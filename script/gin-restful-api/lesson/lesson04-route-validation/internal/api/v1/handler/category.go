package v1handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"trungem.com/hoc-golang/utils"
)

type CategoryHandler struct {
}

type GetCategoryByCategoryV1Param struct {
	Category string `uri:"category" binding:"oneof=php golang python"`
}

type PostCategoryV1Param struct {
	Name   string `form:"name" binding:"required"`
	Status int    `form:"status" binding:"required,oneof=1 2"`
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	var param GetCategoryByCategoryV1Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	log.Println("Into GetCategoryByCategoryV1")

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get category by category (V1)",
		"category": param.Category,
	})
}

func (c *CategoryHandler) PostCategoryV1(ctx *gin.Context) {
	var param PostCategoryV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post category (V1)",
		"name":    param.Name,
		"status":  param.Status,
	})
}
