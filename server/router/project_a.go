package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitProjectARouter(Router *gin.RouterGroup) {
	ProjectARouter := Router.Group("projectA").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ProjectARouter.POST("createProjectA", v1.CreateProjectA)   // 新建ProjectA
		ProjectARouter.DELETE("deleteProjectA", v1.DeleteProjectA) // 删除ProjectA
		ProjectARouter.DELETE("deleteProjectAByIds", v1.DeleteProjectAByIds) // 批量删除ProjectA
		ProjectARouter.PUT("updateProjectA", v1.UpdateProjectA)    // 更新ProjectA
		ProjectARouter.GET("findProjectA", v1.FindProjectA)        // 根据ID获取ProjectA
		ProjectARouter.GET("getProjectAList", v1.GetProjectAList)  // 获取ProjectA列表
	}
}
