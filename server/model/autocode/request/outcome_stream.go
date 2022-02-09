package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type OutcomeStreamSearch struct {
	autocode.OutcomeStream
	request.PageInfo
	CreatedDateRange []time.Time `json:"createdDateRange"`
}
