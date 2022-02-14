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
	"path/filepath"
)

type IncomeStreamService struct {
}

// UpdateProjectWhenCreateIncomeStream 在创建收入流水时计算收入流水的各项分成，以及更新到对应的培训项目中。传进去的income和project都会被修改
func UpdateProjectWhenCreateIncomeStream(income autocode.IncomeStream, project autocode.Project) (autocode.IncomeStream, autocode.Project) {
	// 如果是刚立项的项目，录入第一笔流水以后就设置成进行中状态，无法修改委托方等等的信息，这样在修改项目的时候就不用担心收入流水对不上了
	if project.Status == 0 {
		project.Status = 1
	}
	r := decimal.NewFromFloat(0.01)
	// 入账金额
	iAmount := income.IncomeAmount
	// todo：先计算流水的，再更新项目的
	// 当入账金额>0的时候，计算收入流水的xx分成，并更新项目的xx分成。
	if iAmount.Cmp(decimal.NewFromFloat(0.0)) > 0 {
		// 更新已到账费用，未到账金额
		project.PaidAmount = project.PaidAmount.Add(iAmount)
		project.UnpaidAmount = project.UnpaidAmount.Sub(iAmount)
		// range中的e是拷贝的副本。。。
		for index, e := range project.Client {
			// 计算流水的某个委托方之类的分成
			income.Client[index].Amount = r.Mul(e.Radio).Mul(iAmount)
			// 然后更新对应项目的对应分成。下面同理
			project.Client[index].IncomeAmount = e.IncomeAmount.Add(income.Client[index].Amount)
		}
		for index, e := range project.LandingAgency {
			income.LandingAgency[index].Amount = r.Mul(e.Radio).Mul(iAmount)
			project.LandingAgency[index].IncomeAmount = e.IncomeAmount.Add(income.LandingAgency[index].Amount)
		}
		for index, e := range project.Partner {
			income.Partner[index].Amount = r.Mul(e.Radio).Mul(iAmount)
			project.Partner[index].IncomeAmount = e.IncomeAmount.Add(income.Partner[index].Amount)
		}
		// 学校管理费，等等。。
		income.SAmount = income.SRadio.Mul(r).Mul(iAmount)
		income.DAmount = income.DRadio.Mul(r).Mul(iAmount)
		income.WAmount = income.WRadio.Mul(r).Mul(iAmount)
		income.CAmount = income.CRadio.Mul(r).Mul(iAmount)
		project.SAmount = project.SAmount.Add(income.SAmount)
		project.DAmount = project.DAmount.Add(income.DAmount)
		project.WAmount = project.WAmount.Add(income.WAmount)
		project.CAmount = project.CAmount.Add(income.CAmount)
		// todo 更新项目的总收入-总管理费
		project.IncomeAndOutcome[0].PTotal = project.IncomeAndOutcome[0].PTotal.Add(iAmount).Sub(income.SAmount)

	}
	// 收入,这里就不用反射实现了。
	project.IncomeAndOutcome[0].Pg = project.IncomeAndOutcome[0].Pg.Add(income.Pays.Pg)
	project.IncomeAndOutcome[0].Ph = project.IncomeAndOutcome[0].Ph.Add(income.Pays.Ph)
	project.IncomeAndOutcome[0].Pi = project.IncomeAndOutcome[0].Pi.Add(income.Pays.Pi)
	project.IncomeAndOutcome[0].Pj = project.IncomeAndOutcome[0].Pj.Add(income.Pays.Pj)
	project.IncomeAndOutcome[0].Pk = project.IncomeAndOutcome[0].Pk.Add(income.Pays.Pk)
	project.IncomeAndOutcome[0].Pl = project.IncomeAndOutcome[0].Pl.Add(income.Pays.Pl)
	project.IncomeAndOutcome[0].Pm = project.IncomeAndOutcome[0].Pm.Add(income.Pays.Pm)
	project.IncomeAndOutcome[0].Pn = project.IncomeAndOutcome[0].Pn.Add(income.Pays.Pn)
	project.IncomeAndOutcome[0].Po = project.IncomeAndOutcome[0].Po.Add(income.Pays.Po)
	project.IncomeAndOutcome[0].Pp = project.IncomeAndOutcome[0].Pp.Add(income.Pays.Pp)
	project.IncomeAndOutcome[0].Pq = project.IncomeAndOutcome[0].Pq.Add(income.Pays.Pq)
	project.IncomeAndOutcome[0].Pr = project.IncomeAndOutcome[0].Pr.Add(income.Pays.Pr)
	project.IncomeAndOutcome[0].Ps = project.IncomeAndOutcome[0].Ps.Add(income.Pays.Ps)
	project.IncomeAndOutcome[0].Pt = project.IncomeAndOutcome[0].Pt.Add(income.Pays.Pt)
	return income, project
}

