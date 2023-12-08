package api

import (
	db "Tech_Blog/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listPostsRequests struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=4,max=10"`
}

func (server *Server) listPosts(c *gin.Context) {
	var req listPostsRequests
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.ListPostsParams {
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	posts, err := server.store.ListPosts(c, arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, posts)

}
