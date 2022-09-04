package route

import (
	"gd-blog/src/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListQueryParam struct {
	SeparateId int `json:"separateIdd"`
	Limit      int `json:"limit"`
}

type BlogHandler struct {
	blogDomainService service.BlogDomainService
}

func NewBlogHandler(blogDomainService service.BlogDomainService) BlogHandler {
	return BlogHandler{blogDomainService: blogDomainService}
}

func (bs *BlogHandler) HandleListBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
		var p ListQueryParam
		if err := c.ShouldBindQuery(&p); err != nil {
			c.JSON(http.StatusBadRequest, err)
		}
		blogs, err := bs.blogDomainService.ListBlog(p.SeparateId, p.Limit)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, blogs)
	}
}

func (bs *BlogHandler) HandleCreateBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleGetBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleUpdateBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleLikeBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleUnLikeBlog() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleCreateComment() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleListComment() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}

func (bs *BlogHandler) HandleDeleteComment() func(c *gin.Context) {
	return func(c *gin.Context) {
	}
}
