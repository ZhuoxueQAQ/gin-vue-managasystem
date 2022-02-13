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
)

type OutcomeStreamService struct {
}

// UpdateProjectWhenCreateOutcomeStream 在创建支出流水时将流水的各项分成，更新到对应的培训项目中。传进去的project都会被修改
func UpdateProjectWhenCreateOutcomeStream(outcome autocode.OutcomeStream, project autocode.Project) autocode.Project {
	// 如果是刚立项的项目，录入第一笔流水以后就设置成进行中状态，无法修改委托方等等的信息，这样在修改项目的时候就不用担心收入流水对不上了
	// 如果不是刚立项的项目，就不用在这里修改状态
	if project.Status == 0 {
		project.Status = 1
	}
	// todo：该笔流水的总支出金额
	oAmount := decimal.NewFromFloat(0.0)
	// range中的e是拷贝的副本。。。，所以表达式左边用索引获取第几个合作方
	for index, e := range project.Client {
		// 然后更新对应项目的对应分成。下面同理
		project.Client[index].OutcomeAmount = e.OutcomeAmount.Add(outcome.Client[index].Amount)
		oAmount = oAmount.Add(outcome.Client[index].Amount)
	}
	for index, e := range project.LandingAgency {
		project.LandingAgency[index].OutcomeAmount = e.OutcomeAmount.Add(outcome.LandingAgency[index].Amount)
		oAmount = oAmount.Add(outcome.LandingAgency[index].Amount)
	}
	for index, e := range project.Partner {
		project.Partner[index].OutcomeAmount = e.OutcomeAmount.Add(outcome.Partner[index].Amount)
		oAmount = oAmount.Add(outcome.Partner[index].Amount)
	}
	// 把专家课酬这些费用累加到该笔支出流水的总支出
	oAmount = oAmount.Add(outcome.Pays.Pg)
	oAmount = oAmount.Add(outcome.Pays.Ph)
	oAmount = oAmount.Add(outcome.Pays.Pi)
	oAmount = oAmount.Add(outcome.Pays.Pj)
	oAmount = oAmount.Add(outcome.Pays.Pk)
	oAmount = oAmount.Add(outcome.Pays.Pl)
	oAmount = oAmount.Add(outcome.Pays.Pm)
	oAmount = oAmount.Add(outcome.Pays.Pn)
	oAmount = oAmount.Add(outcome.Pays.Po)
	oAmount = oAmount.Add(outcome.Pays.Pp)
	oAmount = oAmount.Add(outcome.Pays.Pq)
	oAmount = oAmount.Add(outcome.Pays.Pr)
	oAmount = oAmount.Add(outcome.Pays.Ps)
	oAmount = oAmount.Add(outcome.Pays.Pt)

	// IncomeAndOutcome是一个数组。IncomeAndOutcome[0]标识收入，IncomeAndOutcome[1]标识支出,这里就不用反射实现了。
	project.IncomeAndOutcome[1].Pg = project.IncomeAndOutcome[1].Pg.Add(outcome.Pays.Pg)
	project.IncomeAndOutcome[1].Ph = project.IncomeAndOutcome[1].Ph.Add(outcome.Pays.Ph)
	project.IncomeAndOutcome[1].Pi = project.IncomeAndOutcome[1].Pi.Add(outcome.Pays.Pi)
	project.IncomeAndOutcome[1].Pj = project.IncomeAndOutcome[1].Pj.Add(outcome.Pays.Pj)
	project.IncomeAndOutcome[1].Pk = project.IncomeAndOutcome[1].Pk.Add(outcome.Pays.Pk)
	project.IncomeAndOutcome[1].Pl = project.IncomeAndOutcome[1].Pl.Add(outcome.Pays.Pl)
	project.IncomeAndOutcome[1].Pm = project.IncomeAndOutcome[1].Pm.Add(outcome.Pays.Pm)
	project.IncomeAndOutcome[1].Pn = project.IncomeAndOutcome[1].Pn.Add(outcome.Pays.Pn)
	project.IncomeAndOutcome[1].Po = project.IncomeAndOutcome[1].Po.Add(outcome.Pays.Po)
	project.IncomeAndOutcome[1].Pp = project.IncomeAndOutcome[1].Pp.Add(outcome.Pays.Pp)
	project.IncomeAndOutcome[1].Pq = project.IncomeAndOutcome[1].Pq.Add(outcome.Pays.Pq)
	project.IncomeAndOutcome[1].Pr = project.IncomeAndOutcome[1].Pr.Add(outcome.Pays.Pr)
	project.IncomeAndOutcome[1].Ps = project.IncomeAndOutcome[1].Ps.Add(outcome.Pays.Ps)
	project.IncomeAndOutcome[1].Pt = project.IncomeAndOutcome[1].Pt.Add(outcome.Pays.Pt)

	// 更新项目总支出。因为是添加支出流水，所以余额是变少
	project.IncomeAndOutcome[1].PTotal = project.IncomeAndOutcome[1].PTotal.Add(oAmount)
	return project
}

