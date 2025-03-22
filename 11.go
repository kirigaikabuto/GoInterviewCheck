package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func makeSomeWork(workerNumber int, orders chan int, finishedOrders chan string, wg *sync.WaitGroup) {
	for order := range orders {
		fmt.Printf("Worker %d start his work with order number %d\n", workerNumber, order)
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		finishedOrders <- fmt.Sprintf("Worker %d finished his work with order number %d\n", workerNumber, order)
	}
	defer wg.Done()
}

func main() {
	//🔥 Реальная задача для горутин: Обработчик заказов в интернет-магазине
	//Представим, что у нас есть интернет-магазин, где покупатели делают заказы. Нам нужно:
	//✅ Принимать заказы в один канал.
	//✅ Обрабатывать их параллельно с задержкой (как будто работники готовят товар).
	//✅ Записывать выполненные заказы в другой канал.
	//✅ Остановить обработку после всех заказов.
	orders := make(chan int, 3)
	finished := make(chan string, 3)
	wg := sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go makeSomeWork(i, orders, finished, &wg)
	}
	go func() {
		for i := 1; i < 10; i++ {
			orders <- i
		}
		close(orders)
	}()
	go func() {
		wg.Wait()
		close(finished)
	}()
	for work := range finished {
		fmt.Println(work)
	}
}
