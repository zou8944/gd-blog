package route

import (
	"gd-blog/src/facade/dto"
	"gd-blog/src/facade/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListQueryParam struct {
	SeparateId int `json:"separateIdd"`
	Limit      int `json:"limit"`
}

type BlogHandler struct {
	service service.BlogDomainService
}

func NewBlogHandler(blogDomainService service.BlogDomainService) BlogHandler {
	return BlogHandler{service: blogDomainService}
}

func (bh *BlogHandler) ListBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
		var p ListQueryParam
		if err := c.ShouldBindQuery(&p); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		if p.Limit <= 0 {
			p.Limit = 20
		}
		blogs, err := bh.service.ListBlog(p.SeparateId, p.Limit)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, dto.Succeed(blogs))
	}
}

func (bh *BlogHandler) CreateBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) GetBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) UpdateBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) LikeBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) UnLikeBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) CreateComment() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) ListComment() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bh *BlogHandler) DeleteComment() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}
