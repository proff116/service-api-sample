package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//Post is ...
type Post struct {
	ID      string    `json:"id"`
	Time    time.Time `json:"time"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

func getPosts(c echo.Context) error {
	listAr := []Post{
		{
			ID:      "1",
			Time:    time.Now(),
			Title:   "First post",
			Content: "Small content.",
		},
		{
			ID:      "2",
			Time:    time.Now(),
			Title:   "Long post",
			Content: "The second post present a long content. It can be to easy test forms on small devices.",
		},
		{
			ID:      "3",
			Time:    time.Now(),
			Title:   "Third post",
			Content: "Any content",
		},
	}

	return c.JSON(http.StatusOK, listAr)
}

func getPost(c echo.Context) error {
	id := c.Param("id")

	var post = Post{
		ID:      id,
		Time:    time.Now(),
		Title:   "Title",
		Content: "Content",
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
