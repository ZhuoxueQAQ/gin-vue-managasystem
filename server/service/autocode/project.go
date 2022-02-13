package autocode

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"log"
	"path/filepath"
	"strconv"
)

type ProjectService struct {
}

// CreateProject 创建Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) CreateProject(project autocode.Project) (err error) {
	log.Println(project)
	// 对已到账费用和未到账金额赋值
	project.PaidAmount = decimal.NewFromFloat(0.0)
	project.UnpaidAmount = project.ProjectAmount
	err = global.GVA_DB.Create(&project).Error
	return err
}

// DeleteProject 删除Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) DeleteProject(project autocode.Project) (err error) {
	err = global.GVA_DB.Delete(&project).Error
	return err
}

// DeleteProjectByIds 批量删除Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) DeleteProjectByIds(ids request.IdsReq) (err error) {
	// todo delete project files and file records, incomestream and outcomestream
	var projects []autocode.Project
	if err = global.GVA_DB.Model(&[]autocode.Project{}).Where("id in ?", ids.Ids).Find(&projects).Error; err != nil {
		return err
	}
	for _, p := range projects {
		// 删除对应的收入流水
		if err = global.GVA_DB.Unscoped().Delete(&autocode.IncomeStream{}, "project_id = ?", p.ID).Error; err != nil {
			return err
		}
		// 删除对应的支出流水
		if err = global.GVA_DB.Unscoped().Delete(&autocode.OutcomeStream{}, "project_id = ?", p.ID).Error; err != nil {
			return err
		}
		// 删除对应的文件记录和文件
		strProjectID := strconv.Itoa(int(p.ID))
		if err = global.GVA_DB.Unscoped().Delete(&autocode.ProjectFileRecord{}, "project_id = ?", strProjectID).Error; err != nil {
			return err
		}
		if err = utils.DeleteProjectFiles(strProjectID); err != nil {
			return err
		}
		// 最后删除项目
		if err = global.GVA_DB.Unscoped().Delete(&p).Error; err != nil {
			return err
		}
	}
	//err = global.GVA_DB.Delete(&[]autocode.Project{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProject 更新Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) UpdateProject(project autocode.Project) (err error) {
	// todo 如果已经有流水（不在立项状态了），不允许更新。
	// 更新未到账金额(如果项目应收经费更新了)
	project.UnpaidAmount = project.ProjectAmount.Sub(project.PaidAmount)
	err = global.GVA_DB.Save(&project).Error
	return err
}

// GetProject 根据id获取Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) GetProject(id uint) (err error, project autocode.Project) {
	err = global.GVA_DB.Where("id = ?", id).First(&project).Error
	return
}

// GetProjectInfoList 分页获取Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) GetProjectInfoList(info autoCodeReq.ProjectSearch, isExport bool) (err error, list []autocode.Project,
	total int64) {

	// 创建db
	db := global.GVA_DB.Model(&autocode.Project{})
	// 分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var projects []autocode.Project

	// todo 当筛选条件改变时才筛选（pagesize=1？）

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Categories != "" {
		db = db.Where("categories LIKE ?", "%"+info.Categories+"%")
	}
	if info.ProjectCode != "" {
		db = db.Where("project_code LIKE ?", "%"+info.ProjectCode+"%")
	}
	// todo date
	for index, date := range info.CreatedDateRange {
		query := "created_date >= ?"
		if index == 1 {
			query = "created_date <= ?"
		}
		db = db.Where(query, date)
	}
	for index, date := range info.ContractStartDateRange {
		// todo 日期变成范围查询
		query := "contract_start_date > ?"
		if index == 1 {
			query = "contract_start_date < ?"
		}
		db = db.Where(query, date)
	}
	for index, date := range info.TrainStartDateRange {
		query := "train_start_date >= ?"
		if index == 1 {
			query = "train_start_date <= ?"
		}
		db = db.Where(query, date)
	}

	// 根据委托方、落地机构和技术方的名字筛选
	for index, val := range info.Client {
		if val.Name != "" {
			query := fmt.Sprintf("client->'$[%v].name' LIKE ? ", index)
			db = db.Where(query, "%"+val.Name+"%")
		}
	}
	for index, val := range info.LandingAgency {
		if val.Name != "" {
			query := fmt.Sprintf("landing_agency->'$[%v].name' LIKE ? ", index)
			db = db.Where(query, "%"+val.Name+"%")
		}
	}
	for index, val := range info.Partner {
		if val.Name != "" {
			query := fmt.Sprintf("partner->'$[%v].name' LIKE ? ", index)
			db = db.Where(query, "%"+val.Name+"%")
		}
	}
	// total 是已经筛选的数据的条数
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if isExport {
		err = db.Find(&projects).Error
	} else {
		err = db.Limit(limit).Offset(offset).Find(&projects).Error
	}

	return err, projects, total
}

