package database

import (
	"BlogSite/internal/models"
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

func AllPosts() []models.Post {
	db := GetDB()
	query := "SELECT * FROM posts"
	rows, err := db.Query(query)
	if err != nil {
		slog.Error("Failed to querying all posts", "Error", err)
		return nil
	}
	defer rows.Close()

	var allPosts []models.Post
	for rows.Next() {
		post := models.Post{}
		if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.Image); err != nil {
			slog.Error("Failed to scanning row", "Error", err)
			return nil
		}
		allPosts = append(allPosts, post)
	}
	return allPosts
}

func PostByID(id int) models.Post {
	db := GetDB()
	var post models.Post
	query := `SELECT * FROM posts WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&post.ID, &post.Title, &post.Description, &post.Image)
	if err != nil {
		slog.Error("Failed to querying post by id", "Error", err)
		return post
	}
	return post
}

func AddPost(title string, description string, imageName string) {
	db := GetDB()

	relativeImagePath := ""
	if imageName != "" {
		relativeImagePath = "/assets/" + imageName
	}

	query := `INSERT INTO posts (title, description, image) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, title, description, relativeImagePath)
	if err != nil {
		slog.Error("Failed to inserting post into database", "Error", err)
	}
}

func DeletePost(id int) {
	db := GetDB()
	query := `DELETE FROM posts WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		slog.Error("Failed to deleting post", "Error", err)
	}
}

func UpdatePost(id int, newTitle string, newDescription string, newImage string) {
	db := GetDB()
	query := `UPDATE posts SET (title, description, image) = ($1, $2, $3) WHERE id = $4`
	_, err := db.Exec(query, newTitle, newDescription, newImage, id)
	if err != nil {
		slog.Error("Failed to update post", "Error", err)
	}
}
