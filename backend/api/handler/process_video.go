package handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"gaizeapi/internal/config"
)

func processVideo(filePath string) (string, error) {
	base := filepath.Base(filePath)
	name := base[:len(base)-len(filepath.Ext(base))]
	gazeFileName := name + "_gaze.mp4"
	gazeFilePath := filepath.Join(config.UploadFolder, gazeFileName)
	compressedFileName := name + "_compressed.mp4"
	compressedFilePath := filepath.Join(config.ResultFolder, compressedFileName)

	// Run gAIze.bat script
	cmd := exec.Command("cmd", "/C", "scripts/gAIze.bat", filePath)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to execute gAIze.bat: %v", err)
	}

	// Compress the resulting video
	err := compressVideo(gazeFilePath, compressedFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to compress video: %v", err)
	}

	// Move the compressed video to result folder
	err = os.Rename(compressedFilePath, filepath.Join(config.ResultFolder, compressedFileName))
	if err != nil {
		return "", fmt.Errorf("failed to move compressed video: %v", err)
	}

	// Cleanup
	os.Remove(filePath)
	os.Remove(gazeFilePath)

	return filepath.Join(config.ResultFolder, compressedFileName), nil
}
