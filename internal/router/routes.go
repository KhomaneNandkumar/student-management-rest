package router

import (
	"example.com/student-management/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("api/v1/health", handlers.Health)
	router.NoRoute(handlers.NoRoute)

	router.GET("api/v1/student", handlers.GetAll)
	router.GET("api/v1/student/:id", handlers.GetStudent)
	router.POST("api/v1/student", handlers.AddStudent)
	router.PUT("api/v1/student/:id", handlers.UpdateStudent)
	router.DELETE("api/v1/student/:id", handlers.DeleteStudent)
}
