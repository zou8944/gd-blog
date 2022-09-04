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
	service service.BlogService
}

func NewBlogHandler(blogDomainService service.BlogService) BlogHandler {
	return BlogHandler{service: blogDomainService}
}

func (bh *BlogHandler) GetStat() func(c *gin.Context) {
	return func(c *gin.Context) {
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
}

func (bh *BlogHandler) ListCategories() func(c *gin.Context) {
	return func(c *gin.Context) {
		categories, err := bh.service.ListCategories()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, dto.Succeed(categories))
	}
}

func (bh *BlogHandler) ListTags() func(c *gin.Context) {
	return func(c *gin.Context) {
		tags, err := bh.service.ListTag()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, dto.Succeed(tags))
	}
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
