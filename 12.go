package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/async/", func(context *gin.Context) {
		ctx := context.Copy()
		go func() {
			fmt.Printf("Goroutine start on %s\n", ctx.Request.URL.Path)
		}()
		context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
