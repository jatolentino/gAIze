package handler

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"gaizeapi/internal/config"
)

func ReceiveVideo(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	defer file.Close()

	filename := filepath.Base(fileHeader.Filename)
	if filepath.Ext(filename) != ".mp4" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format"})
		return
	}

	filepath := filepath.Join(config.UploadFolder, filename)
	dst, err := os.Create(filepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": filename})
}
