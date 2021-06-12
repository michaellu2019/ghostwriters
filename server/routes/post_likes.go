package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaellu2019/ghostwriters/models"
	"github.com/michaellu2019/ghostwriters/utils"
)

// structs for JSON objects
type PostLike struct {
	Id        int       `json:"id"`
	PostId    int       `json:"post_id"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

type PostLikesListPayload struct {
	Status string `json:"status"`
	Data   struct {
		PostLikes []PostLike `json:"post_likes"`
	} `json:"data"`
}

func GetAllPostLikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		// get all rows in the post_likes table
		rows, err := models.DB.Query("SELECT id, post_id, author, created_at FROM post_likes")
		utils.ErrorCheck(err)

		var postLike PostLike
		payload := &PostLikesListPayload{
			Status: "OK",
		}

		// add all row column data to the payload data struct
		for rows.Next() {
			err = rows.Scan(&postLike.Id, &postLike.PostId, &postLike.Author, &postLike.CreatedAt)
			utils.ErrorCheck(err)

			payload.Data.PostLikes = append(payload.Data.PostLikes, postLike)
		}

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only GET requests allowed.",
		})
		return
	}
}

func GetPostLikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		// parse the URL for the provided post ID
		params, ok := r.URL.Query()["id"]
		if !ok || len(params[0]) < 1 {
			fmt.Fprintf(w, "Missing URL parameter.")
			return
		}

		postId := params[0]

		// get all rows in the post_likes table with the specified post ID
		rows, err := models.DB.Query("SELECT id, post_id, author, created_at FROM post_likes WHERE post_id=?", postId)
		utils.ErrorCheck(err)

		var postLike PostLike
		payload := &PostLikesListPayload{
			Status: "OK",
		}

		// add all row column data to the payload data struct
		for rows.Next() {
			err = rows.Scan(&postLike.Id, &postLike.PostId, &postLike.Author, &postLike.CreatedAt)
			utils.ErrorCheck(err)

			payload.Data.PostLikes = append(payload.Data.PostLikes, postLike)
		}

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only GET requests allowed.",
		})
		return
	}
}

func GetAuthorPostLikes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		// parse the URL for the provided post ID
		params, ok := r.URL.Query()["author"]
		if !ok || len(params[0]) < 1 {
			fmt.Fprintf(w, "Missing URL parameter.")
			return
		}

		author := params[0]

		// get all rows in the post_likes table with the specified author
		rows, err := models.DB.Query("SELECT id, post_id, author, created_at FROM post_likes WHERE author=?", author)
		utils.ErrorCheck(err)

		var postLike PostLike
		payload := &PostLikesListPayload{
			Status: "OK",
		}

		// add all row column data to the payload data struct
		for rows.Next() {
			err = rows.Scan(&postLike.Id, &postLike.PostId, &postLike.Author, &postLike.CreatedAt)
			utils.ErrorCheck(err)

			payload.Data.PostLikes = append(payload.Data.PostLikes, postLike)
		}

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only GET requests allowed.",
		})
		return
	}
}