// UpdateProjectWhenDeleteIncomeStream 在删除收入流水时更新到对应的培训项目中。传进去的project会被修改
func UpdateProjectWhenDeleteIncomeStream(income autocode.IncomeStream, project autocode.Project) autocode.Project {
	iAmount := income.IncomeAmount
	// todo 先计算流水的，再更新项目的
	// 当入账金额>0的时候，计算收入流水的xx分成，并更新项目的xx分成。
	if iAmount.Cmp(decimal.NewFromFloat(0.0)) > 0 {
		// 更新已到账费用，未到账金额
		project.PaidAmount = project.PaidAmount.Sub(iAmount)
		project.UnpaidAmount = project.UnpaidAmount.Add(iAmount)
		for index, e := range project.Client {
			// 更新对应项目的对应分成。下面同理
			project.Client[index].IncomeAmount = e.IncomeAmount.Sub(income.Client[index].Amount)
		}
		for index, e := range project.LandingAgency {
			project.LandingAgency[index].IncomeAmount = e.IncomeAmount.Sub(income.LandingAgency[index].Amount)
		}
		for index, e := range project.Partner {
			project.Partner[index].IncomeAmount = e.IncomeAmount.Sub(income.Partner[index].Amount)
		}
		// 学校管理费，等等。。
		project.SAmount = project.SAmount.Sub(income.SAmount)
		project.DAmount = project.DAmount.Sub(income.DAmount)
		project.WAmount = project.WAmount.Sub(income.WAmount)
		project.CAmount = project.CAmount.Sub(income.CAmount)
		// todo 更新项目的总收入-总管理费。减去该流水的入账金额然后加上该流水的学校管理费
		project.IncomeAndOutcome[0].PTotal = project.IncomeAndOutcome[0].PTotal.Sub(iAmount).Add(income.SAmount)
	}
	// 收入,这里就不用反射实现了。
	project.IncomeAndOutcome[0].Pg = project.IncomeAndOutcome[0].Pg.Sub(income.Pays.Pg)
	project.IncomeAndOutcome[0].Ph = project.IncomeAndOutcome[0].Ph.Sub(income.Pays.Ph)
	project.IncomeAndOutcome[0].Pi = project.IncomeAndOutcome[0].Pi.Sub(income.Pays.Pi)
	project.IncomeAndOutcome[0].Pj = project.IncomeAndOutcome[0].Pj.Sub(income.Pays.Pj)
	project.IncomeAndOutcome[0].Pk = project.IncomeAndOutcome[0].Pk.Sub(income.Pays.Pk)
	project.IncomeAndOutcome[0].Pl = project.IncomeAndOutcome[0].Pl.Sub(income.Pays.Pl)
	project.IncomeAndOutcome[0].Pm = project.IncomeAndOutcome[0].Pm.Sub(income.Pays.Pm)
	project.IncomeAndOutcome[0].Pn = project.IncomeAndOutcome[0].Pn.Sub(income.Pays.Pn)
	project.IncomeAndOutcome[0].Po = project.IncomeAndOutcome[0].Po.Sub(income.Pays.Po)
	project.IncomeAndOutcome[0].Pp = project.IncomeAndOutcome[0].Pp.Sub(income.Pays.Pp)
	project.IncomeAndOutcome[0].Pq = project.IncomeAndOutcome[0].Pq.Sub(income.Pays.Pq)
	project.IncomeAndOutcome[0].Pr = project.IncomeAndOutcome[0].Pr.Sub(income.Pays.Pr)
	project.IncomeAndOutcome[0].Ps = project.IncomeAndOutcome[0].Ps.Sub(income.Pays.Ps)
	project.IncomeAndOutcome[0].Pt = project.IncomeAndOutcome[0].Pt.Sub(income.Pays.Pt)
	return project
}

