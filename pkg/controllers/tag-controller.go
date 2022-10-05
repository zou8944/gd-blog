package controllers

import (
	"gd-blog/pkg/models"
	"gd-blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ListTag(c *gin.Context) {
	tags, err := models.AllTag()
	if err != nil {
		utils.ResponseError(c, err)
	}
	utils.ResponseSuccess(c, tags)
}
