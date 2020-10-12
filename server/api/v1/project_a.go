package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags ProjectA
// @Summary 创建ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "创建ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectA/createProjectA [post]
func CreateProjectA(c *gin.Context) {
	var projectA model.ProjectA
	_ = c.ShouldBindJSON(&projectA)
	err := service.CreateProjectA(projectA)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ProjectA
// @Summary 删除ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "删除ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectA/deleteProjectA [delete]
func DeleteProjectA(c *gin.Context) {
	var projectA model.ProjectA
	_ = c.ShouldBindJSON(&projectA)
	err := service.DeleteProjectA(projectA)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ProjectA
// @Summary 批量删除ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /projectA/deleteProjectAByIds [delete]
func DeleteProjectAByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	err := service.DeleteProjectAByIds(IDS)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ProjectA
// @Summary 更新ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "更新ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /projectA/updateProjectA [put]
func UpdateProjectA(c *gin.Context) {
	var projectA model.ProjectA
	_ = c.ShouldBindJSON(&projectA)
	err := service.UpdateProjectA(&projectA)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ProjectA
// @Summary 用id查询ProjectA
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProjectA true "用id查询ProjectA"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /projectA/findProjectA [get]
func FindProjectA(c *gin.Context) {
	var projectA model.ProjectA
	_ = c.ShouldBindQuery(&projectA)
	err, reprojectA := service.GetProjectA(projectA.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	} else {
		response.OkWithData(gin.H{"reprojectA": reprojectA}, c)
	}
}

// @Tags ProjectA
// @Summary 分页获取ProjectA列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ProjectASearch true "分页获取ProjectA列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /projectA/getProjectAList [get]
func GetProjectAList(c *gin.Context) {
	var pageInfo request.ProjectASearch
	_ = c.ShouldBindQuery(&pageInfo)
	err, list, total := service.GetProjectAInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}
