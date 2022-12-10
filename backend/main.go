package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

const (
	uploadFolder = "uploads"
	resultFolder = "results"
)

func main() {
	// Create necessary directories if they don't exist
	os.MkdirAll(uploadFolder, os.ModePerm)
	os.MkdirAll(resultFolder, os.ModePerm)

	r := gin.Default()

	// CORS setup to allow requests from the frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	r.POST("/data", receiveVideo)
	r.GET("/result", sendVideo)

	r.Run(":8080")
}

func receiveVideo(c *gin.Context) {
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

	filepath := filepath.Join(uploadFolder, filename)
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

func sendVideo(c *gin.Context) {
	filename := c.Query("file")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File name is required"})
		return
	}

	uploadPath := filepath.Join(uploadFolder, filename)
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

func processVideo(filePath string) (string, error) {
	base := filepath.Base(filePath)
	name := base[:len(base)-len(filepath.Ext(base))]
	gazeFileName := name + "_gaze.mp4"
	gazeFilePath := filepath.Join(uploadFolder, gazeFileName)
	compressedFileName := name + "_compressed.mp4"
	compressedFilePath := filepath.Join(resultFolder, compressedFileName)

	// Run gAIze.bat script
	cmd := exec.Command("cmd", "/C", "gAIze.bat", filePath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute gAIze.bat: %v", err)
	}

	// Compress the resulting video
	err := compressVideo(gazeFilePath, compressedFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to compress video: %v", err)
	}

	// Move the compressed video to result folder
	err = os.Rename(compressedFilePath, filepath.Join(resultFolder, compressedFileName))
	if err != nil {
		return "", fmt.Errorf("failed to move compressed video: %v", err)
	}

	// Cleanup
	os.Remove(filePath)
	os.Remove(gazeFilePath)

	return filepath.Join(resultFolder, compressedFileName), nil
}

func compressVideo(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=1920:1080", "-c:a", "aac", outputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to compress video: %v", err)
	}
	return nil
}