// CreateIncomeStream 创建IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) CreateIncomeStream(income autocode.IncomeStream) (err error) {
	// 更新对应的培训项目
	projectDb := global.GVA_DB.Model(&autocode.Project{})
	var project autocode.Project
	// 获取对应的培训项目
	if err = projectDb.First(&project, income.ProjectId).Error; err != nil {
		return err
	}
	// todo 状态不是“进行中|立项”的时候不允许创建收入和支出流水
	if project.Status > 1 {
		err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法修改|删除流水", project.Name, utils.FormatProjectStatus(project.Status)))
		return err
	}
	income, project = UpdateProjectWhenCreateIncomeStream(income, project)
	// 把新的收入流水和更新了的培训项目添加到db
	if err = projectDb.Save(&project).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Create(&income).Error
	return err
}

// DeleteIncomeStream 删除IncomeStream记录
// 对培训项目的修改和创建基本一样，区别是add变成sub
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) DeleteIncomeStream(income autocode.IncomeStream) (err error) {
	// 更新对应的培训项目
	var project autocode.Project
	// 获取对应的培训项目
	if err = global.GVA_DB.Where("id = ?", income.ProjectId).First(&project).Error; err != nil {
		return err
	}
	// todo 状态不是“进行中”的时候不允许更新收入和支出流水
	if project.Status != 1 {
		err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法修改流水", project.Name, utils.FormatProjectStatus(project.Status)))
		return err
	}
	project = UpdateProjectWhenDeleteIncomeStream(income, project)
	// 更新了的培训项目添加到db
	if err = global.GVA_DB.Save(&project).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Unscoped().Delete(&income).Error
	return err
}

// DeleteIncomeStreamByIds 批量删除IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) DeleteIncomeStreamByIds(ids request.IdsReq) (err error) {
	// todo 一条一条删除，动态更新流水对应的培训项目
	var incomes []autocode.IncomeStream
	if err = global.GVA_DB.Model(&[]autocode.IncomeStream{}).Where("id in ?", ids.Ids).Find(&incomes).Error; err != nil {
		return err
	}
	for _, income := range incomes {
		var project autocode.Project
		// 找到待更新的培训项目
		if err = global.GVA_DB.Where("id = ?", income.ProjectId).First(&project).Error; err != nil {
			return err
		}
		// todo 状态不是“进行中”的时候不允许删除收入和支出流水
		if project.Status != 1 {
			err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法删除流水", project.Name, utils.FormatProjectStatus(project.Status)))
			return err
		}
		// 更新培训项目
		project = UpdateProjectWhenDeleteIncomeStream(income, project)
		if err = global.GVA_DB.Unscoped().Delete(&income).Error; err != nil {
			return err
		}
		if err = global.GVA_DB.Save(&project).Error; err != nil {
			return err
		}
	}
	// err = global.GVA_DB.Delete(&[]autocode.IncomeStream{}, "id in ?", ids.Ids).Error
	return nil
}

