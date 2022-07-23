package main;

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	
	booksAPI := app.Party("/books")

	booksAPI.Get("/", list)

	app.Listen(":8080")
}

func list(ctx iris.Context) {
	ctx.JSON("List of books")
}