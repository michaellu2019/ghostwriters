package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaellu2019/ghostwriters/models"
	"github.com/michaellu2019/ghostwriters/utils"
)

// structs for JSON objects
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

		// get all rows in the stories table
		storyRows, err := models.DB.Query("SELECT id, author, title, created_at FROM stories")
		utils.ErrorCheck(err)

		var story Story
		storyMap := make(map[int]*StoryPayload)

		// add all story row column data to a map that connects story IDs to a story payload data struct for the particular story ID
		for storyRows.Next() {
			err = storyRows.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
			utils.ErrorCheck(err)

			storyMap[story.Id] = &StoryPayload{
				Id:        story.Id,
				Author:    story.Author,
				Title:     story.Title,
				CreatedAt: story.CreatedAt,
			}
		}

		// get all rows in the posts table
		postRows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts ORDER BY created_at ASC")
		utils.ErrorCheck(err)

		var post Post

		// add all post row column data to the map that connects story IDs to a story payload data struct for the particular post's story ID
		for postRows.Next() {
			err = postRows.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.ErrorCheck(err)
			storyMap[post.StoryId].Content = append(storyMap[post.StoryId].Content, post)
		}

		// go through the map of story IDs and data and add the content to the payload data struct
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

		// get all rows in the stories table
		rows, err := models.DB.Query("SELECT id, author, title, created_at FROM stories")
		utils.ErrorCheck(err)

		var story Story
		payload := &StoryListPayload{
			Status: "OK",
		}

		// add all story row column data to the payload data struct
		for rows.Next() {
			err = rows.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
			utils.ErrorCheck(err)

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

		// parse the URL for the provided story ID
		params, ok := r.URL.Query()["id"]
		if !ok || len(params[0]) < 1 {
			fmt.Fprintf(w, "Missing URL parameter.")
			return
		}

		storyId := params[0]
		var post Post
		var story Story

		// get the row in the posts table with the specified story ID
		storyRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE id=?", storyId)

		// add all story row column data to the data struct
		err := storyRow.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
		utils.ErrorCheck(err)

		storyPayload := StoryPayload{
			Id:        story.Id,
			Author:    story.Author,
			Title:     story.Title,
			CreatedAt: story.CreatedAt,
		}

		// get all rows in the stories table with the specified story ID
		rows, err := models.DB.Query("SELECT id, story_id, author, text, likes, dislikes, created_at FROM posts WHERE story_id=? ORDER BY created_at ASC", storyId)
		utils.ErrorCheck(err)

		// add all post row column data to the payload data struct
		for rows.Next() {
			err = rows.Scan(&post.Id, &post.StoryId, &post.Author, &post.Text, &post.Likes, &post.Dislikes, &post.CreatedAt)
			utils.ErrorCheck(err)
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

		// make sure no empty data was sent
		if len(s.Author) > 0 && len(s.Title) > 0 {
			// insert a new row with the specified values in the JSON object into the stories table
			row, err := models.DB.Query("INSERT INTO stories(author, title) VALUES(?, ?)", s.Author, s.Title)
			utils.ErrorCheck(err)

			defer row.Close()

			// get the row column data that was just inserted and add it to the payload
			insertedRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE author=? AND title=?", s.Author, s.Title)
			err = insertedRow.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
			utils.ErrorCheck(err)

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

		// get the row in the stories table that has the specified story ID so that its data can be returned when deleted
		deletedRow := models.DB.QueryRow("SELECT id, author, title, created_at FROM stories WHERE id=?", s.Id)
		err = deletedRow.Scan(&story.Id, &story.Author, &story.Title, &story.CreatedAt)
		utils.ErrorCheck(err)

		// delete the row with the specified story ID from the stories table
		row, err := models.DB.Query("DELETE FROM stories WHERE id=?", s.Id)
		utils.ErrorCheck(err)

		defer row.Close()

		// return the data of the deleted story
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
