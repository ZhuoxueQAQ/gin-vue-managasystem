// 自动生成模板IncomeStream
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"time"
)

// IncomeStream 结构体
// 如果含有time.Time 请自行import time包
type IncomeStream struct {
	global.GVA_MODEL
	ProjectId    int `json:"projectId" form:"projectId" gorm:"column:project_id;comment:所属项目ID;"`
	InvoiceCount int `json:"invoiceCount" form:"invoiceCount" gorm:"column:invoice_count;comment:发票数量;"`

	CreatedDate      time.Time `json:"createdDate" form:"createdDate" gorm:"column:created_date;comment:入账日期;"`
	InvoiceIssueDate time.Time `json:"invoiceIssueDate" form:"invoiceIssueDate" gorm:"column:invoice_issue_date;comment:发票开出日期;"`
	SplitAmountDate  time.Time `json:"splitAmountDate" form:"splitAmountDate" gorm:"column:split_amount_date;comment:分账日期;"`

	IncomeAmount  decimal.Decimal `json:"incomeAmount" form:"incomeAmount" gorm:"column:income_amount;comment:入账金额;"`
	InvoiceAmount decimal.Decimal `json:"invoiceAmount" form:"invoiceAmount" gorm:"column:invoice_amount;comment:发票金额;"`

	SRadio  decimal.Decimal `json:"sRadio" form:"sRadio" gorm:"column:s_radio;comment:学校管理费;"`
	DRadio  decimal.Decimal `json:"dRadio" form:"dRadio" gorm:"column:d_radio;comment:发展基金;"`
	WRadio  decimal.Decimal `json:"wRadio" form:"wRadio" gorm:"column:w_radio;comment:福利;"`
	CRadio  decimal.Decimal `json:"cRadio" form:"cRadio" gorm:"column:c_radio;comment:课酬;"`
	SAmount decimal.Decimal `json:"sAmount" form:"sAmount" gorm:"column:s_amount;comment:学校管理费;"`
	DAmount decimal.Decimal `json:"dAmount" form:"dAmount" gorm:"column:d_amount;comment:发展基金;"`
	WAmount decimal.Decimal `json:"wAmount" form:"wAmount" gorm:"column:w_amount;comment:福利;"`
	CAmount decimal.Decimal `json:"cAmount" form:"cAmount" gorm:"column:c_amount;comment:课酬;"`

	ProjectName     string `json:"projectName" form:"projectName" gorm:"column:project_name;comment:所属项目名称;size:200;"`
	Categories      string `json:"categories" form:"categories" gorm:"column:categories;comment:所属项目分类;size:50;"`
	IsOffset        string `json:"isOffset" form:"isOffset" gorm:"column:is_offset;comment:是否预先解票;size:200;"`
	IsBorrowed      string `json:"isBorrowed" form:"isBorrowed" gorm:"column:is_borrowed;comment:是否已冲账;size:200;"`
	PayUnit         string `json:"payUnit" form:"payUnit" gorm:"column:pay_unit;comment:缴费单位;size:200;"`
	Invoice         string `json:"invoice" form:"invoice" gorm:"column:invoice;comment:发票抬头;size:200;"`
	IncomeValidCode string `json:"incomeValidCode" form:"incomeValidCode" gorm:"column:income_valid_code;comment:入账凭证编码;size:200;"`
	InvoiceCode     string `json:"invoiceCode" form:"invoiceCode" gorm:"column:invoice_code;comment:invoiceCode;size:200;"`
	Remark          string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`

	Client        AmountOfIncomeOrOutComeStreamJSONArr `json:"client" form:"client" gorm:"column:client;comment:委托方;"`
	LandingAgency AmountOfIncomeOrOutComeStreamJSONArr `json:"landingAgency" form:"landingAgency" gorm:"column:landing_agency;comment:落地机构;"`
	Partner       AmountOfIncomeOrOutComeStreamJSONArr `json:"partner" form:"partner" gorm:"column:partner;comment:合作方;"`
	Pays          IncomeAndOutcomeObj                  `json:"pays" form:"pays" gorm:"column:pays;comment:一堆预算列表;"`
}

// TableName IncomeStream 表名
func (IncomeStream) TableName() string {
	return "income_stream"
}
