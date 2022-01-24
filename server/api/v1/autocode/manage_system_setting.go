package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ManageSystemSettingApi struct {
}

var manageSystemSettingService = service.ServiceGroupApp.AutoCodeServiceGroup.ManageSystemSettingService


// CreateManageSystemSetting 创建ManageSystemSetting
// @Tags ManageSystemSetting
// @Summary 创建ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ManageSystemSetting true "创建ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manageSystemSetting/createManageSystemSetting [post]
func (manageSystemSettingApi *ManageSystemSettingApi) CreateManageSystemSetting(c *gin.Context) {
	var manageSystemSetting autocode.ManageSystemSetting
	_ = c.ShouldBindJSON(&manageSystemSetting)
	if err := manageSystemSettingService.CreateManageSystemSetting(manageSystemSetting); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteManageSystemSetting 删除ManageSystemSetting
// @Tags ManageSystemSetting
// @Summary 删除ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ManageSystemSetting true "删除ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /manageSystemSetting/deleteManageSystemSetting [delete]
func (manageSystemSettingApi *ManageSystemSettingApi) DeleteManageSystemSetting(c *gin.Context) {
	var manageSystemSetting autocode.ManageSystemSetting
	_ = c.ShouldBindJSON(&manageSystemSetting)
	if err := manageSystemSettingService.DeleteManageSystemSetting(manageSystemSetting); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteManageSystemSettingByIds 批量删除ManageSystemSetting
// @Tags ManageSystemSetting
// @Summary 批量删除ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /manageSystemSetting/deleteManageSystemSettingByIds [delete]
func (manageSystemSettingApi *ManageSystemSettingApi) DeleteManageSystemSettingByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := manageSystemSettingService.DeleteManageSystemSettingByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateManageSystemSetting 更新ManageSystemSetting
// @Tags ManageSystemSetting
// @Summary 更新ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.ManageSystemSetting true "更新ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /manageSystemSetting/updateManageSystemSetting [put]
func (manageSystemSettingApi *ManageSystemSettingApi) UpdateManageSystemSetting(c *gin.Context) {
	var manageSystemSetting autocode.ManageSystemSetting
	_ = c.ShouldBindJSON(&manageSystemSetting)
	if err := manageSystemSettingService.UpdateManageSystemSetting(manageSystemSetting); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindManageSystemSetting 用id查询ManageSystemSetting
// @Tags ManageSystemSetting
// @Summary 用id查询ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.ManageSystemSetting true "用id查询ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /manageSystemSetting/findManageSystemSetting [get]
func (manageSystemSettingApi *ManageSystemSettingApi) FindManageSystemSetting(c *gin.Context) {
	var manageSystemSetting autocode.ManageSystemSetting
	_ = c.ShouldBindQuery(&manageSystemSetting)
	if err, remanageSystemSetting := manageSystemSettingService.GetManageSystemSetting(manageSystemSetting.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remanageSystemSetting": remanageSystemSetting}, c)
	}
}

// GetManageSystemSettingList 分页获取ManageSystemSetting列表
// @Tags ManageSystemSetting
// @Summary 分页获取ManageSystemSetting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ManageSystemSettingSearch true "分页获取ManageSystemSetting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manageSystemSetting/getManageSystemSettingList [get]
func (manageSystemSettingApi *ManageSystemSettingApi) GetManageSystemSettingList(c *gin.Context) {
	var pageInfo autocodeReq.ManageSystemSettingSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := manageSystemSettingService.GetManageSystemSettingInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
