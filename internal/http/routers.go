package http

import (
	"github.com/askaroe/hexago/internal/domain/service"
	"github.com/askaroe/jsonlog"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(logger *jsonlog.Logger, services *service.Builder) (*Router, error) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	handler := newHandler(logger, services)

	// swagger settings
	//router.GET("/swagger/*any")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	rg := router.Group("/api/v1")
	{
		user := rg.Group("/users")
		{
			user.GET("/:id", handler.user.GetById)
			user.GET("/", handler.user.GetAll)
		}
	}

	return &Router{router}, nil
}
