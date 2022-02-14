import service from '@/utils/request'

// @Tags IncomeStream
// @Summary 创建IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeStream true "创建IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Income/createIncomeStream [post]
export const createIncomeStream = (data) => {
  return service({
    url: '/Income/createIncomeStream',
    method: 'post',
    data
  })
}

// @Tags IncomeStream
// @Summary 删除IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeStream true "删除IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Income/deleteIncomeStream [delete]
export const deleteIncomeStream = (data) => {
  return service({
    url: '/Income/deleteIncomeStream',
    method: 'delete',
    data
  })
}

// @Tags IncomeStream
// @Summary 删除IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Income/deleteIncomeStream [delete]
export const deleteIncomeStreamByIds = (data) => {
  return service({
    url: '/Income/deleteIncomeStreamByIds',
    method: 'delete',
    data
  })
}

// @Tags IncomeStream
// @Summary 更新IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IncomeStream true "更新IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Income/updateIncomeStream [put]
export const updateIncomeStream = (data) => {
  return service({
    url: '/Income/updateIncomeStream',
    method: 'put',
    data
  })
}

// @Tags IncomeStream
// @Summary 用id查询IncomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IncomeStream true "用id查询IncomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Income/findIncomeStream [get]
export const findIncomeStream = (params) => {
  return service({
    url: '/Income/findIncomeStream',
    method: 'get',
    params
  })
}

// @Tags IncomeStream
// @Summary 分页获取IncomeStream列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IncomeStream列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Income/getIncomeStreamList [get]
export const getIncomeStreamList = (params) => {
  return service({
    url: '/Income/getIncomeStreamList',
    method: 'get',
    params
  })
}

// @Tags IncomeStream
// @Summary 导出收入流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "导出收入流水"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Income/exportToExcel [get]
export const exportToExcel = (params) => {
  return service({
    url: '/Income/exportToExcel',
    method: 'get',
    params,
    responseType: 'blob' // 制定响应类型
  })
}

