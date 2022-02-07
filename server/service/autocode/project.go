package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/shopspring/decimal"
	"log"
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
		// todo 删除对应的支出流水
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
func (projectService *ProjectService) GetProjectInfoList(info autoCodeReq.ProjectSearch) (err error, list interface{},
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
	err = db.Limit(limit).Offset(offset).Find(&projects).Error

	// todo 导出的时候，取消分页，返回total。前端来进行导出？

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
