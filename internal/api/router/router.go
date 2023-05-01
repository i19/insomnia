package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"insomnia/internal/api/handler"
	"insomnia/internal/api/middleware"
)

func Get() *gin.Engine {
	router := gin.Default()

	ctlGroup := router.Group("/api/linting_rule")
	ctlGroup.Use(middleware.GetGrantMiddleware)
	{
		ctlGroup.POST("/:project_id", middleware.IsAdminForThisProject, handler.PostByProjectID)
		ctlGroup.POST("/:project_id/validate", middleware.IsGrantedForThisProject, handler.Validate)
	}
	return router
}

func Run(port int) {
	Get().Run(fmt.Sprintf(":%d", port))
}
