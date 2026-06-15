package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello trung em",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Danh sach thanh vien",
		})
	})

	r.GET("/users/:user_id", func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		ctx.JSON(200, gin.H{
			"data":    "Thong tin thanh vien",
			"user_id": userId,
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Danh sach san pham",
		})
	})

	r.GET("/products/:product_name", func(ctx *gin.Context) {
		productName := ctx.Param("product_name")

		price := ctx.Query("price")
		color := ctx.Query("color")

		ctx.JSON(200, gin.H{
			"data":         "Thong tin san pham",
			"product_name": productName,
			"price":        price,
			"color":        color,
		})
	})

	r.Run(":8080")
}
