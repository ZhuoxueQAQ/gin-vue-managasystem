package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ManageSystemSettingService struct {
}

// CreateManageSystemSetting 创建ManageSystemSetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (manageSystemSettingService *ManageSystemSettingService) CreateManageSystemSetting(manageSystemSetting autocode.ManageSystemSetting) (err error) {
	err = global.GVA_DB.Create(&manageSystemSetting).Error
	return err
}

// DeleteManageSystemSetting 删除ManageSystemSetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (manageSystemSettingService *ManageSystemSettingService) DeleteManageSystemSetting(manageSystemSetting autocode.ManageSystemSetting) (err error) {
	err = global.GVA_DB.Delete(&manageSystemSetting).Error
	return err
}

// DeleteManageSystemSettingByIds 批量删除ManageSystemSetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (manageSystemSettingService *ManageSystemSettingService) DeleteManageSystemSettingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.ManageSystemSetting{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateManageSystemSetting 更新ManageSystemSetting记录
// 只更新三个字段中的一个
// Author [piexlmax](https://github.com/piexlmax)
func (manageSystemSettingService *ManageSystemSettingService) UpdateManageSystemSetting(manageSystemSetting autocode.ManageSystemSetting) (err error) {
	var query string
	if len(manageSystemSetting.ProjectTableCols) > 0 {
		query = "project_table_cols"
	}
	if len(manageSystemSetting.IncomeStreamTableCols) > 0 {
		query = "income_stream_table_cols"
	}
	if len(manageSystemSetting.OutcomeStreamTableCols) > 0 {
		query = "outcome_stream_table_cols"
	}
	err = global.GVA_DB.Model(&manageSystemSetting).Select(query).Updates(manageSystemSetting).Error
	return err
}

// GetManageSystemSetting 根据id获取ManageSystemSetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (manageSystemSettingService *ManageSystemSettingService) GetManageSystemSetting(id uint) (err error, manageSystemSetting autocode.ManageSystemSetting) {
	err = global.GVA_DB.Where("id = ?", id).First(&manageSystemSetting).Error
	return
}

// GetManageSystemSettingInfoList 分页获取ManageSystemSetting记录
// Author [piexlmax](https://github.com/piexlmax)
func (manageSystemSettingService *ManageSystemSettingService) GetManageSystemSettingInfoList(info autoCodeReq.ManageSystemSettingSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.ManageSystemSetting{})
	var manageSystemSettings []autocode.ManageSystemSetting
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&manageSystemSettings).Error
	return err, manageSystemSettings, total
}
