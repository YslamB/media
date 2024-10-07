package utils

import (
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"media/pkg/config"
	"media/pkg/database"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

// ResizeImage resizes an image to the specified width and saves it as .webp
func ResizeImage(imagePath string, newWidth uint) (int, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return 500, errors.New("failed to open image file: " + err.Error())
	}
	defer file.Close()

	// Detect the image format and decode it
	var img image.Image
	// img, err = webp.Decode(file)
	// if err != nil {
	// 	return 500, errors.New("failed to decode WebP image: " + err.Error())
	// }

	switch {
	case strings.HasSuffix(imagePath, ".jpg"), strings.HasSuffix(imagePath, ".jpeg"):
		img, err = jpeg.Decode(file)
		if err != nil {
			return 500, errors.New("failed to decode JPEG image: " + err.Error())
		}
	case strings.HasSuffix(imagePath, ".png"):
		img, err = png.Decode(file)
		if err != nil {
			return 500, errors.New("failed to decode PNG image: " + err.Error())
		}
	case strings.HasSuffix(imagePath, ".webp"):
		img, err = webp.Decode(file)
		if err != nil {
			return 500, errors.New("failed to decode WebP image: " + err.Error())
		}
	default:
		return 400, errors.New("unsupported image format")
	}

	// Resize the image using the specified width, preserving the aspect ratio

	resizedImg := resize.Resize(newWidth, 0, img, resize.Lanczos3)

	// Create the output .webp file
	outputPath := strings.TrimSuffix(imagePath, filepath.Ext(imagePath)) + ".webp"
	outFile, err := os.Create(outputPath)
	if err != nil {
		return 400, errors.New("failed to create output file: " + err.Error())
	}
	defer outFile.Close()

	// Encode the resized image as .webp
	if err := webp.Encode(outFile, resizedImg, &webp.Options{Lossless: true}); err != nil {
		return 500, errors.New("failed to encode and save .webp image: " + err.Error())
	}

	os.Remove(imagePath)

	return 200, nil
}

func ConvertToHLS(filepath, filename, runType string) error {

	if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	if runType == "film" {
		cmd := exec.Command(
			"ffmpeg",
			"-i", filepath+filename,
			"-c:v", config.ENV.HLS_RUN_ON,
			"-flags", "+cgop",
			"-g", "30",
			"-hls_time", "10",
			"-hls_playlist_type", "event",
			filepath+removeExt(filename)+"HLS.m3u8",
		)

		err := cmd.Run()

		if err != nil {
			return err
		}

		updateStatus(filepath[1:]+removeExt(filename)+"HLS.m3u8", "films")
		err = os.Remove(filepath + filename)
		return err
	}

	cmd := exec.Command("ffmpeg",
		"-i", filepath+filename,
		"-c:a", "aac",
		"-b:a", "128k",
		"-hls_time", "10",
		"-hls_playlist_type", "event",
		filepath+removeExt(filename)+"HLS.m3u8",
	)

	err := cmd.Run()

	if err == nil {
		updateStatus(filepath[1:]+removeExt(filename)+"HLS.m3u8", "musics")
	}

	err = os.Remove(filepath + filename)
	return err
}

func removeExt(filename string) string {
	base := filepath.Base(filename)
	ext := filepath.Ext(base)
	return base[:len(base)-len(ext)]
}

func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()

	if err != nil {
		return err
	}

	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func updateStatus(path, table string) {
	fmt.Println("UpdateStatus:", path)
	_, err := database.DB.Exec(context.Background(), "update "+table+" set status = true where path = $1", path)
	fmt.Println("err")
	fmt.Println(err)
}
