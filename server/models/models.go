package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/michaellu2019/democracy/utils"
)

var DB *sql.DB

type Stories struct {
	Id        int       `name:"id" props:"INT PRIMARY KEY AUTO_INCREMENT"`
	Author    string    `name:"author" props:"TEXT"`
	Title     string    `name:"title" props:"TEXT"`
	CreatedAt time.Time `name:"created_at" props:"TIMESTAMP DEFAULT CURRENT_TIMESTAMP"`
}

type Posts struct {
	Id        int       `name:"id" props:"INT PRIMARY KEY AUTO_INCREMENT"`
	StoryId   int       `name:"story_id" props:"INT NOT NULL"`
	Author    string    `name:"author" props:"TEXT"`
	Text      string    `name:"text" props:"TEXT"`
	Likes     int       `name:"likes" props:"INT DEFAULT 0"`
	Dislikes  int       `name:"dislikes" props:"INT DEFAULT 0"`
	CreatedAt time.Time `name:"created_at" props:"TIMESTAMP DEFAULT CURRENT_TIMESTAMP"`
}

func createTable(t reflect.Type, ctx context.Context) {
	query := "CREATE TABLE IF NOT EXISTS "
	var tableName string
	if t.String() == "models.Stories" {
		tableName = "stories"
	} else if t.String() == "models.Posts" {
		tableName = "posts"
	}

	query += tableName + "("
	for i := 0; i < t.NumField(); i++ {
		columnName := t.Field(i).Tag.Get("name")
		columnProperties := t.Field(i).Tag.Get("props")
		column := columnName + " " + columnProperties

		query += column
		if i < t.NumField()-1 {
			query += ", "
		}
	}
	query += ")"

	res, err := DB.ExecContext(ctx, query)
	utils.QueryErrorCheck(err)

	rows, err := res.RowsAffected()
	utils.QueryErrorCheck(err)
	fmt.Println(rows, "Created table: "+tableName+".")
}

func InitDB() {
	var dbErr error
	DB, dbErr = sql.Open("mysql", "root:Squatifa7.@tcp(127.0.0.1:3306)/golangtestdb?parseTime=true")
	utils.QueryErrorCheck(dbErr)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var stories Stories
	createTable(reflect.TypeOf(stories), ctx)

	var posts Posts
	createTable(reflect.TypeOf(posts), ctx)
}
