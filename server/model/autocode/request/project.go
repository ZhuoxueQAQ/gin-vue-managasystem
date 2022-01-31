package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProjectSearch struct {
	autocode.Project
	request.PageInfo
	// 这里是为了日期范围筛选添加的几个日期数组字段
	CreatedDateRange       []time.Time `json:"createdDateRange"`
	TrainStartDateRange    []time.Time `json:"trainStartDateRange"`
	ContractStartDateRange []time.Time `json:"contractStartDateRange"`
}

type ProjectFileRecordSearch struct {
	autocode.ProjectFileRecord
	request.PageInfo
}
