package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Logic trước khi bắt đầu vào handler (Before)
		log.Println("Start func - Check from Middleware")
		ctx.Writer.Write([]byte("Start func - Check from Middleware"))

		ctx.Next()

		// Logic sau khi handler hoàn thành (After)
		log.Println("End func - Check from Middleware")
		ctx.Writer.Write([]byte("End func - Check from Middleware"))
	}
}
