package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaellu2019/ghostwriters/models"
	"github.com/michaellu2019/ghostwriters/utils"
)

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
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		rows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts")
		utils.QueryErrorCheck(err)

		var post Post
		payload := &PostListPayload{
			Status: "OK",
		}

		for rows.Next() {
			err = rows.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.QueryErrorCheck(err)

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

		params, ok := r.URL.Query()["id"]
		if !ok || len(params[0]) < 1 {
			fmt.Fprintf(w, "Missing URL parameter.")
			return
		}

		postId := params[0]

		row := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", postId)

		var post Post
		payload := &PostListPayload{
			Status: "OK",
		}

		err := row.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
		utils.QueryErrorCheck(err)

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
		var post Post
		var s Story

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		if len(p.Text) > 0 && len(p.Author) > 0 {
			selectedRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE id=?", p.StoryId)

			err = selectedRow.Scan(&s.Id, &s.Author, &s.Title, &s.CreatedAt)
			utils.QueryErrorCheck(err)

			row, err := models.DB.Query("INSERT INTO posts(story_id, author, text) VALUES(?, ?, ?)", p.StoryId, p.Author, p.Text)
			utils.QueryErrorCheck(err)

			defer row.Close()

			insertedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE story_id=? AND author=? AND text=?", p.StoryId, p.Author, p.Text)
			err = insertedRow.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.QueryErrorCheck(err)

			payload := PostListPayload{
				Status: "OK",
			}
			payload.Data.Posts = append(payload.Data.Posts, post)

			utils.SendJSON(w, payload)
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

		insertedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", p.Id)
		err = insertedRow.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
		utils.QueryErrorCheck(err)

		row, err := models.DB.Query("DELETE FROM posts WHERE id=?", p.Id)
		utils.QueryErrorCheck(err)

		defer row.Close()

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

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		selectedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", p.Id)

		err = selectedRow.Scan(&oldPost.Id, &oldPost.StoryId, &oldPost.Author, &oldPost.Text, &oldPost.Likes, &oldPost.Dislikes, &oldPost.CreatedAt)
		utils.QueryErrorCheck(err)

		oldPost.Likes += 1
		updatedRow, err := models.DB.Query("UPDATE posts SET likes=? WHERE id=?", oldPost.Likes, p.Id)
		utils.QueryErrorCheck(err)

		defer updatedRow.Close()

		payload := PostListPayload{
			Status: "OK",
		}
		payload.Data.Posts = append(payload.Data.Posts, oldPost)

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var p Post
		var oldPost Post

		err := utils.DecodeJSONBody(w, r, &p)
		utils.JSONErrorCheck(w, err)

		selectedRow := models.DB.QueryRow("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE id=?", p.Id)

		err = selectedRow.Scan(&oldPost.Id, &oldPost.StoryId, &oldPost.Author, &oldPost.Text, &oldPost.Likes, &oldPost.Dislikes, &oldPost.CreatedAt)
		utils.QueryErrorCheck(err)

		oldPost.Dislikes += 1
		updatedRow, err := models.DB.Query("UPDATE posts SET dislikes=? WHERE id=?", oldPost.Dislikes, p.Id)
		utils.QueryErrorCheck(err)

		defer updatedRow.Close()

		payload := PostListPayload{
			Status: "OK",
		}
		payload.Data.Posts = append(payload.Data.Posts, oldPost)

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}
