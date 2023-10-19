package Posts

import (
	"net/http"
	Post "redditClone/pkg/common/models/posts"

	"github.com/gin-gonic/gin"
)

func (h handler) GetPost(c *gin.Context) {
	Post := Post.New()
	Post.UUID = c.Param("id")

	result := h.DB.First(&Post)

	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)

		return
	}

	c.JSON(http.StatusOK, &Post)
}
