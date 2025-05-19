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

	// Sanitize the filename to avoid any directory traversal issues
	filename := time.Now().Format("20060102150405") + "_" + filepath.Base(handler.Filename)
	saveDir := "/backend/static/uploads/"

	// Make sure the directory exists before creating the file
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return "", err
	}

	savePath := saveDir + filename

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
	return "/uploads/" + filename, nil
}
