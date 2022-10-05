package controllers

import (
	"gd-blog/pkg/models"
	"gd-blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBind(&comment); err != nil {
		utils.ResponseReject(c, err)
	}
	utils.ResponseSuccess(c, nil)
}

func ListComment(c *gin.Context) {

}
