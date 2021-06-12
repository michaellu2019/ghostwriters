package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/michaellu2019/ghostwriters/models"
	"github.com/michaellu2019/ghostwriters/routes"
	"github.com/michaellu2019/ghostwriters/utils"
)

func main() {
	err := godotenv.Load("db_config.env")
	utils.ErrorCheck(err)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Printf("Running server on %s.\n", port)

	// initialize database
	models.InitDB()
	defer models.DB.Close()

	fs := http.FileServer(http.Dir("../client/build"))
	http.Handle("/", fs)

	// assign routes to functions
	http.HandleFunc("/api", routes.GetStories)

	http.HandleFunc("/api/get-stories", routes.GetStories)
	http.HandleFunc("/api/get-story", routes.GetStory)
	http.HandleFunc("/api/create-story", routes.CreateStory)
	http.HandleFunc("/api/delete-story", routes.DeleteStory)

	http.HandleFunc("/api/get-posts", routes.GetPosts)
	http.HandleFunc("/api/get-post", routes.GetPost)
	http.HandleFunc("/api/create-post", routes.CreatePost)
	http.HandleFunc("/api/delete-post", routes.DeletePost)
	http.HandleFunc("/api/like-post", routes.LikePost)
	http.HandleFunc("/api/unlike-post", routes.UnlikePost)

	http.HandleFunc("/api/get-all-post-likes", routes.GetAllPostLikes)
	http.HandleFunc("/api/get-post-likes", routes.GetPostLikes)
	http.HandleFunc("/api/get-author-post-likes", routes.GetAuthorPostLikes)

	// start server
	http.ListenAndServe(":"+port, nil)
}