// UpdateProjectWhenDeleteOutcomeStream 在删除支出流水时更新到对应的培训项目中。传进去的project会被修改
func UpdateProjectWhenDeleteOutcomeStream(outcome autocode.OutcomeStream, project autocode.Project) autocode.Project {
	// todo：该笔流水的总支出金额
	oAmount := decimal.NewFromFloat(0.0)
	// 当入账金额>0的时候，计算收入流水的xx分成，并更新项目的xx分成。
	for index, e := range project.Client {
		// 更新对应项目的对应分成。下面同理
		project.Client[index].OutcomeAmount = e.OutcomeAmount.Sub(outcome.Client[index].Amount)
		oAmount = oAmount.Add(outcome.Client[index].Amount)
	}
	for index, e := range project.LandingAgency {
		project.LandingAgency[index].OutcomeAmount = e.OutcomeAmount.Sub(outcome.LandingAgency[index].Amount)
		oAmount = oAmount.Add(outcome.LandingAgency[index].Amount)
	}
	for index, e := range project.Partner {
		project.Partner[index].OutcomeAmount = e.OutcomeAmount.Sub(outcome.Partner[index].Amount)
		oAmount = oAmount.Add(outcome.Partner[index].Amount)
	}

	// 把专家课酬这些费用累加到该笔支出流水的总支出
	oAmount = oAmount.Add(outcome.Pays.Pg)
	oAmount = oAmount.Add(outcome.Pays.Ph)
	oAmount = oAmount.Add(outcome.Pays.Pi)
	oAmount = oAmount.Add(outcome.Pays.Pj)
	oAmount = oAmount.Add(outcome.Pays.Pk)
	oAmount = oAmount.Add(outcome.Pays.Pl)
	oAmount = oAmount.Add(outcome.Pays.Pm)
	oAmount = oAmount.Add(outcome.Pays.Pn)
	oAmount = oAmount.Add(outcome.Pays.Po)
	oAmount = oAmount.Add(outcome.Pays.Pp)
	oAmount = oAmount.Add(outcome.Pays.Pq)
	oAmount = oAmount.Add(outcome.Pays.Pr)
	oAmount = oAmount.Add(outcome.Pays.Ps)
	oAmount = oAmount.Add(outcome.Pays.Pt)

	// 专家课酬这些费用..,这里就不用反射实现了。
	project.IncomeAndOutcome[1].Pg = project.IncomeAndOutcome[1].Pg.Sub(outcome.Pays.Pg)
	project.IncomeAndOutcome[1].Ph = project.IncomeAndOutcome[1].Ph.Sub(outcome.Pays.Ph)
	project.IncomeAndOutcome[1].Pi = project.IncomeAndOutcome[1].Pi.Sub(outcome.Pays.Pi)
	project.IncomeAndOutcome[1].Pj = project.IncomeAndOutcome[1].Pj.Sub(outcome.Pays.Pj)
	project.IncomeAndOutcome[1].Pk = project.IncomeAndOutcome[1].Pk.Sub(outcome.Pays.Pk)
	project.IncomeAndOutcome[1].Pl = project.IncomeAndOutcome[1].Pl.Sub(outcome.Pays.Pl)
	project.IncomeAndOutcome[1].Pm = project.IncomeAndOutcome[1].Pm.Sub(outcome.Pays.Pm)
	project.IncomeAndOutcome[1].Pn = project.IncomeAndOutcome[1].Pn.Sub(outcome.Pays.Pn)
	project.IncomeAndOutcome[1].Po = project.IncomeAndOutcome[1].Po.Sub(outcome.Pays.Po)
	project.IncomeAndOutcome[1].Pp = project.IncomeAndOutcome[1].Pp.Sub(outcome.Pays.Pp)
	project.IncomeAndOutcome[1].Pq = project.IncomeAndOutcome[1].Pq.Sub(outcome.Pays.Pq)
	project.IncomeAndOutcome[1].Pr = project.IncomeAndOutcome[1].Pr.Sub(outcome.Pays.Pr)
	project.IncomeAndOutcome[1].Ps = project.IncomeAndOutcome[1].Ps.Sub(outcome.Pays.Ps)
	project.IncomeAndOutcome[1].Pt = project.IncomeAndOutcome[1].Pt.Sub(outcome.Pays.Pt)

	// 更新项目余额。因为是删除支出流水，所以余额是变多
	project.IncomeAndOutcome[1].PTotal = project.IncomeAndOutcome[1].PTotal.Sub(oAmount)
	return project
}

