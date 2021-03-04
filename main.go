package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//Post is ...
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func getPost(c echo.Context) error {
	id := c.Param("id")

	var post = Post{
		ID:    id,
		Title: "Title",
		Text:  "Text",
	}

	return c.JSON(http.StatusOK, post)
}

//User is ...
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	var user = User{
		ID:   id,
		Name: "Name",
	}

	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/posts/:id", getPost)
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":80"))
}
