package route

import (
	"gd-blog/src/ioc"
	"github.com/gin-gonic/gin"
)

func InitRoutes() (*gin.Engine, error) {

	blogHandler, err := ioc.Provide("blogHandler")
	if err != nil {
		return nil, err
	}
	handler := blogHandler.(BlogHandler)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/info", handler.GetStat())
	br := r.Group("/blogs")
	{
		br.GET("", handler.ListBlog())
		br.POST("", handler.CreateBlog())
		bdr := br.Group("/{id}")
		{
			bdr.GET("", handler.GetBlog())
			bdr.PUT("", handler.UpdateBlog())
			bdrl := bdr.Group("/like")
			{
				bdrl.POST("", handler.LikeBlog())
				bdrl.DELETE("", handler.UnLikeBlog())
			}
			bdrc := bdr.Group("/comments")
			{
				bdrc.GET("", handler.ListComment())
				bdrc.POST("", handler.CreateComment())
				bdrc.DELETE("", handler.DeleteComment())
			}
		}
	}

	return r, nil
}
