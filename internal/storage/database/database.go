package database

import (
	"Carbon/internal/models"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
)

func CreateTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		image TEXT
	);
	`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		slog.Error("Failed to create table", "Error", err)
		return
	}
}

func GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "internal/storage/database/posts.db")
	if err != nil {
		slog.Error("Failed to open database", "Error", err)
		return nil
	}
	return db
}

func AllPosts() ([]models.Post, error) {
	db := GetDB()
	query := "SELECT * FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allPosts []models.Post
	for rows.Next() {
		post := models.Post{}
		if err = rows.Scan(&post.ID, &post.Title, &post.Description, &post.Image); err != nil {
			return nil, err
		}
		allPosts = append(allPosts, post)
	}
	return allPosts, nil
}

func PostByID(id int) (models.Post, error) {
	db := GetDB()
	var post models.Post
	query := `SELECT * FROM posts WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Description, &post.Image)
	if err != nil {
		return post, err
	}
	return post, nil
}

func AddPost(title string, description string, imageName string) error {
	db := GetDB()

	relativeImagePath := ""
	if imageName != "" {
		relativeImagePath = "/assets/" + imageName
	}

	query := `INSERT INTO posts (title, description, image) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, title, description, relativeImagePath)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(id int) error {
	db := GetDB()
	query := `DELETE FROM posts WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(id int, newTitle string, newDescription string, newImage string) error {
	db := GetDB()
	query := `UPDATE posts SET (title, description, image) = ($1, $2, $3) WHERE id = $4`
	_, err := db.Exec(query, newTitle, newDescription, newImage, id)
	if err != nil {
		return err
	}
	return nil
}