// UpdateIncomeStream 更新IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) UpdateIncomeStream(income autocode.IncomeStream) (err error) {
	// todo 一条一条更新，动态更新流水对应的培训项目

	var project autocode.Project
	var incomeBeforeUpdate autocode.IncomeStream
	if err = global.GVA_DB.Where("id = ?", income.ProjectId).First(&project).Error; err != nil {
		return err
	}
	// todo 状态不是“进行中”的时候不允许更新收入和支出流水
	if project.Status != 1 {
		err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法修改流水", project.Name, utils.FormatProjectStatus(project.Status)))
		return err
	}
	// 找到更新前的收入流水
	if err = global.GVA_DB.Where("id = ?", income.ID).First(&incomeBeforeUpdate).Error; err != nil {
		return err
	}
	// 把更新前的收入流水传入该函数更新项目，就把更新前的流水叠加在项目上的金额全去掉了
	project = UpdateProjectWhenDeleteIncomeStream(incomeBeforeUpdate, project)
	// 然后把更新后的收入流水传进去更新项目，就正确更新的需要更新的项目,以及收入流水（收入流水更新以后，分成需要重新计算，因为入账金额可能变了）
	// 可以根据入账金额是否变化优化，但是懒得搞了
	income, project = UpdateProjectWhenCreateIncomeStream(income, project)
	// 然后将正确更新的项目和流水保存到数据库里
	if err = global.GVA_DB.Save(&project).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Save(&income).Error
	return err
}

// GetIncomeStream 根据id获取IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) GetIncomeStream(id uint) (err error, Income autocode.IncomeStream) {
	err = global.GVA_DB.Where("id = ?", id).First(&Income).Error
	return
}

