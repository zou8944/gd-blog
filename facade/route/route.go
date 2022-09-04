package route

import (
	"gd-blog/facade/controller"
	"gd-blog/ioc"
	"github.com/gin-gonic/gin"
)

func Init() (*gin.Engine, error) {

	var blogController controller.BlogController
	if object, err := ioc.Provide("blogController"); err != nil {
		return nil, err
	} else {
		blogController = object.(controller.BlogController)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/info", blogController.GetStat)
	r.GET("/tags", blogController.ListTags)
	r.GET("/categories", blogController.ListCategories)
	br := r.Group("/blogs")
	{
		br.GET("", blogController.ListBlog)
		br.POST("", blogController.CreateBlog)
		bdr := br.Group("/{id}")
		{
			bdr.GET("", blogController.GetBlog)
			bdr.PUT("", blogController.UpdateBlog)
			bdrl := bdr.Group("/like")
			{
				bdrl.POST("", blogController.LikeBlog)
				bdrl.DELETE("", blogController.UnLikeBlog)
			}
			bdrc := bdr.Group("/comments")
			{
				bdrc.GET("", blogController.ListComment)
				bdrc.POST("", blogController.CreateComment)
				bdrc.DELETE("", blogController.DeleteComment)
			}
		}
	}

	return r, nil
}
