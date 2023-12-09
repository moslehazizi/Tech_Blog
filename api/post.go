package api

import (
	db "Tech_Blog/db/sqlc"
	"database/sql"
	_ "encoding/json"
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
		return
	}

	arg := db.ListPostsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	posts, err := server.store.ListPosts(c, arg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	c.JSON(http.StatusOK, posts)

}

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPost(c *gin.Context) {
	var req getPostRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	post, err_1 := server.store.GetPost(c, req.ID)
	if err_1 != nil {
		if err_1 == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err_1))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err_1))
		return
	}

	arg := db.ListPostsParams{
		Limit:  10000,
		Offset: 0,
	}
	posts, err_2 := server.store.ListPosts(c, arg)
	if err_2 != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err_2))
	}


	arg_R := db.ListReviewsParams{
		Post: post.ID,
		Limit: 1000,
		Offset: 0,
	}
	reviews, err_3 := server.store.ListReviews(c, arg_R)
	if err_3 != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err_3))
	}

	c.JSON(http.StatusOK, post)
	c.JSON(http.StatusOK, posts[post.ID:post.ID+3])
	c.JSON(http.StatusOK, reviews)
}
