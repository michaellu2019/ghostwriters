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

	// assign routes to functions
	http.HandleFunc("/", routes.GetStories)

	http.HandleFunc("/get-stories", routes.GetStories)
	http.HandleFunc("/get-story", routes.GetStory)
	http.HandleFunc("/create-story", routes.CreateStory)
	http.HandleFunc("/delete-story", routes.DeleteStory)

	http.HandleFunc("/get-posts", routes.GetPosts)
	http.HandleFunc("/get-post", routes.GetPost)
	http.HandleFunc("/create-post", routes.CreatePost)
	http.HandleFunc("/delete-post", routes.DeletePost)
	http.HandleFunc("/like-post", routes.LikePost)
	http.HandleFunc("/dislike-post", routes.DislikePost)

	// start server
	http.ListenAndServe("127.0.0.1:"+port, nil)
}
