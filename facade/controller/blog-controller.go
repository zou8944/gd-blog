package controller

import (
	"gd-blog/facade/dto"
	"gd-blog/facade/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BlogController struct {
	service service.BlogService
}

func NewBlogController(blogDomainService service.BlogService) BlogController {
	return BlogController{service: blogDomainService}
}

func (bh *BlogController) GetStat(c *gin.Context) {
	stat, err := bh.service.GetSiteStat()
	if err != nil {
		panic(err)
	}
	siteInfo := dto.SiteInfo{
		Author: dto.Author{
			Name:   "果冻",
			Desc:   "果冻的碎碎念",
			Avatar: "https://thirdwx.qlogo.cn/mmopen/vi_32/DYAIOgq83equib0YGKeGrRww67LyZ7hSONtAW59RHDTd2JuKmSfQLEs8zWIB14hUcHibNG41zNibv5mr5QhM5QDMQ/132",
			CSDN:   "https://blog.csdn.net/zou8944",
			Github: "https://github.com/zou8944",
		},
		Statistics: *stat,
		Beian:      "粤ICP备2021024139号",
	}
	c.JSON(http.StatusOK, dto.Succeed(siteInfo))
}

func (bh *BlogController) ListCategory(c *gin.Context) {
	categories, err := bh.service.ListCategories()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, dto.Succeed(categories))
}

func (bh *BlogController) ListTag(c *gin.Context) {
	tags, err := bh.service.ListTag()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, dto.Succeed(tags))
}

func (bh *BlogController) ListBlog(c *gin.Context) {
	checkpoint, err := strconv.ParseInt(c.Query("checkpoint"), 10, strconv.IntSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Reject(err))
	}
	limit, err := strconv.ParseInt(c.DefaultQuery("limit", "20"), 10, strconv.IntSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Reject(err))
	}
	blogs, err := bh.service.ListBlog(int(checkpoint), int(limit))
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, dto.Succeed(blogs))
}

func (bh *BlogController) CreateBlog(c *gin.Context) {
}

func (bh *BlogController) GetBlog(c *gin.Context) {
}

func (bh *BlogController) UpdateBlog(c *gin.Context) {
}

func (bh *BlogController) LikeBlog(c *gin.Context) {

}

func (bh *BlogController) UnLikeBlog(c *gin.Context) {
}
