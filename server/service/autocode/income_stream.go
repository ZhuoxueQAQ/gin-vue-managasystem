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

// UpdateProjectWhenCreateIncomeStream 在创建收入流水时计算收入流水的各项分成，以及更新到对应的培训项目中。传进去的income和project都会被修改
func UpdateProjectWhenCreateIncomeStream(income autocode.IncomeStream, project autocode.Project) (autocode.IncomeStream, autocode.Project) {
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
	if err = global.GVA_DB.Where("id = ?", income.ProjectName).First(&project).Error; err != nil {
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
