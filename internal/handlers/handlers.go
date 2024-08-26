package handlers

import (
	"Carbon/internal/storage/database"
	"Carbon/internal/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
)

func HomeHandler(c *gin.Context) {
	allPosts, err := database.AllPosts()
	if err != nil {
		slog.Error("Failed to retrieve posts:", "error", err)
		return
	}
	c.HTML(http.StatusOK, "home.html", allPosts)

}

func AddPostHandler(c *gin.Context) {
	postTitle := c.PostForm("postTitle")
	postDescription := c.PostForm("postDescription")
	postImage, err := c.FormFile("postImage")

	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		slog.Error("Failed to retrieve post image:", "error", err)
	}
	imageName, err := utils.SaveImage(c, postImage)
	if err != nil {
		slog.Error("Failed to save image:", "error", err)
		return
	}
	err = database.AddPost(postTitle, postDescription, imageName)
	if err != nil {
		slog.Error("Failed to add a post:", "error", err)
	}
}

func AddPostPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "add-post.html", nil)
}

func PostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedPost, err := database.PostByID(id)
	if err != nil {
		slog.Error("Failed to get post with id:", id, "error:", err)
	}
	selectedPost.Description, err = utils.ConvertMarkdownToHTML(selectedPost.Description)
	if err != nil {
		slog.Error("Failed to convert markdown to HTML:", err)
	}

	c.HTML(http.StatusOK, "post-page.html", selectedPost)
}

func DeletePostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := database.DeletePost(id)
	if err != nil {
		slog.Error("Failed to delete post with id:", id, "error:", err)
	}
}

func UpdatePostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	postTitle := c.PostForm("postTitle")
	postDescription := c.PostForm("postDescription")
	postImage, err := c.FormFile("postImage")
	uploadPath := "internal/assets/"
	imageName := ""

	if err == nil {
		imageName = uuid.New().String() + filepath.Ext(postImage.Filename)
		imagePath := uploadPath + imageName
		err = c.SaveUploadedFile(postImage, imagePath)
	}
	err = database.UpdatePost(id, postTitle, postDescription, imageName)
	if err != nil {
		slog.Error("Failed to update post with id:", id, "error:", err)
	}
}