func (projectService *ProjectService) GetProjectFileRecordInfoList(info autoCodeReq.ProjectFileRecordSearch) (err error, list interface{},
	total int64) {
	var projectFileRecords []autocode.ProjectFileRecord
	// 创建db
	db := global.GVA_DB.Model(&autocode.ProjectFileRecord{})
	// 分页参数
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db = db.Where("project_id = ? and file_type_id = ?", info.ProjectID, info.FileTypeID)
	// total 是已经筛选的数据的条数
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projectFileRecords).Error
	return err, projectFileRecords, total
}

func (projectService *ProjectService) DeleteProjectFileByIds(ids request.IdsReq) (err error) {
	var records []autocode.ProjectFileRecord
	// 按数组查询
	if err := global.GVA_DB.Model(&[]autocode.ProjectFileRecord{}).Where("id in ?", ids.Ids).Find(&records).Error; err != nil {
		return err
	}
	for _, record := range records {
		// 查询待删除的记录
		// 删除文件
		if err = utils.DeleteProjectFile(record); err != nil {
			return err
		}
		// 删除文件记录
		if err = global.GVA_DB.Unscoped().Delete(&record).Error; err != nil {
			return err
		}
	}
	return nil
}

func (projectService *ProjectService) CreateUploadProjectFileRecord(record autocode.ProjectFileRecord) (err error) {
	return global.GVA_DB.Create(&record).Error
}

func (projectService *ProjectService) ChangeProjectStatus(projectStatus autocode.ProjectStatus) (err error) {
	if projectStatus.Status == 0 {
		return errors.New("无法将项目状态修改为[立项]！")
	}
	if projectStatus.Status > 4 {
		return errors.New("该状态未定义！")
	}
	var project autocode.Project
	projectDB := global.GVA_DB.Model(&autocode.Project{})
	// 获取对应的培训项目修改其状态
	if err = projectDB.First(&project, projectStatus.ProjectId).Error; err != nil {
		return err
	}
	// 只更新项目的状态列
	if err = global.GVA_DB.Model(&project).Update("status", projectStatus.Status).Error; err != nil {
		return err
	}
	return nil
}

