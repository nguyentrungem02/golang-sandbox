package v1handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"trungem.com/hoc-golang/utils"
)

type NewsHandler struct {
}

type PostNewsV1Param struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) GetNewsV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if slug == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get news (V1)",
			"slug":    "No news",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get news (V1)",
			"slug":    slug,
		})
	}

}

func (n *NewsHandler) PostNewsV1(ctx *gin.Context) {
	var param PostNewsV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
		})
		return
	}

	// Yeu cau gioi han file < 5MB
	// 1 << 20 = 2^20 = 1048576 B = 1MB
	// 5 << 20 = 5 * 2^20 = 5 * 1048576 B = 5MB
	if image.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Image size is too big (5MB)",
		})
		return
	}

	// os.ModePerm = 0o777 (octal)
	// Có nghĩa: đọc, ghi, thực thi (read, write, execute) cho tất cả mọi người (owner, group, others)
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Create upload dir fail",
		})
		return
	}

	dst := fmt.Sprintf("./uploads/%s", filepath.Base(image.Filename))

	if err := ctx.SaveUploadedFile(image, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Server error save file",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post news (V1)",
		"title":   param.Title,
		"status":  param.Status,
		"image":   image.Filename,
		"path":    dst,
	})
}

func (n *NewsHandler) PostUploadFileNewsV1(ctx *gin.Context) {
	var param PostNewsV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "File is required",
		})
		return
	}

	filename, err := utils.ValidateAndSaveFile(image, "./uploads")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post news (V1)",
		"title":   param.Title,
		"status":  param.Status,
		"image":   filename,
		"path":    "./uploads/" + filename,
	})
}

func (n *NewsHandler) PostUploadMultipleFileNewsV1(ctx *gin.Context) {
	const publicURL = "http://localhost:8080/images/"
	var param PostNewsV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid multipart form",
		})
		return
	}

	images := form.File["images"]
	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "No files to upload",
		})
		return
	}

	var successFiles []string
	var filedFiles []map[string]string
	for _, image := range images {
		filename, err := utils.ValidateAndSaveFile(image, "./uploads")
		if err != nil {
			filedFiles = append(filedFiles, map[string]string{
				"filename": filename,
				"error":    err.Error(),
			})

			continue
		}

		publicImageURL := publicURL + filename
		successFiles = append(successFiles, publicImageURL)
	}

	resp := gin.H{
		"message":       "Post news (V1)",
		"title":         param.Title,
		"status":        param.Status,
		"success_files": successFiles,
	}

	if len(filedFiles) != 0 {
		resp["message"] = "Upload complete with partial error"
		resp["error_files"] = filedFiles
	}

	ctx.JSON(http.StatusCreated, resp)
}
