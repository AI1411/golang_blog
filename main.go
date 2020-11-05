package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-blog/handler"
	"go-blog/repository"
	"log"
)

const tmplPath = "src/template/"

var db *sqlx.DB
var e = createMux()

func main() {
	db = connectDB()
	repository.SetDb(db)
	e.GET("/", handler.ArticleIndex)
	e.GET("/new", handler.ArticleNew)
	e.GET("/:id", handler.ArticleShow)
	e.GET("/:id/edit", handler.ArticleEdit)

	e.Logger.Fatal(e.Start(":8080"))
}

func connectDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/techblog")
	if err != nil {
		e.Logger.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")
	return db
}

func createMux() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("css", "src/css")
	e.Static("js", "src/js")

	return e
}
