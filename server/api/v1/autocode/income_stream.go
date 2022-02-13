package autocode

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/url"
)

type IncomeStreamApi struct {
}

var IncomeService = service.ServiceGroupApp.AutoCodeServiceGroup.IncomeStreamService

// CreateIncomeStream 创建IncomeStream
// @Tags IncomeStream
// @Summary 创建IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IncomeStream true "创建IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Income/createIncomeStream [post]
func (IncomeApi *IncomeStreamApi) CreateIncomeStream(c *gin.Context) {
	var Income autocode.IncomeStream
	_ = c.ShouldBindJSON(&Income)
	if err := IncomeService.CreateIncomeStream(Income); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIncomeStream 删除IncomeStream
// @Tags IncomeStream
// @Summary 删除IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IncomeStream true "删除IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Income/deleteIncomeStream [delete]
func (IncomeApi *IncomeStreamApi) DeleteIncomeStream(c *gin.Context) {
	var Income autocode.IncomeStream
	_ = c.ShouldBindJSON(&Income)
	if err := IncomeService.DeleteIncomeStream(Income); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIncomeStreamByIds 批量删除IncomeStream
// @Tags IncomeStream
// @Summary 批量删除IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Income/deleteIncomeStreamByIds [delete]
func (IncomeApi *IncomeStreamApi) DeleteIncomeStreamByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := IncomeService.DeleteIncomeStreamByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("批量删除失败:%d", err.Error()), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIncomeStream 更新IncomeStream
// @Tags IncomeStream
// @Summary 更新IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.IncomeStream true "更新IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Income/updateIncomeStream [put]
func (IncomeApi *IncomeStreamApi) UpdateIncomeStream(c *gin.Context) {
	var Income autocode.IncomeStream
	_ = c.ShouldBindJSON(&Income)
	if err := IncomeService.UpdateIncomeStream(Income); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("更新失败:%d", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIncomeStream 用id查询IncomeStream
// @Tags IncomeStream
// @Summary 用id查询IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.IncomeStream true "用id查询IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Income/findIncomeStream [get]
func (IncomeApi *IncomeStreamApi) FindIncomeStream(c *gin.Context) {
	var Income autocode.IncomeStream
	_ = c.ShouldBindQuery(&Income)
	if err, reIncome := IncomeService.GetIncomeStream(Income.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reIncome": reIncome}, c)
	}
}

// GetIncomeStreamList 分页获取IncomeStream列表
// @Tags IncomeStream
// @Summary 分页获取IncomeStream列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.IncomeStreamSearch true "分页获取IncomeStream列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Income/getIncomeStreamList [get]
func (IncomeApi *IncomeStreamApi) GetIncomeStreamList(c *gin.Context) {
	var info autocodeReq.IncomeStreamSearch
	// 从get请求中获取searchInfo，解码，然后转成json对象，绑定到info结构体
	searchInfo, _ := url.QueryUnescape(c.Query("searchInfo"))
	// 如果解码失败，直接返回
	if err := json.Unmarshal([]byte(searchInfo), &info); err != nil {
		global.GVA_LOG.Error("查询JSON参数解析失败!", zap.Error(err))
		response.FailWithMessage("查询JSON参数解析失败", c)
		return
	}
	if err, list, total := IncomeService.GetIncomeStreamInfoList(info); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     info.Page,
			PageSize: info.PageSize,
		}, "获取成功", c)
	}
}
