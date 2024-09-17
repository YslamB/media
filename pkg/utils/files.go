package utils

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"media/pkg/config"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

var extensions map[string]bool = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"svg":  true,
	"webp": true,
	"mp4":  true,
}

func WriteImage(c *gin.Context, dir string) string {
	image, header, _ := c.Request.FormFile("image")

	if image == nil {
		return ""
	}

	splitedFileName := strings.Split(header.Filename, ".")
	extension := splitedFileName[len(splitedFileName)-1]

	if extension == "webp" || extension == "svg" || extension == "jpeg" ||
		extension == "jpg" || extension == "png" {

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, image)
		os.WriteFile(
			config.ENV.UPLOAD_PATH+dir+header.Filename,
			buf.Bytes(), os.ModePerm,
		)

		return header.Filename
	}

	return ""
}

func SaveFiles(c *gin.Context) ([]string, error) {
	form, _ := c.MultipartForm()

	if form == nil {
		return nil, errors.New("didn't upload the files")
	}

	files := form.File["files"]

	if len(files) == 0 {
		return nil, errors.New("must load minimum 1 file")
	}

	var filePaths []string
	var fileNames []string
	var video = 0
	var images = 0

	for _, file := range files {
		const maxFileSize = 50 * 1024 * 1024 // 50MB

		if file.Size > maxFileSize {
			return nil, fmt.Errorf("file %s is too large", file.Filename)
		}
		splitedFileName := strings.Split(file.Filename, ".")
		extension := splitedFileName[len(splitedFileName)-1]

		extensionExists := extensions[extension]

		if !extensionExists {
			return nil, fmt.Errorf("this file (extension) is forbidden: .%s", extension)
		}

		if extension == "mp4" {
			video += 1
		} else {
			images += 1
		}

		if video > 1 || images > 5 {
			return nil, fmt.Errorf("trying to upload %v video and %v images", video, images)
		}

		fileNames = append(fileNames, uuid.NewString()+"."+extension)
	}

	for index, file := range files {
		readerFile, _ := file.Open()

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, readerFile)
		err := os.WriteFile(
			config.ENV.UPLOAD_PATH+"orders/"+fileNames[index],
			buf.Bytes(),
			os.ModePerm,
		)

		if err != nil {
			return nil, err
		}

		if strings.ToLower(filepath.Ext(fileNames[index])) != ".mp4" {

			err = ResizeImage(config.ENV.UPLOAD_PATH+"orders/"+fileNames[index], 700)
			if err != nil {
				return nil, err
			}

		}

		filePaths = append(filePaths, "/uploads/orders/"+fileNames[index])
	}

	return filePaths, nil
}

func ResizeImage(imagePath string, width uint) error {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Decode the image
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}

	// Resize the image to the specified width
	newImage := resize.Resize(width, 0, img, resize.Lanczos3)

	// Close the original file so it can be deleted
	file.Close()

	// Delete the original image file
	err = os.Remove(imagePath)
	if err != nil {
		return fmt.Errorf("failed to delete original image: %w", err)
	}

	// Create a new file with the same name
	out, err := os.Create(imagePath)
	if err != nil {
		return fmt.Errorf("failed to create new image file: %w", err)
	}
	defer out.Close()

	// Encode and save the resized image
	switch format {
	case "jpeg":
		err = jpeg.Encode(out, newImage, nil)
	case "png":
		err = png.Encode(out, newImage)
	default:
		return fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return fmt.Errorf("failed to save resized image: %w", err)
	}

	return nil
}

func ConvertToHLS(filepath, filename, runType string) error {
	// Define the output directory and HLS file names

	// Create the output directory if it doesn't exist
	if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}
	if runType == "film" {
		// Define the FFmpeg command to convert video to HLS and resize
		cmd := exec.Command(
			"ffmpeg",
			"-i", filepath+filename,
			// "-c:v", *runType,
			// "-c:v", "h264_nvenc", //work on gpu //
			// "-c:v", "hevc_nvenc", //work on gpu //
			"-c:v", "h264_videotoolbox", //"hevc_videotoolbox", //"h264", //cpu
			// -hls_segment_filename 'segment_%03d.ts'
			"-flags", "+cgop",
			"-g", "30",
			"-hls_time", "10",
			"-hls_playlist_type", "event",
			filepath+filename+"HLS.m3u8",
		)

		// Run the FFmpeg command and capture any errors
		err := cmd.Run()
		if err != nil {
			return err
		}
		err = os.Remove(filepath + filename)
		return err
	}

	cmd := exec.Command("ffmpeg",
		"-i", filepath+filename, // Input MP3 file
		"-c:a", "aac", // Use AAC audio codec
		"-b:a", "128k", // Set audio bitrate to 128kbps
		"-hls_time", "10", // Segment duration in seconds
		"-hls_playlist_type", "event",
		filepath+filename+"HLS.m3u8", // Output HLS playlist (.m3u8)
	)

	err := cmd.Run()
	fmt.Println("HLS conversion completed successfully!:::", err)
	err = os.Remove(filepath + filename)

	return err
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

func GetType(contentType string) string {
	switch contentType {
	case "video/mp4":
		return "video"
	case "audio/mpeg":
		return "audio"
	case "application/pdf":
		return "book"
	}
	return "other"
}
