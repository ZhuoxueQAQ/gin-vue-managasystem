package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OutcomeStreamRouter struct {
}

// InitOutcomeStreamRouter 初始化 OutcomeStream 路由信息
func (s *OutcomeStreamRouter) InitOutcomeStreamRouter(Router *gin.RouterGroup) {
	outcomeRouter := Router.Group("outcome").Use(middleware.OperationRecord())
	outcomeRouterWithoutRecord := Router.Group("outcome")
	var outcomeApi = v1.ApiGroupApp.AutoCodeApiGroup.OutcomeStreamApi
	{
		outcomeRouter.POST("createOutcomeStream", outcomeApi.CreateOutcomeStream)   // 新建OutcomeStream
		outcomeRouter.DELETE("deleteOutcomeStream", outcomeApi.DeleteOutcomeStream) // 删除OutcomeStream
		outcomeRouter.DELETE("deleteOutcomeStreamByIds", outcomeApi.DeleteOutcomeStreamByIds) // 批量删除OutcomeStream
		outcomeRouter.PUT("updateOutcomeStream", outcomeApi.UpdateOutcomeStream)    // 更新OutcomeStream
	}
	{
		outcomeRouterWithoutRecord.GET("findOutcomeStream", outcomeApi.FindOutcomeStream)        // 根据ID获取OutcomeStream
		outcomeRouterWithoutRecord.GET("getOutcomeStreamList", outcomeApi.GetOutcomeStreamList)  // 获取OutcomeStream列表
	}
}
