package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ProjectService struct {
}

// CreateProject 创建Project记录
// Author [piexlmax](https://github.com/piexlmax)
func (projectService *ProjectService) CreateProject(project autocode.Project) (err error) {
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
func (projectService *ProjectService) GetProjectInfoList(info autoCodeReq.ProjectSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.Project{})
	var projects []autocode.Project
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Categories != "" {
		db = db.Where("categories LIKE ?", "%"+info.Categories+"%")
	}
	if info.ProjectCode != "" {
		db = db.Where("project_code LIKE ?", "%"+info.ProjectCode+"%")
	}
	// date
	if info.CreatedDate != nil {
		db = db.Where("created_date >= ?", info.CreatedDate)
	}
	if info.ContractStartDate != nil {
		db = db.Where("contract_start_date >= ?", info.ContractStartDate)
	}
	if info.TrainStartDate != nil {
		db = db.Where("train_start_date >= ?", info.TrainStartDate)
	}
	// todo 日期右边界

	// 委托方、落地机构和技术方查询
	//if info.Client != nil {
	//	for index, val := range info.Client {
	//		query := fmt.Sprintf("client->'$[%v].Name = (?)'", index)
	//		db = db.Where(query, val)
	//	}
	//}
	//if info.LandingAgency != nil {
	//	for index, val := range info.LandingAgency {
	//		query := fmt.Sprintf("landing_agency->'$[%v].Name = (?)'", index)
	//		db = db.Where(query, val)
	//	}
	//}
	//if info.Partner != nil {
	//	for index, val := range info.LandingAgency {
	//		query := fmt.Sprintf("partners->'$[%v].Name = (?)'", index)
	//		db = db.Where(query, val)
	//	}
	//}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&projects).Error

	return err, projects, total
}
