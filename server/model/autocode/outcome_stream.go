// 自动生成模板OutcomeStream
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// OutcomeStream 结构体
// 如果含有time.Time 请自行import time包
type OutcomeStream struct {
	global.GVA_MODEL
	ProjectId   *int       `json:"projectId" form:"projectId" gorm:"column:project_id;comment:所属项目Id;"`
	CreatedDate *time.Time `json:"createdDate" form:"createdDate" gorm:"column:created_date;comment:支出日期;"`

	Categories         string `json:"categories" form:"categories" gorm:"column:categories;comment:所属项目分类;size:50;"`
	ProjectName        string `json:"projectName" form:"projectName" gorm:"column:project_name;comment:项目名称;size:200;"`
	OutcomeProjectCode string `json:"outcomeProjectCode" form:"outcomeProjectCode" gorm:"column:outcome_project_code;comment:支出项目码;size:100;"`
	Remark             string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`

	Client        AmountOfIncomeOrOutComeStreamJSONArr `json:"client" form:"client" gorm:"column:client;comment:委托方;"`
	LandingAgency AmountOfIncomeOrOutComeStreamJSONArr `json:"landingAgency" form:"landingAgency" gorm:"column:landing_agency;comment:落地机构;"`
	Partner       AmountOfIncomeOrOutComeStreamJSONArr `json:"partner" form:"partner" gorm:"column:partner;comment:技术方;"`
	Pays          IncomeAndOutcomeObj                  `json:"pays" form:"pays" gorm:"column:pays;comment:一堆支出;"`
}

// TableName OutcomeStream 表名
func (OutcomeStream) TableName() string {
	return "outcome_stream"
}
