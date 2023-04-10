package file_saver

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

func SaveFile(ctx context.Context, savePath, dir string, fileHeader *multipart.FileHeader) (string, error) {
	savePath += dir

	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		return "", fmt.Errorf("file saver: save file: %w", err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("file saver: save file: %w", err)
	}
	defer file.Close()

	savePath += "/" + fileHeader.Filename

	out, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("file saver: save file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", fmt.Errorf("file saver: save file: %w", err)
	}

	return savePath[1:], nil
}
