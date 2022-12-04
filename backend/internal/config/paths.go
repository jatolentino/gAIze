package config

import (
	"os"
)

const (
	UploadFolder = "uploads"
	ResultFolder = "results"
)

func EnsureDirectories() {
	os.MkdirAll(UploadFolder, os.ModePerm)
	os.MkdirAll(ResultFolder, os.ModePerm)
}
