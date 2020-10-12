import service from '@/utils/request'

// @Tags ProjectA
// @Summary 创建ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "创建ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectA/createProjectA [post]
export const createProjectA = (data) => {
     return service({
         url: "/projectA/createProjectA",
         method: 'post',
         data
     })
 }


// @Tags ProjectA
// @Summary 删除ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "删除ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectA/deleteProjectA [delete]
 export const deleteProjectA = (data) => {
     return service({
         url: "/projectA/deleteProjectA",
         method: 'delete',
         data
     })
 }

// @Tags ProjectA
// @Summary 删除ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectA/deleteProjectA [delete]
 export const deleteProjectAByIds = (data) => {
     return service({
         url: "/projectA/deleteProjectAByIds",
         method: 'delete',
         data
     })
 }

// @Tags ProjectA
// @Summary 更新ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "更新ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /projectA/updateProjectA [put]
 export const updateProjectA = (data) => {
     return service({
         url: "/projectA/updateProjectA",
         method: 'put',
         data
     })
 }


// @Tags ProjectA
// @Summary 用id查询ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "用id查询ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /projectA/findProjectA [get]
 export const findProjectA = (params) => {
     return service({
         url: "/projectA/findProjectA",
         method: 'get',
         params
     })
 }


// @Tags ProjectA
// @Summary 分页获取ProjectA列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ProjectA列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectA/getProjectAList [get]
 export const getProjectAList = (params) => {
     return service({
         url: "/projectA/getProjectAList",
         method: 'get',
         params
     })
 }