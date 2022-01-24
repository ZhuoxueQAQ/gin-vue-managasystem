// 自动生成模板ManageSystemSetting
package autocode

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type TableCol struct {
	ColType string `json:"colType"`
	Label   string `json:"label"`
	Prop    string `json:"prop"`
	Format  string `json:"format"`

	Show     bool   `json:"show"`
	ShowList []bool `json:"showList"`

	Width          int `json:"width"`
	MaxLengthIndex int `json:"maxLengthIndex"`
	MaxSubItemNum  int `json:"maxSubItemNum"`

	Sub *[]TableCol `json:"sub"`
}
type TableCols []TableCol

func (a TableCols) Value() (driver.Value, error) {
	b, err := json.Marshal(a)
	return string(b), err
}
func (a *TableCols) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), a)
}

// ManageSystemSetting 结构体
// 如果含有time.Time 请自行import time包
type ManageSystemSetting struct {
	global.GVA_MODEL
	ProjectTableCols       TableCols `json:"projectTableCols" form:"projectTableCol" gorm:"type:json;column:project_table_cols;comment:培训项目表格模板;"`
	IncomeStreamTableCols  TableCols `json:"incomeStreamTableCols" form:"incomeStreamTableCols" gorm:"type:json;column:income_stream_table_cols;comment:收入流水表格模板;"`
	OutcomeStreamTableCols TableCols `json:"outcomeStreamTableCols" form:"outcomeStreamTableCols" gorm:"type:json;column:outcome_stream_table_cols;comment:支出流水表格模板;"`
}

// TableName ManageSystemSetting 表名
func (ManageSystemSetting) TableName() string {
	return "manage_sys_setting"
}
