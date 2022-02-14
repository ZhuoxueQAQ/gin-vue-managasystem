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
	"os"
)

type OutcomeStreamApi struct {
}

var outcomeService = service.ServiceGroupApp.AutoCodeServiceGroup.OutcomeStreamService

// CreateOutcomeStream 创建OutcomeStream
// @Tags OutcomeStream
// @Summary 创建OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.OutcomeStream true "创建OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/createOutcomeStream [post]
func (outcomeApi *OutcomeStreamApi) CreateOutcomeStream(c *gin.Context) {
	var outcome autocode.OutcomeStream
	_ = c.ShouldBindJSON(&outcome)
	if err := outcomeService.CreateOutcomeStream(outcome); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("创建失败:%d", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOutcomeStream 删除OutcomeStream
// @Tags OutcomeStream
// @Summary 删除OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.OutcomeStream true "删除OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outcome/deleteOutcomeStream [delete]
func (outcomeApi *OutcomeStreamApi) DeleteOutcomeStream(c *gin.Context) {
	var outcome autocode.OutcomeStream
	_ = c.ShouldBindJSON(&outcome)
	if err := outcomeService.DeleteOutcomeStream(outcome); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("删除失败:%d", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOutcomeStreamByIds 批量删除OutcomeStream
// @Tags OutcomeStream
// @Summary 批量删除OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /outcome/deleteOutcomeStreamByIds [delete]
func (outcomeApi *OutcomeStreamApi) DeleteOutcomeStreamByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := outcomeService.DeleteOutcomeStreamByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("批量删除失败:%d", err), c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOutcomeStream 更新OutcomeStream
// @Tags OutcomeStream
// @Summary 更新OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.OutcomeStream true "更新OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outcome/updateOutcomeStream [put]
func (outcomeApi *OutcomeStreamApi) UpdateOutcomeStream(c *gin.Context) {
	var outcome autocode.OutcomeStream
	_ = c.ShouldBindJSON(&outcome)
	if err := outcomeService.UpdateOutcomeStream(outcome); err != nil {
		global.GVA_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("更新失败:%d", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOutcomeStream 用id查询OutcomeStream
// @Tags OutcomeStream
// @Summary 用id查询OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.OutcomeStream true "用id查询OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outcome/findOutcomeStream [get]
func (outcomeApi *OutcomeStreamApi) FindOutcomeStream(c *gin.Context) {
	var outcome autocode.OutcomeStream
	_ = c.ShouldBindQuery(&outcome)
	if err, reoutcome := outcomeService.GetOutcomeStream(outcome.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reoutcome": reoutcome}, c)
	}
}

// GetOutcomeStreamList 分页获取OutcomeStream列表
// @Tags OutcomeStream
// @Summary 分页获取OutcomeStream列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.OutcomeStreamSearch true "分页获取OutcomeStream列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/getOutcomeStreamList [get]
func (outcomeApi *OutcomeStreamApi) GetOutcomeStreamList(c *gin.Context) {
	var info autocodeReq.OutcomeStreamSearch
	// 从get请求中获取searchInfo（json字符串），解码，然后转成json对象，绑定到info结构体
	searchInfo, _ := url.QueryUnescape(c.Query("searchInfo"))
	// 如果解码searchInfo失败，直接返回
	if err := json.Unmarshal([]byte(searchInfo), &info); err != nil {
		global.GVA_LOG.Error("查询JSON参数解析失败!", zap.Error(err))
		response.FailWithMessage("查询JSON参数解析失败", c)
		return
	}
	if err, list, total := outcomeService.GetOutcomeStreamInfoList(info, false); err != nil {
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

// ExportToExcel 导出已筛选的支出流水数据
// @Tags Project
// @Summary 导出已筛选的project数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProjectSearch true "导出已筛选的project数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/ExportToExcel [get]
func (outcomeApi *OutcomeStreamApi) ExportToExcel(c *gin.Context) {
	var info autocodeReq.OutcomeStreamSearch
	// 从get请求中获取searchInfo，解码，然后转成json对象，绑定到info结构体
	searchInfo, _ := url.QueryUnescape(c.Query("searchInfo"))
	// 如果解码失败，直接返回
	if err := json.Unmarshal([]byte(searchInfo), &info); err != nil {
		global.GVA_LOG.Error("查询JSON参数解析失败!", zap.Error(err))
		response.FailWithMessage("查询JSON参数解析失败", c)
		return
	}

	var path string

	if err, list, _ := outcomeService.GetOutcomeStreamInfoList(info, true); err != nil {
		global.GVA_LOG.Error("获取收入流水列表失败!", zap.Error(err))
		response.FailWithMessage("获取收入流水列表失败", c)
	} else {
		if err, path = outcomeService.ExportToExcel(list); err != nil {
			global.GVA_LOG.Error("生成excel文件失败!", zap.Error(err))
			response.FailWithMessage("生成excel文件失败", c)
		}
		//获取文件类型对应的http ContentType 类型
		c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
		c.Header("Content-Disposition", "导出数据.xlsx")
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(path)
	}

	defer func() {
		if err := os.Remove(path); err != nil {
			global.GVA_LOG.Error("移除临时生成的文件失败!", zap.Error(err))
		}
	}()
}
