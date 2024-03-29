package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/michaellu2019/ghostwriters/models"
	"github.com/michaellu2019/ghostwriters/utils"
)

// structs for JSON objects
type Post struct {
	Id        int       `json:"id"`
	StoryId   int       `json:"story_id"`
	Author    string    `json:"author"`
	Text      string    `json:"text"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"created_at"`
}

type PostListPayload struct {
	Status string `json:"status"`
	Data   struct {
		Posts []Post `json:"posts"`
	} `json:"data"`
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		// get all rows in the posts table
		rows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts")
		utils.ErrorCheck(err)

		var post Post
		payload := &PostListPayload{
			Status: "OK",
		}

		// add all row column data to the payload data struct
		for rows.Next() {
			err = rows.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.ErrorCheck(err)

			payload.Data.Posts = append(payload.Data.Posts, post)
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

func GetPost(w http.ResponseWriter, r *http.Request) {
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

		// get the row in the posts table with the specified ID
		row := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", postId)

		var post Post
		payload := &PostListPayload{
			Status: "OK",
		}

		// add the row column data to the payload data struct
		err := row.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
		utils.ErrorCheck(err)

		payload.Data.Posts = append(payload.Data.Posts, post)

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only GET requests allowed.",
		})
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var p Post
		var oldPost Post
		var post Post
		var s Story

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		// make sure no empty data was sent
		if len(p.Text) > 0 && len(p.Author) > 0 {
			// make sure character limit was not exceeded
			if len(p.Text) < 1600 {
				// get the row in the stories table that has the specified story ID to check that the post can be assigned to a story
				selectedRow := models.DB.QueryRow("SELECT id, author, title, image_url, created_at FROM stories WHERE id=?", p.StoryId)

				err = selectedRow.Scan(&s.Id, &s.Author, &s.Title, &s.ImageURL, &s.CreatedAt)
				utils.ErrorCheck(err)

				postRows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE story_id=? ORDER BY created_at DESC", p.StoryId)
				utils.ErrorCheck(err)

				if postRows.Next() {
					err = postRows.Scan(&oldPost.Id, &oldPost.StoryId, &oldPost.Author, &oldPost.Text, &oldPost.Likes, &oldPost.Dislikes, &oldPost.CreatedAt)
					utils.ErrorCheck(err)

					if oldPost.Author == p.Author {
						utils.SendJSON(w, ErrorPayload{
							Status: "ERROR",
							Data:   "Posting two times in a row not permitted. Please wait for another user to post on this story before posting again.",
						})
						return
					}
				}

				// insert a new row with the specified values in the JSON object into the posts table
				row, err := models.DB.Query("INSERT INTO posts(story_id, author, text) VALUES(?, ?, ?)", p.StoryId, p.Author, p.Text)
				utils.ErrorCheck(err)

				defer row.Close()

				// get the row column data that was just inserted and add it to the payload
				insertedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE story_id=? AND author=? AND text=?", p.StoryId, p.Author, p.Text)
				err = insertedRow.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
				utils.ErrorCheck(err)

				payload := PostListPayload{
					Status: "OK",
				}
				payload.Data.Posts = append(payload.Data.Posts, post)

				utils.SendJSON(w, payload)
			} else {
				utils.SendJSON(w, ErrorPayload{
					Status: "ERROR",
					Data:   "Too many characters.",
				})
			}
		} else {
			utils.SendJSON(w, ErrorPayload{
				Status: "ERROR",
				Data:   "Empty values.",
			})
		}
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var p Post
		var post Post

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		// get the row in the posts table that has the specified post ID so that its data can be returned when deleted
		insertedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", p.Id)
		err = insertedRow.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
		utils.ErrorCheck(err)

		// delete the row with the specified post ID from the posts table
		row, err := models.DB.Query("DELETE FROM posts WHERE id=?", p.Id)
		utils.ErrorCheck(err)

		defer row.Close()

		// return the data of the deleted post
		payload := PostListPayload{
			Status: "OK",
		}
		payload.Data.Posts = append(payload.Data.Posts, post)

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var p Post
		var oldPost Post
		var postLike PostLike

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		// get the row in the posts table with the specified ID
		selectedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", p.Id)

		err = selectedRow.Scan(&oldPost.Id, &oldPost.StoryId, &oldPost.Author, &oldPost.Text, &oldPost.Likes, &oldPost.Dislikes, &oldPost.CreatedAt)
		utils.ErrorCheck(err)

		// check that the author has not already liked the post
		existingRow := models.DB.QueryRow("SELECT id, post_id, author, created_at FROM post_likes WHERE post_id=? AND author=?", p.Id, p.Author)
		err = existingRow.Scan(&postLike.Id, &postLike.PostId, &postLike.Author, &postLike.CreatedAt)

		if err != nil && err == sql.ErrNoRows {
			// update the number of likes for the post
			oldPost.Likes += 1
			updatedRow, err := models.DB.Query("UPDATE posts SET likes=? WHERE id=?", oldPost.Likes, p.Id)
			utils.ErrorCheck(err)

			defer updatedRow.Close()

			// insert a new row with the specified values in the JSON object into the post_likes table
			// author in this case is not the author of the post but another author/user who is liking the author's post
			row, err := models.DB.Query("INSERT INTO post_likes(post_id, author) VALUES(?, ?)", oldPost.Id, p.Author)
			utils.ErrorCheck(err)

			defer row.Close()

			// return the data of the updated post
			payload := PostListPayload{
				Status: "OK",
			}
			payload.Data.Posts = append(payload.Data.Posts, oldPost)

			utils.SendJSON(w, payload)
		} else {
			utils.SendJSON(w, ErrorPayload{
				Status: "ERROR",
				Data:   "Author already liked post.",
			})
		}
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}

func UnlikePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var p Post
		var oldPost Post
		var postLike PostLike

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		// get the row in the posts table with the specified ID
		selectedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", p.Id)

		err = selectedRow.Scan(&oldPost.Id, &oldPost.StoryId, &oldPost.Author, &oldPost.Text, &oldPost.Likes, &oldPost.Dislikes, &oldPost.CreatedAt)
		utils.ErrorCheck(err)

		// check that the author has already liked the post
		existingRow := models.DB.QueryRow("SELECT id, post_id, author, created_at FROM post_likes WHERE post_id=? AND author=?", p.Id, p.Author)
		err = existingRow.Scan(&postLike.Id, &postLike.PostId, &postLike.Author, &postLike.CreatedAt)

		if err != nil && err == sql.ErrNoRows {
			utils.SendJSON(w, ErrorPayload{
				Status: "ERROR",
				Data:   "Author never liked post.",
			})
		} else {
			// update the number of likes for the post
			oldPost.Likes -= 1
			updatedRow, err := models.DB.Query("UPDATE posts SET likes=? WHERE id=?", oldPost.Likes, p.Id)
			utils.ErrorCheck(err)

			defer updatedRow.Close()

			// delete the row with the specified values in the JSON object from the post_likes table
			// author in this case is not the author of the post but another author/user who is liking the author's post
			row, err := models.DB.Query("DELETE FROM post_likes WHERE post_id=? AND author=?", oldPost.Id, p.Author)
			utils.ErrorCheck(err)

			defer row.Close()

			// return the data of the updated post
			payload := PostListPayload{
				Status: "OK",
			}
			payload.Data.Posts = append(payload.Data.Posts, oldPost)

			utils.SendJSON(w, payload)
		}
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}
