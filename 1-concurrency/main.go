package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make(chan int)
	squares := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer close(nums)
		defer wg.Done()
		slice := make([]int, 10)
		for i := 0; i < 10; i++ {
			slice[i] = r.Intn(101)
		}

		for _, num := range slice {
			nums <- num
		}
	}()

	go func() {
		defer close(squares)
		defer wg.Done()

		for num := range nums {
			squares <- num * num
		}
	}()

	for square := range squares {
		fmt.Print(square, " ")
	}
	wg.Wait()

}
