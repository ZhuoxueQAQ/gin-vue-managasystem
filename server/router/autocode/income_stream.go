package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IncomeStreamRouter struct {
}

// InitIncomeStreamRouter 初始化 IncomeStream 路由信息
func (s *IncomeStreamRouter) InitIncomeStreamRouter(Router *gin.RouterGroup) {
	IncomeRouter := Router.Group("Income").Use(middleware.OperationRecord())
	IncomeRouterWithoutRecord := Router.Group("Income")
	var IncomeApi = v1.ApiGroupApp.AutoCodeApiGroup.IncomeStreamApi
	{
		IncomeRouter.POST("createIncomeStream", IncomeApi.CreateIncomeStream)             // 新建IncomeStream
		IncomeRouter.DELETE("deleteIncomeStream", IncomeApi.DeleteIncomeStream)           // 删除IncomeStream
		IncomeRouter.DELETE("deleteIncomeStreamByIds", IncomeApi.DeleteIncomeStreamByIds) // 批量删除IncomeStream
		IncomeRouter.PUT("updateIncomeStream", IncomeApi.UpdateIncomeStream)              // 更新IncomeStream
	}
	{
		IncomeRouterWithoutRecord.GET("findIncomeStream", IncomeApi.FindIncomeStream)       // 根据ID获取IncomeStream
		IncomeRouterWithoutRecord.GET("getIncomeStreamList", IncomeApi.GetIncomeStreamList) // 获取IncomeStream列表
		IncomeRouterWithoutRecord.GET("exportToExcel", IncomeApi.ExportToExcel)             // 导出incomeStreams
	}
}