// CreateOutcomeStream 创建OutcomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (outcomeService *OutcomeStreamService) CreateOutcomeStream(outcome autocode.OutcomeStream) (err error) {
	var project autocode.Project
	projectDb := global.GVA_DB.Model(&autocode.Project{})
	// 获取对应的培训项目
	if err = projectDb.First(&project, outcome.ProjectId).Error; err != nil {
		return err
	}
	// todo 状态不是“进行中|立项”的时候不允许创建收入和支出流水
	if project.Status > 1 {
		err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法修改流水", project.Name, utils.FormatProjectStatus(project.Status)))
		return err
	}
	project = UpdateProjectWhenCreateOutcomeStream(outcome, project)
	// 把新的支出流水和更新了的培训项目添加到db
	if err = projectDb.Save(&project).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Create(&outcome).Error
	return err
}

// DeleteOutcomeStream 删除OutcomeStream记录c
// Author [piexlmax](https://github.com/piexlmax)
func (outcomeService *OutcomeStreamService) DeleteOutcomeStream(outcome autocode.OutcomeStream) (err error) {
	var project autocode.Project
	projectDb := global.GVA_DB.Model(&autocode.Project{})
	// 获取对应的培训项目
	if err = projectDb.First(&project, outcome.ProjectId).Error; err != nil {
		return err
	}
	// todo 状态不是“进行中”的时候不允许更新收入和支出流水
	if project.Status != 1 {
		err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法修改流水", project.Name, utils.FormatProjectStatus(project.Status)))
		return err
	}
	project = UpdateProjectWhenDeleteOutcomeStream(outcome, project)
	// 更新了的培训项目添加到db
	if err = projectDb.Save(&project).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Unscoped().Delete(&outcome).Error
	return err
}

