import service from '@/utils/request'

// @Tags Project
// @Summary 创建Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "创建Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/createProject [post]
export const createProject = (data) => {
  return service({
    url: '/project/createProject',
    method: 'post',
    data
  })
}

// @Tags Project
// @Summary 删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProject [delete]
export const deleteProject = (data) => {
  return service({
    url: '/project/deleteProject',
    method: 'delete',
    data
  })
}

// @Tags Project
// @Summary 删除Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /project/deleteProject [delete]
export const deleteProjectByIds = (data) => {
  return service({
    url: '/project/deleteProjectByIds',
    method: 'delete',
    data
  })
}

// @Tags Project
// @Summary 更新Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Project true "更新Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /project/updateProject [put]
export const updateProject = (data) => {
  return service({
    url: '/project/updateProject',
    method: 'put',
    data
  })
}

// @Tags Project
// @Summary 用id查询Project
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Project true "用id查询Project"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/findProject [get]
export const findProject = (params) => {
  return service({
    url: '/project/findProject',
    method: 'get',
    params
  })
}

// @Tags Project
// @Summary 分页获取Project列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Project列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /project/getProjectList [get]
export const getProjectList = (params) => {
  return service({
    url: '/project/getProjectList',
    method: 'get',
    params
  })
}

// @Tags Project
// @Summary 上传项目附件（分片）
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Project true "上传项目附件（分片）"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/uploadProjectFile [post]
export const uploadProjectChunk = (data) => {
  return service({
    url: '/project/uploadProjectFileChunk',
    method: 'post',
    data
  })
}

// @Tags Project
// @Summary 合并分片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Project true "合并分片"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /project/uploadProjectFile [post]
export const mergeProjectFileChunk = (data) => {
  return service({
    url: '/project/mergeProjectFileChunk',
    method: 'post',
    data
  })
}
