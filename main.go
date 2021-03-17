package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Post is ...
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func getPosts(c echo.Context) error {
	listAr := []Post{
		{
			ID:    "1",
			Title: "First post",
			Text:  "Small content.",
		},
		{
			ID:    "2",
			Title: "Long post",
			Text:  "The second post present a long content. It can be to easy test forms on small devices.",
		},
		{
			ID:    "3",
			Title: "Third post",
			Text:  "Any content",
		},
	}

	return c.JSON(http.StatusOK, listAr)
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

//File is ...
func uploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "File uploaded")
}

func downloadFile(c echo.Context) error {
	file := c.Param("file")
	return c.Attachment(file, file)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route Posts
	e.GET("/posts", getPosts)
	e.GET("/posts/:id", getPost)

	// Route Users
	e.GET("/users/:id", getUser)

	// Route Files
	e.POST("/upload", uploadFile)
	e.GET("/download/:file", downloadFile)

	e.Logger.Fatal(e.Start(":80"))
}
