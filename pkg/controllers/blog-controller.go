package controllers

import (
	"gd-blog/pkg/models"
	"gd-blog/pkg/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetSiteInfo(c *gin.Context) {
	siteInfo, err := models.GetSiteInfo()
	if err != nil {
		utils.ResponseError(c, err)
	}
	utils.ResponseSuccess(c, siteInfo)
}

func ListBlog(c *gin.Context) {
	pageNo, err := strconv.ParseInt(c.DefaultQuery("pageNo", "1"), 0, 0)
	if err != nil {
		utils.ResponseReject(c, err)
	}
	pageSize, err := strconv.ParseInt(c.DefaultQuery("pageSize", "20"), 0, 0)
	if err != nil {
		utils.ResponseReject(c, err)
	}
	categoryId, err := strconv.ParseInt(c.DefaultQuery("cid", "0"), 0, 0)
	if err != nil {
		utils.ResponseReject(c, err)
	}
	blogs, err := models.ListBlog(categoryId, int(pageNo), int(pageSize))
	if err != nil {
		utils.ResponseError(c, err)
	}
	blogCount, err := models.CountBlog(categoryId)
	if err != nil {
		utils.ResponseError(c, err)
	}
	result := map[string]interface{}{
		"articles":     models.ModifyBlogSummary(blogs),
		"articleCount": blogCount,
	}
	utils.ResponseSuccess(c, result)
}

func GetBlog(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, strconv.IntSize)
	if err != nil {
		utils.ResponseReject(c, err)
	}
	blog, err := models.GetBlogById(id)
	if err != nil {
		utils.ResponseError(c, err)
	}
	utils.ResponseSuccess(c, blog)
}

func CreateBlog(c *gin.Context) {

}

func UpdateBlog(c *gin.Context) {

}

func DeleteBlog(c *gin.Context) {

}

func LikeBlog(c *gin.Context) {

}

func UnLikeBlog(c *gin.Context) {

}
