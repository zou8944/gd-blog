package controller

import (
	"gd-blog/facade/dto"
	"gd-blog/facade/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentController struct {
	service service.CommentService
}

func NewCommentController(service service.CommentService) CommentController {
	return CommentController{service: service}
}

func (cc *CommentController) Create(c *gin.Context) {
	var commentDTO dto.Comment
	if err := c.ShouldBind(&commentDTO); err != nil {
		c.JSON(http.StatusBadRequest, dto.Reject(err))
	}
	cm, err := cc.service.Create(&commentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(err))
	}
	c.JSON(http.StatusOK, dto.Succeed(cm))
}

func (cc *CommentController) List(c *gin.Context) {
	blogId, err := strconv.ParseUint(c.Query("checkpoint"), 10, strconv.IntSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Reject(err))
	}
	cms, err := cc.service.List(uint(blogId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error(err))
	}
	c.JSON(http.StatusOK, dto.Succeed(cms))
}
