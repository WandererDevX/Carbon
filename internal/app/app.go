package server

import (
	"BlogSite/internal/handlers"
	"BlogSite/internal/storage/database"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	app := &App{
		Router: gin.Default(),
	}
	app.setupRouter()
	app.setupDatabase()

	return app
}

func (app *App) setupRouter() {
	app.Router.SetFuncMap(template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	})
	app.Router.LoadHTMLGlob("web/templates/*")

	app.Router.GET("/", handlers.HomeHandler)
	app.Router.GET("/add", handlers.AddPostPageHandler)
	app.Router.GET("/posts/:id", handlers.PostHandler)

	app.Router.POST("/addPost", handlers.AddPostHandler)

	app.Router.PUT("/posts/:id", handlers.UpdatePostHandler)

	app.Router.DELETE("/posts/:id", handlers.DeletePostHandler)

	app.Router.Static("/assets", "./internal/assets")
	app.Router.Static("/templates", "./web/templates")
}

func (app *App) setupDatabase() {
	db := database.GetDB()
	database.CreateTable(db)

}

func (app *App) Run(addr string) {
	if err := app.Router.Run(addr); err != nil {
		log.Fatal(err)
	}

}
