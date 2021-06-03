package models

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/michaellu2019/ghostwriters/utils"
)

var DB *sql.DB

// structs for database tables
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
	// create a MySQL query to create the speciifed table
	query := "CREATE TABLE IF NOT EXISTS "
	var tableName string
	if t.String() == "models.Stories" {
		tableName = "stories"
	} else if t.String() == "models.Posts" {
		tableName = "posts"
	}

	// parse the struct field names and tags to generate a MySQL query to create the table
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

	// execute the created MySQL query
	res, err := DB.ExecContext(ctx, query)
	utils.ErrorCheck(err)

	rows, err := res.RowsAffected()
	utils.ErrorCheck(err)
	fmt.Println(rows, "Created table: "+tableName+".")
}

func InitDB() {
	// open a database connection
	var dbErr error

	const dbLang = "mysql"

	err := godotenv.Load("db_config.env")
	utils.ErrorCheck(err)

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUsername, dbPassword, dbURL, dbName)

	DB, dbErr = sql.Open(dbLang, dbConn)
	utils.ErrorCheck(dbErr)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	var stories Stories
	createTable(reflect.TypeOf(stories), ctx)

	var posts Posts
	createTable(reflect.TypeOf(posts), ctx)
}
