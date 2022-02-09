import service from '@/utils/request'

// @Tags OutcomeStream
// @Summary 创建OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutcomeStream true "创建OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/createOutcomeStream [post]
export const createOutcomeStream = (data) => {
  return service({
    url: '/outcome/createOutcomeStream',
    method: 'post',
    data
  })
}

// @Tags OutcomeStream
// @Summary 删除OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutcomeStream true "删除OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outcome/deleteOutcomeStream [delete]
export const deleteOutcomeStream = (data) => {
  return service({
    url: '/outcome/deleteOutcomeStream',
    method: 'delete',
    data
  })
}

// @Tags OutcomeStream
// @Summary 删除OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /outcome/deleteOutcomeStream [delete]
export const deleteOutcomeStreamByIds = (data) => {
  return service({
    url: '/outcome/deleteOutcomeStreamByIds',
    method: 'delete',
    data
  })
}

// @Tags OutcomeStream
// @Summary 更新OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.OutcomeStream true "更新OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /outcome/updateOutcomeStream [put]
export const updateOutcomeStream = (data) => {
  return service({
    url: '/outcome/updateOutcomeStream',
    method: 'put',
    data
  })
}

// @Tags OutcomeStream
// @Summary 用id查询OutcomeStream
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.OutcomeStream true "用id查询OutcomeStream"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /outcome/findOutcomeStream [get]
export const findOutcomeStream = (params) => {
  return service({
    url: '/outcome/findOutcomeStream',
    method: 'get',
    params
  })
}

// @Tags OutcomeStream
// @Summary 分页获取OutcomeStream列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取OutcomeStream列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /outcome/getOutcomeStreamList [get]
export const getOutcomeStreamList = (params) => {
  return service({
    url: '/outcome/getOutcomeStreamList',
    method: 'get',
    params
  })
}
