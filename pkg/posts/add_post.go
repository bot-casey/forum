package Posts

import (
	"net/http"
	Post "redditClone/pkg/common/models/posts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddPostRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Flairs  string `json:"flairs"`
}

func (h handler) AddPost(c *gin.Context) {
	body := AddPostRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Post := Post.New()
	Post.Author = uuid.New().String()
	Post.Title = body.Title
	Post.Content = body.Content

	if result := h.DB.Create(&Post); result.Error != nil {
		c.AbortWithError(http.StatusUnprocessableEntity, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &Post)
	return
}
