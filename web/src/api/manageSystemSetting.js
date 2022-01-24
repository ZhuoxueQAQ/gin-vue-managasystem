import service from '@/utils/request'

// @Tags ManageSystemSetting
// @Summary 创建ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ManageSystemSetting true "创建ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manageSystemSetting/createManageSystemSetting [post]
export const createManageSystemSetting = (data) => {
  return service({
    url: '/manageSystemSetting/createManageSystemSetting',
    method: 'post',
    data
  })
}

// @Tags ManageSystemSetting
// @Summary 删除ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ManageSystemSetting true "删除ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /manageSystemSetting/deleteManageSystemSetting [delete]
export const deleteManageSystemSetting = (data) => {
  return service({
    url: '/manageSystemSetting/deleteManageSystemSetting',
    method: 'delete',
    data
  })
}

// @Tags ManageSystemSetting
// @Summary 删除ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /manageSystemSetting/deleteManageSystemSetting [delete]
export const deleteManageSystemSettingByIds = (data) => {
  return service({
    url: '/manageSystemSetting/deleteManageSystemSettingByIds',
    method: 'delete',
    data
  })
}

// @Tags ManageSystemSetting
// @Summary 更新ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ManageSystemSetting true "更新ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /manageSystemSetting/updateManageSystemSetting [put]
export const updateManageSystemSetting = (data) => {
  return service({
    url: '/manageSystemSetting/updateManageSystemSetting',
    method: 'put',
    data
  })
}

// @Tags ManageSystemSetting
// @Summary 用id查询ManageSystemSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ManageSystemSetting true "用id查询ManageSystemSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /manageSystemSetting/findManageSystemSetting [get]
export const findManageSystemSetting = (params) => {
  return service({
    url: '/manageSystemSetting/findManageSystemSetting',
    method: 'get',
    params
  })
}

// @Tags ManageSystemSetting
// @Summary 分页获取ManageSystemSetting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ManageSystemSetting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /manageSystemSetting/getManageSystemSettingList [get]
export const getManageSystemSettingList = (params) => {
  return service({
    url: '/manageSystemSetting/getManageSystemSettingList',
    method: 'get',
    params
  })
}
