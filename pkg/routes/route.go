package routes

import (
	"gd-blog/pkg/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Init() (*gin.Engine, error) {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
	}))

	r.GET("/info", controllers.GetSiteInfo)
	r.GET("/tags", controllers.ListTag)
	r.GET("/categories", controllers.ListCategory)
	r.GET("/blogs", controllers.ListBlog)
	r.POST("/blogs", controllers.CreateBlog)
	r.GET("/blogs/:id", controllers.GetBlog)
	r.PUT("/blogs/:id", controllers.UpdateBlog)
	r.POST("/blogs/:id/like", controllers.LikeBlog)
	r.DELETE("/blogs/:id/like", controllers.UnLikeBlog)
	r.GET("/blogs/:id/comments", controllers.ListComment)
	r.POST("/blogs/:id/comments", controllers.CreateComment)

	return r, nil
}

func RegisterRoutes(r *gin.Engine) {

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
	}))

	r.GET("/info", controllers.GetSiteInfo)
	r.GET("/tags", controllers.ListTag)
	r.GET("/categories", controllers.ListCategory)
	r.GET("/blogs", controllers.ListBlog)
	r.POST("/blogs", controllers.CreateBlog)
	r.GET("/blogs/:id", controllers.GetBlog)
	r.PUT("/blogs/:id", controllers.UpdateBlog)
	r.POST("/blogs/:id/like", controllers.LikeBlog)
	r.DELETE("/blogs/:id/like", controllers.UnLikeBlog)
	r.GET("/blogs/:id/comments", controllers.ListComment)
	r.POST("/blogs/:id/comments", controllers.CreateComment)
}
