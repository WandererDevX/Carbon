package server

import (
	"Carbon/internal/handlers"
	"Carbon/internal/storage/database"
	"github.com/gin-gonic/gin"
	"html/template"
)

var router *gin.Engine

func NewApp() *gin.Engine {
	router = gin.Default()
	setupDatabase()
	setupRouter()

	return router
}

func setupRouter() {
	router.SetFuncMap(template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	})
	router.LoadHTMLGlob("web/templates/*")

	router.GET("/", handlers.HomeHandler)
	router.GET("/add", handlers.AddPostPageHandler)
	router.GET("/posts/:id", handlers.PostHandler)

	router.POST("/addPost", handlers.AddPostHandler)

	router.PUT("/posts/:id", handlers.UpdatePostHandler)

	router.DELETE("/posts/:id", handlers.DeletePostHandler)

	router.Static("/assets", "./internal/assets")
	router.Static("/templates", "./web/templates")
}

func setupDatabase() {
	database.CreateTable()
}

func Run() error {
	if err := router.Run(); err != nil {
		return err
	}
	return nil
}
