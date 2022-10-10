package handler

import (
	"fmt"
	"os/exec"
)

func compressVideo(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vf", "scale=1920:1080", "-c:a", "aac", outputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to compress video: %v", err)
	}
	return nil
}
