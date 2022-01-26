package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"log"
)

type ProjectService struct {
}

// CreateProject 创建Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) CreateProject(project autocode.Project) (err error) {
	log.Println(project)
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
	err = global.GVA_DB.Delete(&[]autocode.Project{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateProject 更新Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) UpdateProject(project autocode.Project) (err error) {
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
	// todo 修复bug
	for index, val := range *(info.Client) {
		if val.Name != "" {
			query := fmt.Sprintf("client->'$[%v].name' LIKE ? ", index)
			db = db.Where(query, "%"+val.Name+"%")
		}
	}
	for index, val := range *(info.LandingAgency) {
		if val.Name != "" {
			query := fmt.Sprintf("landing_agency->'$[%v].name' LIKE ? ", index)
			db = db.Where(query, "%"+val.Name+"%")
		}
	}
	for index, val := range *(info.Partner) {
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
