package Votes

import (
	"errors"
	"net/http"

	Post "redditClone/pkg/common/models/posts"
	Vote "redditClone/pkg/common/models/votes"

	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Resource      string                  `json:"resource"`
	Resource_Type Vote.Enum_resource_type `json:"resource_type"`
	User          string                  `json:""`
}

func (h handler) getVoteByUser(c *gin.Context) {
	body := requestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	resource_type, err := Vote.Enum_resource_type.String(body.Resource_Type)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid resource_type"))
		return
	}

	var Vote Vote.Vote_Relations

	Vote.Resource = body.Resource
	Vote.User = body.User
	Vote.Resource_Type = resource_type

	var Post Post.Post
	Post.UUID = Vote.Resource

	if result := h.DB.First(&Post); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	Post.IncrementVote()

	if result := h.DB.Save(&Post); result.Error != nil {
		c.AbortWithError(http.StatusConflict, result.Error)
		return
	}

	if result := h.DB.Create(&Vote); result.Error != nil {
		c.Status(http.StatusAlreadyReported)
		return
	}

	c.JSON(http.StatusCreated, &Vote)
	return
}