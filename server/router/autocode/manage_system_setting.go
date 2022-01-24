package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ManageSystemSettingRouter struct {
}

// InitManageSystemSettingRouter 初始化 ManageSystemSetting 路由信息
func (s *ManageSystemSettingRouter) InitManageSystemSettingRouter(Router *gin.RouterGroup) {
	manageSystemSettingRouter := Router.Group("manageSystemSetting").Use(middleware.OperationRecord())
	manageSystemSettingRouterWithoutRecord := Router.Group("manageSystemSetting")
	var manageSystemSettingApi = v1.ApiGroupApp.AutoCodeApiGroup.ManageSystemSettingApi
	{
		manageSystemSettingRouter.POST("createManageSystemSetting", manageSystemSettingApi.CreateManageSystemSetting)   // 新建ManageSystemSetting
		manageSystemSettingRouter.DELETE("deleteManageSystemSetting", manageSystemSettingApi.DeleteManageSystemSetting) // 删除ManageSystemSetting
		manageSystemSettingRouter.DELETE("deleteManageSystemSettingByIds", manageSystemSettingApi.DeleteManageSystemSettingByIds) // 批量删除ManageSystemSetting
		manageSystemSettingRouter.PUT("updateManageSystemSetting", manageSystemSettingApi.UpdateManageSystemSetting)    // 更新ManageSystemSetting
	}
	{
		manageSystemSettingRouterWithoutRecord.GET("findManageSystemSetting", manageSystemSettingApi.FindManageSystemSetting)        // 根据ID获取ManageSystemSetting
		manageSystemSettingRouterWithoutRecord.GET("getManageSystemSettingList", manageSystemSettingApi.GetManageSystemSettingList)  // 获取ManageSystemSetting列表
	}
}
