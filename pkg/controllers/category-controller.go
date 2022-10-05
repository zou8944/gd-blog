package controllers

import (
	"gd-blog/pkg/models"
	"gd-blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	categories, err := models.AllCategory()
	if err != nil {
		utils.ResponseError(c, err)
	}
	utils.ResponseSuccess(c, categories)
}