func (projectService *ProjectService) ExportToExcel(projects []autocode.Project) (err error, path string) {
	f := excelize.NewFile()
	sheetName := "表"
	f.SetSheetName("Sheet1", sheetName)
	data := [][]interface{}{
		{"状态", "项目名称", "备案申请日期", "项目分类", "项目所属地", "收费标准", "负责人", "培训模式", "学时数", "培训开始时间", "培训结束时间",
			"合同开始时间", "合同结束时间", "项目应收费用", "已到账费用", "未到账金额", "委托方1", nil, nil, nil, "委托方2", nil, nil, nil,
			"委托方3", nil, nil, nil, "落地机构1", nil, nil, nil, "落地机构2", nil, nil, nil, "落地机构3", nil, nil, nil, "技术方1", nil, nil, nil,
			"技术方2", nil, nil, nil, "技术方3", nil, nil, nil, "学校管理费180043", nil, "发展基金210098", nil, "福利220121", nil, "课酬440120", nil,
			"专家课酬", nil, "方案费", nil, "网络研修教学费", nil, "管理及工作人员劳务费", nil, "考察、跟岗、线下指导费", nil, "专家食宿", nil,
			"专家及工作人员交通费", nil, "学员伙食费", nil, "学员住宿费", nil, "学员交通费", nil, "场租", nil, "课程资源建设费(人员支出)", nil,
			"课程资源建设费(购买)", nil, "培训资料费、办公用品费、保险费、医药费及其他(含购买标书、中标服务费等)", nil, "余额", "备注"},
		{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "名称", "分成比例", "分成(预算)", "分成(支出)", "比例", "分成", "比例", "分成", "比例", "分成", "比例", "分成", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", "预算", "支出", nil, nil},
	}
	// 插入project data
	for _, p := range projects {
		left := []interface{}{
			utils.FormatProjectStatus(p.Status), p.Name, utils.FormatDate(p.CreatedDate, false), p.Categories, p.Area, p.ChargeStandard, p.Manager, p.TrainMode,
			p.TrainTime, utils.FormatDate(p.TrainStartDate, false), utils.FormatDate(p.TrainEndDate, false),
			utils.FormatDate(p.ContractStartDate, false), utils.FormatDate(p.ContractEndDate, false),
			p.ProjectAmount, p.PaidAmount, p.UnpaidAmount,
		}
		targetNum := 3
		for _, c := range p.Client {
			tmp := []interface{}{c.Name, c.Radio.String() + "%", c.IncomeAmount, c.OutcomeAmount}
			left = append(left, tmp...)
		}
		for i := len(p.Client) * 4; i < targetNum*4; i++ {
			left = append(left, nil)
		}
		for _, c := range p.LandingAgency {
			tmp := []interface{}{c.Name, c.Radio.String() + "%", c.IncomeAmount, c.OutcomeAmount}
			left = append(left, tmp...)
		}
		for i := len(p.LandingAgency) * 4; i < targetNum*4; i++ {
			left = append(left, nil)
		}
		for _, c := range p.Partner {
			tmp := []interface{}{c.Name, c.Radio.String() + "%", c.IncomeAmount, c.OutcomeAmount}
			left = append(left, tmp...)
		}
		for i := len(p.Partner) * 4; i < targetNum*4; i++ {
			left = append(left, nil)
		}
		right := []interface{}{
			p.SRadio.String() + "%", p.SAmount, p.DRadio.String() + "%", p.DAmount, p.WRadio.String() + "%", p.WAmount, p.CRadio.String() + "%", p.CAmount,
			p.IncomeAndOutcome[0].Pg, p.IncomeAndOutcome[1].Pg, p.IncomeAndOutcome[0].Ph, p.IncomeAndOutcome[1].Ph,
			p.IncomeAndOutcome[0].Pi, p.IncomeAndOutcome[1].Pi, p.IncomeAndOutcome[0].Pj, p.IncomeAndOutcome[1].Pj,
			p.IncomeAndOutcome[0].Pk, p.IncomeAndOutcome[1].Pk, p.IncomeAndOutcome[0].Pl, p.IncomeAndOutcome[1].Pl,
			p.IncomeAndOutcome[0].Pm, p.IncomeAndOutcome[1].Pm, p.IncomeAndOutcome[0].Pn, p.IncomeAndOutcome[1].Pn,
			p.IncomeAndOutcome[0].Po, p.IncomeAndOutcome[1].Po, p.IncomeAndOutcome[0].Pp, p.IncomeAndOutcome[1].Pp,
			p.IncomeAndOutcome[0].Pq, p.IncomeAndOutcome[1].Pq, p.IncomeAndOutcome[0].Pr, p.IncomeAndOutcome[1].Pr,
			p.IncomeAndOutcome[0].Ps, p.IncomeAndOutcome[1].Ps, p.IncomeAndOutcome[0].Pt, p.IncomeAndOutcome[1].Pt,
			p.IncomeAndOutcome[0].PTotal.Sub(p.IncomeAndOutcome[1].PTotal), p.Remark,
		}
		// concrete
		left = append(left, right...)
		data = append(data, left)
	}

	// 创建表头
	for i, row := range data {
		// 按行赋值
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			return err, ""
		}
		if err = f.SetSheetRow(sheetName, startCell, &row); err != nil {
			return err, ""
		}
	}
	// 合并单元格
	mergeCellRanges := [][]string{
		{"A1", "A2"}, {"B1", "B2"}, {"C1", "C2"}, {"D1", "D2"}, {"E1", "E2"}, {"F1", "F2"}, {"G1", "G2"}, {"H1", "H2"}, {"I1", "I2"}, {"J1", "J2"}, {"K1", "K2"}, {"L1", "L2"}, {"M1", "M2"}, {"N1", "N2"}, {"O1", "O2"}, {"P1", "P2"},
		{"Q1", "T1"}, {"U1", "X1"}, {"Y1", "AB1"}, {"AC1", "AF1"}, {"AG1", "AJ1"}, {"AK1", "AN1"}, {"AO1", "AR1"}, {"AS1", "AV1"}, {"AW1", "AZ1"},
		{"BA1", "BB1"}, {"BC1", "BD1"}, {"BE1", "BF1"}, {"BG1", "BH1"},
		{"BI1", "BJ1"}, {"BK1", "BL1"}, {"BM1", "BN1"}, {"BO1", "BP1"}, {"BQ1", "BR1"}, {"BS1", "BT1"}, {"BU1", "BV1"}, {"BW1", "BX1"}, {"BY1", "BZ1"}, {"CA1", "CB1"}, {"CC1", "CD1"}, {"CE1", "CF1"}, {"CG1", "CH1"}, {"CI1", "CJ1"},
		{"CK1", "CK2"}, {"CL1", "CL2"},
	}

	for _, ranges := range mergeCellRanges {
		if err = f.MergeCell(sheetName, ranges[0], ranges[1]); err != nil {
			return err, ""
		}
	}
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			Indent:          0,
			JustifyLastLine: true,
			ReadingOrder:    0,
			RelativeIndent:  0,
			ShrinkToFit:     true,
			TextRotation:    0,
			Vertical:        "center",
			WrapText:        true,
		},
	})
	if err != nil {
		return err, ""
	}
	// CLn n是数据量
	if err = f.SetCellStyle(sheetName, "A1", fmt.Sprintf("CL%v", len(projects)+2), style); err != nil {
		return err, ""
	}
	if err = f.SetColWidth(sheetName, "A", "CL", 15); err != nil {

	}
	path = filepath.Join("./fileDir", "导出数据.xlsx")
	if err = f.SaveAs(path); err != nil {
		return err, ""
	}
	return nil, path
}
