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

	br := r.Group("/blogs")
	{
		br.GET("", handler.HandleListBlog())
		br.POST("", handler.HandleCreateBlog())
		bdr := br.Group("/{id}")
		{
			bdr.GET("", handler.HandleGetBlog())
			bdr.PUT("", handler.HandleUpdateBlog())
			bdrl := bdr.Group("/like")
			{
				bdrl.POST("", handler.HandleLikeBlog())
				bdrl.DELETE("", handler.HandleUnLikeBlog())
			}
			bdrc := bdr.Group("/comments")
			{
				bdrc.GET("", handler.HandleListComment())
				bdrc.POST("", handler.HandleCreateComment())
				bdrc.DELETE("", handler.HandleDeleteComment())
			}
		}
	}

	return r, nil
}