// DeleteOutcomeStreamByIds 批量删除OutcomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (outcomeService *OutcomeStreamService) DeleteOutcomeStreamByIds(ids request.IdsReq) (err error) {
	// todo 一条一条删除，动态更新流水对应的培训项目
	var outcomes []autocode.OutcomeStream
	if err = global.GVA_DB.Model(&[]autocode.OutcomeStream{}).Where("id in ?", ids.Ids).Find(&outcomes).Error; err != nil {
		return err
	}
	for _, outcome := range outcomes {
		var project autocode.Project
		// 找到待更新的培训项目
		if err = global.GVA_DB.Where("id = ?", outcome.ProjectId).First(&project).Error; err != nil {
			return err
		}
		// 状态不是“进行中”的时候不允许更新收入和支出流水
		if project.Status != 1 {
			err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法删除流水", project.Name, utils.FormatProjectStatus(project.Status)))
			return err
		}
		// 更新培训项目
		project = UpdateProjectWhenDeleteOutcomeStream(outcome, project)
		if err = global.GVA_DB.Unscoped().Delete(&outcome).Error; err != nil {
			return err
		}
		if err = global.GVA_DB.Save(&project).Error; err != nil {
			return err
		}
	}
	// err = global.GVA_DB.Delete(&[]autocode.IncomeStream{}, "id in ?", ids.Ids).Error
	return nil
}

// UpdateOutcomeStream 更新OutcomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (outcomeService *OutcomeStreamService) UpdateOutcomeStream(outcome autocode.OutcomeStream) (err error) {
	// todo 一条一条更新，动态更新流水对应的培训项目
	var project autocode.Project
	var outcomeBeforeUpdate autocode.OutcomeStream
	if err = global.GVA_DB.Where("id = ?", outcome.ProjectId).First(&project).Error; err != nil {
		return err
	}
	// todo 状态不是“进行中”的时候不允许更新收入和支出流水
	if project.Status != 1 {
		err = errors.New(fmt.Sprintf("该流水所属的项目%d处于[%d]状态，无法修改流水", project.Name, utils.FormatProjectStatus(project.Status)))
		return err
	}
	// 找到更新前的收入流水
	if err = global.GVA_DB.Where("id = ?", outcome.ID).First(&outcomeBeforeUpdate).Error; err != nil {
		return err
	}
	// 把更新前的收入流水传入该函数更新项目，就把更新前的流水叠加在项目上的金额全去掉了
	project = UpdateProjectWhenDeleteOutcomeStream(outcomeBeforeUpdate, project)
	// 然后把更新后的收入流水传进去更新项目，就正确更新的需要更新的项目,以及收入流水（收入流水更新以后，分成需要重新计算，因为入账金额可能变了）
	// 可以根据入账金额是否变化优化，但是懒得搞了
	project = UpdateProjectWhenCreateOutcomeStream(outcome, project)
	// 然后将正确更新的项目和流水保存到数据库里
	if err = global.GVA_DB.Save(&project).Error; err != nil {
		return err
	}
	err = global.GVA_DB.Save(&outcome).Error
	return err
}

// GetOutcomeStream 根据id获取OutcomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (outcomeService *OutcomeStreamService) GetOutcomeStream(id uint) (err error, outcome autocode.OutcomeStream) {
	err = global.GVA_DB.Where("id = ?", id).First(&outcome).Error
	return
}

// GetOutcomeStreamInfoList 分页获取OutcomeStream记录
// Author [piexlmax](https://github.com/piexlmax)
func (outcomeService *OutcomeStreamService) GetOutcomeStreamInfoList(info autoCodeReq.OutcomeStreamSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.OutcomeStream{})
	var outcomes []autocode.OutcomeStream
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProjectName != "" {
		db = db.Where("project_name LIKE ?", "%"+info.ProjectName+"%")
	}
	if info.OutcomeProjectCode != "" {
		db = db.Where("outcome_project_code LIKE ?", "%"+info.OutcomeProjectCode+"%")
	}
	for index, date := range info.CreatedDateRange {
		query := "created_date >= ?"
		if index == 1 {
			query = "created_date <= ?"
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
	err = db.Limit(limit).Offset(offset).Find(&outcomes).Error
	return err, outcomes, total
}
