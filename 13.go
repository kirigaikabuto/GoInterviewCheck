package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

func processData(number int) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("process number %d finished", number)
}

func main() {
	router := gin.Default()
	router.GET("/getData/", func(context *gin.Context) {
		results := make(chan string, 3)
		wg := sync.WaitGroup{}
		for i := 1; i <= 3; i++ {
			wg.Add(1)
			go func(val int) {
				defer wg.Done()
				results <- processData(val)
			}(i)
		}
		go func() {
			wg.Wait()
			close(results)
		}()
		var response []string
		for newVal := range results {
			response = append(response, newVal)
		}
		context.JSON(http.StatusOK, gin.H{"response": response})
	})
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
