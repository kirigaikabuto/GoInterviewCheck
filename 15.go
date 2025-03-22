package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func workPool(workNumber int, jobs chan int, results chan string) {
	for job := range jobs {
		time.Sleep(1 * time.Second)
		results <- fmt.Sprintf("job number %d was compelted by worker %d", job, workNumber)
	}
}

func main() {
	r := gin.Default()
	r.GET("/test/", func(c *gin.Context) {
		jobs := make(chan int)
		results := make(chan string)

		for i := 1; i <= 5; i++ {
			go workPool(i, jobs, results)
		}

		go func() {
			for i := 1; i <= 10; i++ {
				jobs <- i
			}
			close(jobs)
		}()
		var response []string
		for i := 0; i < 10; i++ {
			response = append(response, <-results)
		}
		c.JSON(http.StatusOK, gin.H{"message": response})
	})
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
