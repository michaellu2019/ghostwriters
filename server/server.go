package main

import (
	"fmt"
	"net/http"

	"github.com/michaellu2019/ghostwriters/models"
	"github.com/michaellu2019/ghostwriters/routes"
)

func main() {
	const port string = "8000"
	fmt.Printf("Running server on %s.\n", port)

	// initialize database
	models.InitDB()
	defer models.DB.Close()

	fs := http.FileServer(http.Dir("../client/build"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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
	http.HandleFunc("/api/dislike-post", routes.DislikePost)

	// start server
	http.ListenAndServe("127.0.0.1:"+port, nil)
}
