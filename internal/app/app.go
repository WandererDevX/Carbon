package server

import (
	"Carbon/internal/handlers"
	"Carbon/internal/storage/database"
	"database/sql"
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
	db := app.setupDatabase()
	app.setupRouter(db)

	return app
}

func (app *App) setupRouter(db *sql.DB) {
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

func (app *App) setupDatabase() *sql.DB {
	db := database.GetDB()
	database.CreateTable(db)
	return db

}

func (app *App) Run() {
	addr := ":8080"
	if err := app.Router.Run(addr); err != nil {
		log.Fatal(err)
	}

}
