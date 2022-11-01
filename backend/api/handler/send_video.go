package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
	"github.com/gin-gonic/gin"
	"gaizeapi/internal/config"
)

func SendVideo(c *gin.Context) {
	filename := c.Query("file")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File name is required"})
		return
	}

	uploadPath := filepath.Join(config.UploadFolder, filename)
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	resultPath, err := processVideo(uploadPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing video"})
		return
	}

	// Schedule file deletion after sending the response
	go func(path string) {
		time.Sleep(5 * time.Second) // Adjust delay as necessary
		os.Remove(path)
	}(resultPath)

	c.File(resultPath)
}
