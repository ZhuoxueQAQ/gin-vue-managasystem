package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct {
}

// InitProjectRouter 初始化 Project 路由信息
func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("project").Use(middleware.OperationRecord())
	projectRouterWithoutRecord := Router.Group("project")
	var projectApi = v1.ApiGroupApp.AutoCodeApiGroup.ProjectApi
	{
		projectRouter.POST("createProject", projectApi.CreateProject)                     // 新建Project
		projectRouter.DELETE("deleteProject", projectApi.DeleteProject)                   // 删除Project
		projectRouter.DELETE("deleteProjectByIds", projectApi.DeleteProjectByIds)         // 批量删除Project
		projectRouter.DELETE("deleteProjectFileByIds", projectApi.DeleteProjectFileByIds) // 批量删除附件
		projectRouter.PUT("updateProject", projectApi.UpdateProject)                      // 更新Project
		projectRouter.POST("uploadProjectFileChunk", projectApi.UploadProjectChunk)       // 上传项目附件分片
		projectRouter.POST("mergeProjectFileChunk", projectApi.MergeProjectFileChunk)     // 合并项目附件分片
		projectRouter.POST("changeProjectStatus", projectApi.ChangeProjectStatus)         // 修改项目状态
	}
	{
		projectRouterWithoutRecord.GET("findProject", projectApi.FindProject)                           // 根据ID获取Project
		projectRouterWithoutRecord.GET("getProjectList", projectApi.GetProjectList)                     // 获取Project列表
		projectRouterWithoutRecord.GET("getProjectFileRecordList", projectApi.GetProjectFileRecordList) // 获取项目附件
		projectRouterWithoutRecord.GET("downloadFile", projectApi.DownloadFile)                         // 下载文件
	}
}
