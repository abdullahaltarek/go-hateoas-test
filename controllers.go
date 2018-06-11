package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func GetAllBooks(c context.Context) {
	var booksList []Book
	query := map[string]string{}
	if c.URLParams()["genre"] != "" {
		query["genre"] = c.URLParams()["genre"]
	}
	err := books.Find(query).All(&booksList)
	if err != nil {
		c.StatusCode(500)
		return
	}

	var returnBooks []Book

	for _, j := range booksList {
		self := map[string]string{"FilterByThisGenre": "http://" + c.Host() + "/api/books/" + j.ID.Hex()}
		j.Links = append(j.Links, self)

		returnBooks = append(returnBooks, j)
	}
	c.StatusCode(200)
	c.JSON(returnBooks)
}

func GetSingleBook(c context.Context) {
	var book Book
	idCheck := bson.IsObjectIdHex(c.Params().Get("bookId"))

	if idCheck {
		err := books.FindId(bson.ObjectIdHex(c.Params().Get("bookId"))).One(&book)
		if err != nil {
			c.StatusCode(404)
			c.JSON(iris.Map{"msg": "sorry, not found"})
			return
		} else {
			self := map[string]string{"self": "http://" + c.Host() + "/api/books/?genre=" + strings.Replace(book.Genre, " ", "+", -1)}
			book.Links = append(book.Links, self)

			c.StatusCode(200)
			c.JSON(book)
		}

	} else {
		c.StatusCode(404)
		c.JSON(iris.Map{"msg": "sorry, not found"})
	}
}

func CreateBook(c context.Context) {
	var book Book
	err := c.ReadJSON(&book)

	if err != nil {
		c.StatusCode(500)
		return
	}

	book.ID = bson.NewObjectId()

	err = books.Insert(book)
	if err != nil {
		c.StatusCode(500)
		return
	}

	c.StatusCode(201)
	c.JSON(book)
}

func UpdateBook(c context.Context) {
	var book Book

	idCheck := bson.IsObjectIdHex(c.Params().Get("bookId"))

	if idCheck {
		err := books.FindId(bson.ObjectIdHex(c.Params().Get("bookId"))).One(&book)
		if err != nil {
			c.StatusCode(500)
			return
		} else {
			c.ReadJSON(&book)
			err = books.UpdateId(book.ID, book)

			if err != nil {
				c.StatusCode(500)
				return
			} else {
				c.StatusCode(200)
				c.JSON(book)
			}
		}
	} else {
		c.StatusCode(404)
		c.JSON(iris.Map{"msg": "sorry, not found"})
	}
}

func DeleteBook(c context.Context) {
	idCheck := bson.IsObjectIdHex(c.Params().Get("bookId"))

	if idCheck {
		err := books.RemoveId(bson.ObjectIdHex(c.Params().Get("bookId")))

		if err != nil {
			c.StatusCode(404)
			c.JSON(iris.Map{"msg": "sorry, not found"})
			return
		} else {
			c.StatusCode(200)
			c.JSON(iris.Map{"msg": "removed"})
		}
	} else {
		c.StatusCode(404)
		c.JSON(iris.Map{"msg": "sorry, not found"})
	}
}
