package routers

import (
	v2 "github.com/EDDYCJY/go-gin-example/routers/api/v2"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	"github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	// Init API V2
	apiV2 := r.Group("api/v2")

	// Using API V2
	apiV2.Use()
	{
		apiV2.POST("/user", v2.GetUsers)
		apiV2.GET("/user", v2.GetUsers)
	}

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		//
		apiV1.GET("/tags", v1.GetTags)
		//
		apiV1.POST("/tags", v1.AddTag)
		//
		apiV1.PUT("/tags/:id", v1.EditTag)
		//
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		//
		r.POST("/tags/export", v1.ExportTag)
		//
		r.POST("/tags/import", v1.ImportTag)

		//
		apiV1.GET("/articles", v1.GetArticles)
		//
		apiV1.GET("/articles/:id", v1.GetArticle)
		//
		apiV1.POST("/articles", v1.AddArticle)
		//
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
		//
		apiV1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
