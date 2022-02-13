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
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type ProjectApi struct {
}

var projectService = service.ServiceGroupApp.AutoCodeServiceGroup.ProjectService

// CreateProject 创建Project
// @Tags Project
// @Summary 创建Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "创建Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/createProject [post]
func (projectApi *ProjectApi) CreateProject(c *gin.Context) {
	var project autocode.Project
	_ = c.ShouldBindJSON(&project)
	if err := projectService.CreateProject(project); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProject 删除Project
// @Tags Project
// @Summary 删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProject [delete]
func (projectApi *ProjectApi) DeleteProject(c *gin.Context) {
	var project autocode.Project
	_ = c.ShouldBindJSON(&project)
	if err := projectService.DeleteProject(project); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProjectByIds 批量删除Project
// @Tags Project
// @Summary 批量删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /project/deleteProjectByIds [delete]
func (projectApi *ProjectApi) DeleteProjectByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := projectService.DeleteProjectByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// todo 添加一个修改状态的api。并且状态大于1以后不能在设置成0（立项状态）防止修改项目信息

// UpdateProject 更新Project
// @Tags Project
// @Summary 更新Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "更新Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project/updateProject [put]
func (projectApi *ProjectApi) UpdateProject(c *gin.Context) {
	var project autocode.Project
	_ = c.ShouldBindJSON(&project)
	// todo 检查状态，如果进行中，不准更新项目中一切会影响到收入和支出流水的信息

	if err := projectService.UpdateProject(project); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProject 用id查询Project
// @Tags Project
// @Summary 用id查询Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Project true "用id查询Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/findProject [get]
func (projectApi *ProjectApi) FindProject(c *gin.Context) {
	var project autocode.Project
	_ = c.ShouldBindQuery(&project)
	if err, reproject := projectService.GetProject(project.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproject": reproject}, c)
	}
}

// GetProjectList 分页获取Project列表
// @Tags Project
// @Summary 分页获取Project列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProjectSearch true "分页获取Project列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/getProjectList [get]
func (projectApi *ProjectApi) GetProjectList(c *gin.Context) {
	var info autocodeReq.ProjectSearch
	// 从get请求中获取searchInfo，解码，然后转成json对象，绑定到info结构体
	searchInfo, _ := url.QueryUnescape(c.Query("searchInfo"))
	// 如果解码失败，直接返回
	if err := json.Unmarshal([]byte(searchInfo), &info); err != nil {
		global.GVA_LOG.Error("查询JSON参数解析失败!", zap.Error(err))
		response.FailWithMessage("查询JSON参数解析失败", c)
		return
	}
	if err, list, total := projectService.GetProjectInfoList(info, false); err != nil {
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

// uploadProjectFile 分片上传附件
// @Tags Project
// @Summary 分片上传附件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "分片上传附件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/uploadProjectFile [post]

func (projectApi *ProjectApi) UploadProjectChunk(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber"))
	// chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	_, FileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := FileHeader.Open()
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Error(err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	cen, _ := ioutil.ReadAll(f)
	if !utils.CheckMd5(cen, chunkMd5) {
		global.GVA_LOG.Error("检查分片md5失败!", zap.Error(err))
		response.FailWithMessage("检查分片md5失败", c)
		return
	}
	err, _ = utils.MakeFileChunk(cen, fileName, chunkNumber, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("创建分片文件失败!", zap.Error(err))
		response.FailWithMessage("创建分片文件失败", c)
		return
	}
	response.OkWithMessage("切片创建成功", c)

}

func (projectApi *ProjectApi) MergeProjectFileChunk(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	projectID := c.Query("projectID")
	fileTypeID := c.Query("fileTypeID")
	// chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))
	err, _ := utils.MergeChunk(fileName, fileMd5, projectID, fileTypeID)
	if err != nil {
		global.GVA_LOG.Error("合并分片文件失败!", zap.Error(err))
		response.FailWithMessage("合并分片文件失败", c)
		return
	}
	content, err := ioutil.ReadFile(filepath.Join("./fileDir", projectID, fileTypeID, fileName))
	if err != nil {
		response.FailWithMessage("读取已上传文件失败！", c)
		return
	}
	if !utils.CheckMd5(content, fileMd5) {
		response.FailWithMessage("校验文件md5失败！", c)
		return
	}
	// todo 添加上传记录
	fileRecord := autocode.ProjectFileRecord{FileMD5: fileMd5, FileName: fileName, FileTypeID: fileTypeID, ProjectID: projectID}
	if err := projectService.CreateUploadProjectFileRecord(fileRecord); err != nil {
		response.FailWithMessage("创建文件上传记录失败！", c)
		return
	}
	response.OkWithMessage("合并分片成功", c)
}

// GetProjectFileList
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小，projectID, fileTypeID"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router /fileUploadAndDownload/getFileList [get]
func (projectApi *ProjectApi) GetProjectFileRecordList(c *gin.Context) {
	var info autocodeReq.ProjectFileRecordSearch
	// bind query
	_ = c.ShouldBindQuery(&info)

	err, list, total := projectService.GetProjectFileRecordInfoList(info)
	if err != nil {
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

// DeleteProjectByIds 批量删除ProjectFileRecord
// @Tags Project
// @Summary 批量删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /project/deleteProjectByIds [delete]
func (projectApi *ProjectApi) DeleteProjectFileByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := projectService.DeleteProjectFileByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

//DownloadFile 根据文件路径读取返回流文件 参数url
func (projectApi *ProjectApi) DownloadFile(c *gin.Context) {
	var record autocode.ProjectFileRecord
	// bind query
	_ = c.ShouldBindQuery(&record)
	path := utils.GetProjectFilePath(record)
	//获取文件类型对应的http ContentType 类型
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", record.FileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path)
}

// ChangeProjectStatus 修改project的状态
// @Tags Project
// @Summary 修改project的状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Project true "修改project的状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /project/createProject [post]
func (projectApi *ProjectApi) ChangeProjectStatus(c *gin.Context) {
	var projectStatus autocode.ProjectStatus
	_ = c.ShouldBindJSON(&projectStatus)
	if err := projectService.ChangeProjectStatus(projectStatus); err != nil {
		global.GVA_LOG.Error("更新项目状态失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// ExportToExcel 导出已筛选的project数据
// @Tags Project
// @Summary 导出已筛选的project数据
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ProjectSearch true "导出已筛选的project数据"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/ExportToExcel [get]
func (projectApi *ProjectApi) ExportToExcel(c *gin.Context) {
	var info autocodeReq.ProjectSearch
	// 从get请求中获取searchInfo，解码，然后转成json对象，绑定到info结构体
	searchInfo, _ := url.QueryUnescape(c.Query("searchInfo"))
	// 如果解码失败，直接返回
	if err := json.Unmarshal([]byte(searchInfo), &info); err != nil {
		global.GVA_LOG.Error("查询JSON参数解析失败!", zap.Error(err))
		response.FailWithMessage("查询JSON参数解析失败", c)
		return
	}

	var path string

	if err, projects, _ := projectService.GetProjectInfoList(info, true); err != nil {
		global.GVA_LOG.Error("获取project列表失败!", zap.Error(err))
		response.FailWithMessage("获取project列表失败", c)
	} else {
		if err, path = projectService.ExportToExcel(projects); err != nil {
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
