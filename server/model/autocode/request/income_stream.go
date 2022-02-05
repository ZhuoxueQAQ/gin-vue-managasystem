package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type IncomeStreamSearch struct {
	autocode.IncomeStream
	request.PageInfo
	CreatedDateRange     []time.Time `json:"createdDateRange"`
	SplitAmountDateRange []time.Time `json:"splitAmountDateRange"`
}
