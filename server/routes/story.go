package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaellu2019/democracy/models"
	"github.com/michaellu2019/democracy/utils"
)

type Story struct {
	Id        int       `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type StoryPayload struct {
	Id        int       `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   []Post    `json:"content"`
}

type StoryListPayload struct {
	Status string `json:"status"`
	Data   struct {
		Stories []StoryPayload `json:"stories"`
	} `json:"data"`
}

type ErrorPayload struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func GetStories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		storyRows, err := models.DB.Query("SELECT id, author, title, created_at FROM stories")
		utils.QueryErrorCheck(err)

		var story Story
		storyMap := make(map[int]*StoryPayload)

		for storyRows.Next() {
			err = storyRows.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
			utils.QueryErrorCheck(err)

			storyMap[story.Id] = &StoryPayload{
				Id:        story.Id,
				Author:    story.Author,
				Title:     story.Title,
				CreatedAt: story.CreatedAt,
			}
		}

		postRows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts ORDER BY created_at ASC")
		utils.QueryErrorCheck(err)

		var post Post

		for postRows.Next() {
			err = postRows.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.QueryErrorCheck(err)
			storyMap[post.StoryId].Content = append(storyMap[post.StoryId].Content, post)
		}

		payload := &StoryListPayload{
			Status: "OK",
		}
		for _, v := range storyMap {
			payload.Data.Stories = append(payload.Data.Stories, *v)
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

func GetStoryTitles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		rows, err := models.DB.Query("SELECT id, author, title, created_at FROM stories")
		utils.QueryErrorCheck(err)

		var story Story
		payload := &StoryListPayload{
			Status: "OK",
		}

		for rows.Next() {
			err = rows.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
			utils.QueryErrorCheck(err)

			storyPayload := StoryPayload{
				Id:        story.Id,
				Author:    story.Author,
				Title:     story.Title,
				CreatedAt: story.CreatedAt,
			}
			payload.Data.Stories = append(payload.Data.Stories, storyPayload)
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

func GetStory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		defer utils.RecoverErrorCheck(w)

		params, ok := r.URL.Query()["id"]
		if !ok || len(params[0]) < 1 {
			fmt.Fprintf(w, "Missing URL parameter.")
			return
		}

		var story Story
		storyId := params[0]

		storyRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE id=?", storyId)
		err := storyRow.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
		utils.QueryErrorCheck(err)

		storyPayload := StoryPayload{
			Id:        story.Id,
			Author:    story.Author,
			Title:     story.Title,
			CreatedAt: story.CreatedAt,
		}

		rows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE story_id=? ORDER BY created_at ASC", storyId)
		utils.QueryErrorCheck(err)

		var post Post

		for rows.Next() {
			err = rows.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.QueryErrorCheck(err)
			storyPayload.Content = append(storyPayload.Content, post)
		}

		payload := &StoryListPayload{
			Status: "OK",
		}
		payload.Data.Stories = append(payload.Data.Stories, storyPayload)

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only GET requests allowed.",
		})
		return
	}
}

func CreateStory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var s Story
		var story Story

		err := utils.DecodeJSONBody(w, r, &s)
		utils.JSONErrorCheck(w, err)

		if len(s.Author) > 0 && len(s.Title) > 0 {
			row, err := models.DB.Query("INSERT INTO stories(author, title) VALUES(?, ?)", s.Author, s.Title)
			utils.QueryErrorCheck(err)

			defer row.Close()

			insertedRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE author=? AND title=?", s.Author, s.Title)
			err = insertedRow.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
			utils.QueryErrorCheck(err)

			storyPayload := StoryPayload{
				Id:        story.Id,
				Author:    story.Author,
				Title:     story.Title,
				CreatedAt: story.CreatedAt,
			}
			payload := StoryListPayload{
				Status: "OK",
			}
			payload.Data.Stories = append(payload.Data.Stories, storyPayload)

			utils.SendJSON(w, payload)
		} else {
			utils.SendJSON(w, ErrorPayload{
				Status: "ERROR",
				Data:   "Empty author/title value.",
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

func DeleteStory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer utils.RecoverErrorCheck(w)

		var s Story
		var story Story

		err := utils.DecodeJSONBody(w, r, &s)
		utils.JSONErrorCheck(w, err)

		deletedRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE id=?", s.Id)
		err = deletedRow.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
		utils.QueryErrorCheck(err)

		row, err := models.DB.Query("DELETE FROM stories WHERE id=?", s.Id)
		utils.QueryErrorCheck(err)

		defer row.Close()

		storyPayload := StoryPayload{
			Id:        story.Id,
			Author:    story.Author,
			Title:     story.Title,
			CreatedAt: story.CreatedAt,
		}
		payload := StoryListPayload{
			Status: "OK",
		}
		payload.Data.Stories = append(payload.Data.Stories, storyPayload)

		utils.SendJSON(w, payload)
	default:
		utils.SendJSON(w, ErrorPayload{
			Status: "ERROR",
			Data:   "Only POST requests allowed.",
		})
		return
	}
}
