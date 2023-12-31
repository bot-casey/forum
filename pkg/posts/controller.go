package Posts

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/posts")
	routes.POST("/", h.AddPost)
	routes.GET("/:id", h.GetPost)
	routes.PUT("/:id", h.UpdatePost)
	routes.DELETE("/:id", h.DeletePost)
	routes.GET("/recommendations", h.GetRecommendations)
}
