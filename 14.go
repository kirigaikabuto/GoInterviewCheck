package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func processTask(ctx context.Context) string {
	select {
	case <-time.After(1 * time.Second):
		return "Task was completed"
	case <-ctx.Done():
		return "Task shutdown"
	}
}

func main() {
	r := gin.Default()
	r.GET("/data/", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()
		result := processTask(ctx)
		c.JSON(http.StatusOK, gin.H{"message": result})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
