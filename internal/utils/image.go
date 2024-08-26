package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"path/filepath"
)

var uploadPath = "internal/assets/"

func SaveImage(c *gin.Context, postImage *multipart.FileHeader) (string, error) {
	if postImage == nil {
		return "", nil
	}

	imageName := uuid.New().String() + filepath.Ext(postImage.Filename)
	imagePath := uploadPath + imageName
	err := c.SaveUploadedFile(postImage, imagePath)
	if err != nil {
		return "", err
	}
	return imageName, nil
}
