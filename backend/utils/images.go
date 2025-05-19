package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveUploadedImage(file multipart.File, handler *multipart.FileHeader) (string, error) {
	defer file.Close()

	// Generate a unique filename using the current timestamp
	filename := time.Now().Format("20060102150405") + "_" + filepath.Base(handler.Filename)
	saveDir := "static/uploads/"  // Corrected path

	// Make sure the directory exists
	if err := os.MkdirAll(saveDir, 0o755); err != nil {
		return "", err
	}

	savePath := filepath.Join(saveDir, filename)

	dst, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	// Return the relative path from the web server's perspective (for frontend use)
	return "/static/uploads/" + filename, nil
}
