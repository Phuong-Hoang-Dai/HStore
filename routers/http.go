package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpHandler() {
	router := gin.Default()
	{
		v1 := router.Group("/api/v1")
		{
			product := v1.Group("/product")
			product.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"sucess": "hahahaha",
				})
			})
		}
	}
	router.Run()
}
