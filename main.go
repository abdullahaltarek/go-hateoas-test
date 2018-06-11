package main

import (
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"github.com/kataras/iris/middleware/logger"
)

var books *mgo.Collection

func main() {
	app := iris.New()
	//app.Logger().SetLevel("info")
	app.Use(logger.New())

	conn, err := mgo.Dial("127.0.0.1:27017")

	if err != nil {
		panic(err.Error())
	}

	books = conn.DB("bookAPI").C("book")

	booksRoute := app.Party("api/books")
	booksRoute.Get("/", GetAllBooks)
	booksRoute.Get("/:bookId", GetSingleBook)
	booksRoute.Post("/", CreateBook)
	booksRoute.Put("/:bookId", UpdateBook)
	booksRoute.Delete("/:bookId", DeleteBook)


	app.Run(iris.Addr(":8021"))
}
