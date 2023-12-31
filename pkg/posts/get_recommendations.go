package Posts

import (
	"errors"
	"net/http"
	"redditClone/pkg/common/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h handler) GetRecommendations(c *gin.Context) {
	result := h.DB
	limit := 1

	if query, err := c.GetQuery("limit"); !err {
		c.AbortWithError(http.StatusBadRequest, errors.New("limit not set"))
		return
	} else if converted, err := strconv.Atoi(query); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("limit not valid"))
		return
	} else {
		limit = converted
		result.Limit(limit)
	}

	if query, err := c.GetQuery("filter"); !err {
		switch query {
		case "hot":
			result.Scopes(CreatedWithinPastDay, SortByUpVotes)
		case "votes":
			result.Order("vote_count")
		case "new":
			result.Where("? - ")
		default:
			c.AbortWithStatus(
				http.StatusBadRequest,
			)
			return
		}
	}

	var Posts []models.Post

	result.Limit(limit).Find(&Posts)

	if result.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": &Posts,
	})
	return
}
