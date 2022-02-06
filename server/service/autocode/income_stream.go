package autocode

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/shopspring/decimal"
)

type IncomeStreamService struct {
}

// CreateIncomeStream 创建IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) CreateIncomeStream(Income autocode.IncomeStream) (err error) {
	// 更新对应的培训项目
	projectDb := global.GVA_DB.Model(&autocode.Project{})
	var project autocode.Project
	// 获取对应的培训项目
	if err = projectDb.First(&project, Income.ProjectId).Error; err != nil {
		return err
	}
	// 因为比率存的是1-100的百分数，所以
	r := decimal.NewFromFloat(0.01)
	// 入账金额
	iAmount := Income.IncomeAmount
	// todo 先计算流水的，再更新项目的
	// 当入账金额>0的时候，计算收入流水的xx分成，并更新项目的xx分成。
	if iAmount.Cmp(decimal.NewFromFloat(0.0)) > 0 {
		// 更新已到账费用，未到账金额
		project.PaidAmount = project.PaidAmount.Add(iAmount)
		project.UnpaidAmount = project.UnpaidAmount.Sub(iAmount)
		for index, e := range project.Client {
			// 计算流水的某个委托方之类的分成
			Income.Client[index].Amount = r.Mul(e.Radio).Mul(iAmount)
			// 然后更新对应项目的对应分成。下面同理
			e.IncomeAmount = e.IncomeAmount.Add(Income.Client[index].Amount)
		}
		for index, e := range project.LandingAgency {
			Income.LandingAgency[index].Amount = r.Mul(e.Radio).Mul(iAmount)
			e.IncomeAmount = e.IncomeAmount.Add(Income.LandingAgency[index].Amount)
		}
		for index, e := range project.Partner {
			Income.Partner[index].Amount = r.Mul(e.Radio).Mul(iAmount)
			e.IncomeAmount = e.IncomeAmount.Add(Income.Partner[index].Amount)
		}
		// 学校管理费，等等。。
		Income.SAmount = Income.SRadio.Mul(r).Mul(iAmount)
		Income.DAmount = Income.DRadio.Mul(r).Mul(iAmount)
		Income.WAmount = Income.WRadio.Mul(r).Mul(iAmount)
		Income.CAmount = Income.CRadio.Mul(r).Mul(iAmount)
		project.SAmount = project.SAmount.Add(Income.SAmount)
		project.DAmount = project.DAmount.Add(Income.DAmount)
		project.WAmount = project.WAmount.Add(Income.WAmount)
		project.CAmount = project.CAmount.Add(Income.CAmount)
	}
	// 收入,这里就不用反射实现了。
	project.IncomeAndOutcome[0].Pg = project.IncomeAndOutcome[0].Pg.Add(Income.Pays.Pg)
	project.IncomeAndOutcome[0].Ph = project.IncomeAndOutcome[0].Ph.Add(Income.Pays.Ph)
	project.IncomeAndOutcome[0].Pi = project.IncomeAndOutcome[0].Pi.Add(Income.Pays.Pi)
	project.IncomeAndOutcome[0].Pj = project.IncomeAndOutcome[0].Pj.Add(Income.Pays.Pj)
	project.IncomeAndOutcome[0].Pk = project.IncomeAndOutcome[0].Pk.Add(Income.Pays.Pk)
	project.IncomeAndOutcome[0].Pl = project.IncomeAndOutcome[0].Pl.Add(Income.Pays.Pl)
	project.IncomeAndOutcome[0].Pm = project.IncomeAndOutcome[0].Pm.Add(Income.Pays.Pm)
	project.IncomeAndOutcome[0].Pn = project.IncomeAndOutcome[0].Pn.Add(Income.Pays.Pn)
	project.IncomeAndOutcome[0].Po = project.IncomeAndOutcome[0].Po.Add(Income.Pays.Po)
	project.IncomeAndOutcome[0].Pp = project.IncomeAndOutcome[0].Pp.Add(Income.Pays.Pp)
	project.IncomeAndOutcome[0].Pq = project.IncomeAndOutcome[0].Pq.Add(Income.Pays.Pq)
	project.IncomeAndOutcome[0].Pr = project.IncomeAndOutcome[0].Pr.Add(Income.Pays.Pr)
	project.IncomeAndOutcome[0].Ps = project.IncomeAndOutcome[0].Ps.Add(Income.Pays.Ps)
	project.IncomeAndOutcome[0].Pt = project.IncomeAndOutcome[0].Pt.Add(Income.Pays.Pt)
	// 把新的收入流水和更新了的培训项目添加到db
	if err = projectDb.Save(&project).Error; err != nil {
		return err
	}
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
	// todo 一条一条删除，动态更新流水对应的培训项目
	err = global.GVA_DB.Delete(&[]autocode.IncomeStream{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIncomeStream 更新IncomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (IncomeService *IncomeStreamService) UpdateIncomeStream(Income autocode.IncomeStream) (err error) {
	// todo 一条一条更新，动态更新流水对应的培训项目
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
	err = db.Limit(limit).Offset(offset).Find(&Incomes).Error
	return err, Incomes, total
}
