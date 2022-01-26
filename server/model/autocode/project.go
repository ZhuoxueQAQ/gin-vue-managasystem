// 自动生成模板Project
package autocode

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"time"
)

// IncomeAndOutcomeObj 结构体是培训项目的收入|支出金额对象。
type IncomeAndOutcomeObj struct {
	Pg decimal.Decimal `json:"pg"`
	Ph decimal.Decimal `json:"ph"`
	Pi decimal.Decimal `json:"pi"`
	Pj decimal.Decimal `json:"pj"`
	Pk decimal.Decimal `json:"pk"`
	Pl decimal.Decimal `json:"pl"`
	Pm decimal.Decimal `json:"pm"`
	Pn decimal.Decimal `json:"pn"`
	Po decimal.Decimal `json:"po"`
	Pp decimal.Decimal `json:"pp"`
	Pq decimal.Decimal `json:"pq"`
	Pr decimal.Decimal `json:"pr"`
	Ps decimal.Decimal `json:"ps"`
	Pt decimal.Decimal `json:"pt"`
}

// IncomeAndOutcomeJSONArr 类型是一个数组，第一个元素为预算对象，第二个元素为支出对象
type IncomeAndOutcomeJSONArr []IncomeAndOutcomeObj

func (i IncomeAndOutcomeJSONArr) Value() (driver.Value, error) {
	b, err := json.Marshal(i)
	return string(b), err
}
func (i *IncomeAndOutcomeJSONArr) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), i)
}

// AmountObj 结构体是一个金额对象
type AmountObj struct {
	Name   string          `json:"name"`
	Radio  decimal.Decimal `json:"radio"`
	Amount decimal.Decimal `json:"amount"`
}

// AmountJSONArr 结构体是一个金额对象json数组
type AmountJSONArr []AmountObj

func (a AmountJSONArr) Value() (driver.Value, error) {
	b, err := json.Marshal(a)
	return string(b), err
}
func (a *AmountJSONArr) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

// Project 结构体
// 如果含有time.Time 请自行import time包
type Project struct {
	global.GVA_MODEL
	Name           string `json:"name" form:"name" gorm:"column:name;comment:项目名称;size:100;"`
	Categories     string `json:"categories" form:"categories" gorm:"column:categories;comment:项目类别;size:20;"`
	Area           string `json:"area" form:"area" gorm:"column:area;comment:项目所属地;size:50;"`
	ChargeStandard string `json:"chargeStandard" form:"chargeStandard" gorm:"column:charge_standard;comment:收费标准;size:500;"`
	Manager        string `json:"manager" form:"manager" gorm:"column:manager;comment:项目负责人;size:20;"`
	ProjectCode    string `json:"projectCode" form:"projectCode" gorm:"column:project_code;comment:项目码;size:50;"`
	TrainMode      string `json:"trainMode" form:"trainMode" gorm:"column:train_mode;comment:培训模式;size:20;"`
	ContractNum    string `json:"contractNum" form:"contractNum" gorm:"column:contract_num;comment:合同编号;size:200;"`
	Remark         string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:200;"`

	ProjectAmount decimal.Decimal `json:"projectAmount" sql:"type:decimal(12,2);" form:"projectAmount" gorm:"column:project_amount;comment:项目应收费用;"`
	PaidAmount    decimal.Decimal `json:"paidAmount" sql:"type:decimal(12,2);" form:"paidAmount" gorm:"column:paid_amount;comment:已到账费用;"`
	UnpaidAmount  decimal.Decimal `json:"unpaidAmount" sql:"type:decimal(12,2);" form:"unpaidAmount" gorm:"column:unpaid_amount;comment:未到账金额;"`

	SRadio  decimal.Decimal `json:"sRadio" sql:"type:decimal(7,6);" form:"sRadio" gorm:"column:s_radio;comment:学校管理费比例;"`
	DRadio  decimal.Decimal `json:"dRadio" sql:"type:decimal(7,6);" form:"dRadio" gorm:"column:d_radio;comment:发展基金比例;"`
	WRadio  decimal.Decimal `json:"wRadio" sql:"type:decimal(7,6);" form:"wRadio" gorm:"column:w_radio;comment:福利比例;"`
	CRadio  decimal.Decimal `json:"cRadio" sql:"type:decimal(7,6);" form:"cRadio" gorm:"column:c_radio;comment:课酬比例;"`
	SAmount decimal.Decimal `json:"sAmount" sql:"type:decimal(12,2);" form:"SAmount" gorm:"column:s_amount;comment:学校管理费;"`
	DAmount decimal.Decimal `json:"dAmount" sql:"type:decimal(12,2);" form:"dAmount" gorm:"column:d_amount;comment:发展基金;"`
	WAmount decimal.Decimal `json:"wAmount" sql:"type:decimal(12,2);" form:"wAmount" gorm:"column:w_amount;comment:福利;"`
	CAmount decimal.Decimal `json:"cAmount" sql:"type:decimal(12,2);" form:"cAmount" gorm:"column:c_amount;comment:课酬;"`

	TrainNumOfPerson int `json:"trainNumOfPerson" form:"trainNumOfPerson" gorm:"column:train_num_of_person;comment:培训人数;size:19;"`
	TrainTime        int `json:"trainTime" form:"trainTime" gorm:"column:train_time;comment:培训学时数;size:19;"`

	Client        *AmountJSONArr `json:"client" form:"client" gorm:"type:json;column:client;comment:委托方;"`
	LandingAgency *AmountJSONArr `json:"landingAgency" form:"landingAgency" gorm:"type:json;column:landing_agency;comment:落地机构;"`
	Partner       *AmountJSONArr `json:"partner" form:"partner" gorm:"type:json;column:partner;comment:技术方;"`

	IncomeAndOutcome *IncomeAndOutcomeJSONArr `json:"incomeAndOutcome" validate:"required" form:"incomeAndOutcome" gorm:"type:json;column:income_and_outcome;comment:预算和支出;"`

	CreatedDate       time.Time `json:"createdDate" form:"createdDate" validate:"required" gorm:"type:date;column:created_date;comment:备案申请日期;"`
	ContractStartDate time.Time `json:"contractStartDate" form:"contractStartDate" gorm:"type:date;column:contract_start_date;comment:合同开始时间;"`
	ContractEndDate   time.Time `json:"contractEndDate" form:"contractEndDate" gorm:"type:date;column:contract_end_date;comment:合同结束时间;"`
	TrainStartDate    time.Time `json:"trainStartDate" form:"trainStartDate" gorm:"type:date;column:train_start_date;comment:培训开始时间;"`
	TrainEndDate      time.Time `json:"trainEndDate" form:"trainEndDate" gorm:"type:date;column:train_end_date;comment:培训结束时间;"`
}

// TableName Project 表名
func (Project) TableName() string {
	return "project"
}
