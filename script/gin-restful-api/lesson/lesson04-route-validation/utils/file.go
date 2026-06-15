package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var allowExts = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}
var allowMime = map[string]bool{
	"image/jpg":  true,
	"image/jpeg": true,
	"image/png":  true,
}

const maxSize = 5 << 20

func ValidateAndSaveFile(fileHeader *multipart.FileHeader, uploadDir string) (string, error) {
	// 1. Check extension file
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowExts[ext] {
		return fileHeader.Filename, errors.New("unsupported file extension")
	}

	// 2. Check size file
	if fileHeader.Size > maxSize {
		return fileHeader.Filename, errors.New("file too large (max 5MB)")
	}

	// 3. Check file type
	// 3.1 Open file
	file, err := fileHeader.Open()
	if err != nil {
		return fileHeader.Filename, errors.New("cannot open file")
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	// 3.2 Create buffer and Read file
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return fileHeader.Filename, errors.New("cannot read file")
	}
	// 3.3 Check mimeType file
	mimeType := http.DetectContentType(buffer)
	if !allowMime[mimeType] {
		return fileHeader.Filename, fmt.Errorf("invalid mime type %s", mimeType)
	}

	// 4. Rename file
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// 5. Create folder if not exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return fileHeader.Filename, errors.New("cannot create upload directory")
	}

	// 6. uploadDir "./uploads" + filename "abc.png"
	savePath := filepath.Join(uploadDir, filename)
	if err := saveFile(fileHeader, savePath); err != nil {
		return fileHeader.Filename, errors.New("cannot save file")
	}

	return filename, nil
}

func saveFile(fileHeader *multipart.FileHeader, destination string) error {
	// Open file
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	// Create file empty
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)

	// Copy content file src to out
	_, err = io.Copy(out, src)

	return err
}
