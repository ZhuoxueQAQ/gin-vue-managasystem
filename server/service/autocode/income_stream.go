package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type IncomeStreamService struct {
}

// CreateIncomeStream 创建IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) CreateIncomeStream(Income autocode.IncomeStream) (err error) {
	err = global.GVA_DB.Create(&Income).Error
	return err
}

// DeleteIncomeStream 删除IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) DeleteIncomeStream(Income autocode.IncomeStream) (err error) {
	err = global.GVA_DB.Delete(&Income).Error
	return err
}

// DeleteIncomeStreamByIds 批量删除IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) DeleteIncomeStreamByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.IncomeStream{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIncomeStream 更新IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) UpdateIncomeStream(Income autocode.IncomeStream) (err error) {
	err = global.GVA_DB.Save(&Income).Error
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
func (IncomeService *IncomeStreamService) GetIncomeStreamInfoList(info autoCodeReq.IncomeStreamSearch) (err error, list interface{}, total int64) {
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
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&Incomes).Error
	return err, Incomes, total
}