// GetIncomeStreamInfoList 分页获取IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) GetIncomeStreamInfoList(info autoCodeReq.IncomeStreamSearch, isExport bool) (err error,
	list []autocode.IncomeStream, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.IncomeStream{})
	var Incomes []autocode.IncomeStream
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProjectName != "" {
		db = db.Where("project_name LIKE ?", "%"+info.ProjectName+"%")
	}
	for index, date := range info.CreatedDateRange {
		query := "created_date >= ?"
		if index == 1 {
			query = "created_date <= ?"
		}
		db = db.Where(query, date)
	}
	for index, date := range info.SplitAmountDateRange {
		query := "split_amount_date >= ?"
		if index == 1 {
			query = "split_amount_date <= ?"
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
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if isExport {
		err = db.Find(&Incomes).Error
	} else {
		err = db.Limit(limit).Offset(offset).Find(&Incomes).Error
	}

	return err, Incomes, total
}

func (IncomeService *IncomeStreamService) ExportToExcel(incomes []autocode.IncomeStream) (err error, path string) {
	f := excelize.NewFile()
	sheetName := "表"
	f.SetSheetName("Sheet1", sheetName)
	data := [][]interface{}{
		{"培训项目", "项目分类", "是否预先借票", "是否已冲账", "缴费单位", "发票抬头", "入账日期", "入账金额", "入账凭证编码", "发票金额", "发票数量", "发票编号", "分账日期", "发票开出日期",
			"委托方1", nil, nil, "委托方2", nil, nil, "委托方3", nil, nil,
			"落地机构1", nil, nil, "落地机构2", nil, nil, "落地机构3", nil, nil,
			"技术方1", nil, nil, "技术方2", nil, nil, "技术方3", nil, nil,
			"学校管理费180043", nil, "发展基金210098", nil, "福利220121", nil, "课酬440120", nil,
			"专家课酬", "方案费", "网络研修教学费", "管理及工作人员劳务费", "考察、跟岗、线下指导费", "专家食宿", "专家及工作人员交通费", "学员伙食费", "学员住宿费", "学员交通费", "场租", "课程资源建设费(人员支出)", "课程资源建设费(购买)", "培训资料费、办公用品费、保险费、医药费及其他(含购买标书、中标服务费等)",
			"备注"},
		{nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			"名称", "分成比例", "分成金额", "名称", "分成比例", "分成金额", "名称", "分成比例", "分成金额",
			"名称", "分成比例", "分成金额", "名称", "分成比例", "分成金额", "名称", "分成比例", "分成金额",
			"名称", "分成比例", "分成金额", "名称", "分成比例", "分成金额", "名称", "分成比例", "分成金额",
			"比例", "分成金额", "比例", "分成金额", "比例", "分成金额", "比例", "分成金额",
			nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			nil},
	}
	// 插入project data
	for _, p := range incomes {
		left := []interface{}{
			p.ProjectName, p.Categories, p.IsBorrowed, p.IsOffset, p.PayUnit, p.Invoice, utils.FormatDate(p.CreatedDate, false),
			p.IncomeAmount, p.IncomeValidCode, p.InvoiceAmount, p.InvoiceCount, p.InvoiceCode,
			utils.FormatDate(p.SplitAmountDate, false), utils.FormatDate(p.InvoiceIssueDate, false),
		}
		// 委托方、等的最大个数
		maxFormRowNum := 3
		for _, c := range p.Client {
			tmp := []interface{}{c.Name, c.Radio.String() + "%", c.Amount}
			left = append(left, tmp...)
		}
		// 对小于最大长度的部分补充{nil, nil, nil}
		for i := len(p.Client) * 3; i < maxFormRowNum*3; i++ {
			left = append(left, nil)
		}
		for _, c := range p.LandingAgency {
			tmp := []interface{}{c.Name, c.Radio.String() + "%", c.Amount}
			left = append(left, tmp...)
		}
		for i := len(p.LandingAgency) * 3; i < maxFormRowNum*3; i++ {
			left = append(left, nil)
		}
		for _, c := range p.Partner {
			tmp := []interface{}{c.Name, c.Radio.String() + "%", c.Amount}
			left = append(left, tmp...)
		}
		for i := len(p.Partner) * 3; i < maxFormRowNum*3; i++ {
			left = append(left, nil)
		}
		right := []interface{}{
			p.SRadio.String() + "%", p.SAmount, p.DRadio.String() + "%", p.DAmount, p.WRadio.String() + "%", p.WAmount, p.CRadio.String() + "%", p.CAmount,
			p.Pays.Pg, p.Pays.Ph, p.Pays.Pi, p.Pays.Pj, p.Pays.Pk, p.Pays.Pl, p.Pays.Pm, p.Pays.Pn, p.Pays.Po, p.Pays.Pp, p.Pays.Pq, p.Pays.Pr, p.Pays.Ps, p.Pays.Pt,
			p.Remark,
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
		{"A1", "A2"}, {"B1", "B2"}, {"C1", "C2"}, {"D1", "D2"}, {"E1", "E2"}, {"F1", "F2"}, {"G1", "G2"}, {"H1", "H2"}, {"I1", "I2"}, {"J1", "J2"}, {"K1", "K2"}, {"L1", "L2"}, {"M1", "M2"}, {"N1", "N2"},
		{"O1", "Q1"}, {"R1", "T1"}, {"U1", "W1"}, {"X1", "Z1"}, {"AA1", "AC1"}, {"AD1", "AF1"}, {"AG1", "AI1"}, {"AJ1", "AL1"}, {"AM1", "AO1"},
		{"AP1", "AQ1"}, {"AR1", "AS1"}, {"AT1", "AU1"}, {"AV1", "AW1"},
		{"AX1", "AX2"}, {"AY1", "AY2"}, {"AZ1", "AZ2"}, {"BA1", "BA2"}, {"BB1", "BB2"}, {"BC1", "BC2"}, {"BD1", "BD2"}, {"BE1", "BE2"}, {"BF1", "BF2"}, {"BG1", "BG2"}, {"BH1", "BH2"}, {"BI1", "BI2"}, {"BJ1", "BJ2"}, {"BK1", "BK2"},
		{"BL1", "BL2"},
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
	if err = f.SetCellStyle(sheetName, "A1", fmt.Sprintf("CL%v", len(incomes)+2), style); err != nil {
		return err, ""
	}
	if err = f.SetColWidth(sheetName, "A", "BL", 15); err != nil {

	}
	path = filepath.Join("./fileDir", "导出数据.xlsx")
	if err = f.SaveAs(path); err != nil {
		return err, ""
	}
	return nil, path
}
