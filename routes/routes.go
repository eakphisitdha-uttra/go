package routes

import (
	"database/sql"
	"os"
	"strings"

	_ "microservice/docs"
	"microservice/internals/a/handlers"
	"microservice/internals/a/repositories"
	"microservice/internals/a/usecases"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(pg *sql.DB, mg *mongo.Client) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(contentType())
	r.Use(corsMiddleware())

	// Static file serving
	r.StaticFS("/storages/excels", gin.Dir(os.Getenv("STORAGES_FOLDER_EXCEL_PATH"), true))
	r.StaticFS("/storages/pdfs", gin.Dir(os.Getenv("STORAGES_FOLDER_PDF_PATH"), true))
	r.StaticFS("/storages/images", gin.Dir(os.Getenv("STORAGES_FOLDER_IMAGE_PATH"), true))
	r.StaticFS("/storages/htmls", gin.Dir(os.Getenv("STORAGES_FOLDER_HTML_PATH"), true))

	// Swagger documentation (only in non-production environments)
	if os.Getenv("ENV") != "production" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// API path
	path := r.Group("/api/v1")

	//module A
	a := path.Group("/a")
	{
		aRepositories := repositories.NewRepository(pg, mg)
		aUsecases := usecases.NewUsecase(aRepositories)
		aHandler := handlers.NewHandler(aUsecases)

		a.GET("/get", aHandler.Get)
		a.POST("/add", aHandler.Add)
		a.PUT("/update/:id", aHandler.Update)
		a.DELETE("/delete/:id", aHandler.Delete)
	}

	return r
}

func contentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // process request first, then modify headers
		switch {
		case strings.HasSuffix(c.Request.URL.Path, ".png"):
			c.Header("Content-Type", "image/png")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		case strings.HasSuffix(c.Request.URL.Path, ".jpg") || strings.HasSuffix(c.Request.URL.Path, ".jpeg"):
			c.Header("Content-Type", "image/jpeg")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		case strings.HasSuffix(c.Request.URL.Path, ".xlsx"):
			c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		case strings.HasSuffix(c.Request.URL.Path, ".xls"):
			c.Header("Content-Type", "application/vnd.ms-excel")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
