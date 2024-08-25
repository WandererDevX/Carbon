package handlers

import (
	"BlogSite/internal/storage/database"
	"BlogSite/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"strconv"
)

func HomeHandler(c *gin.Context) {
	allPosts := database.AllPosts()
	c.HTML(http.StatusOK, "home.html", allPosts)

}

func AddPostHandler(c *gin.Context) {
	uploadPath := "internal/assets/"
	imageName := ""

	postTitle := c.PostForm("postTitle")
	postDescription := c.PostForm("postDescription")
	postImage, err := c.FormFile("postImage")
	if err == nil {
		imageName = uuid.New().String() + filepath.Ext(postImage.Filename)
		imagePath := uploadPath + imageName
		err = c.SaveUploadedFile(postImage, imagePath)
	}
	database.AddPost(postTitle, postDescription, imageName)
}

func AddPostPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "add-post.html", nil)
}

func PostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	selectedPost := database.PostByID(id)

	selectedPost.Description = utils.ConvertMarkdownToHTML(selectedPost.Description)

	c.HTML(http.StatusOK, "post-page.html", selectedPost)
}

func DeletePostHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DeletePost(id)
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
	database.UpdatePost(id, postTitle, postDescription, imageName)
}
