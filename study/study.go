package main

import (
	"time"
	"fmt"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() { 
	// go processando() // T2
	// go processando() // T3
	// processando() // T1

	canal := make(chan int)

	go func() {
		// canal <- 1 // T2 - incrementar 1 no canal
		for i := 0; i < 10; i++ {
			canal <- 1
			fmt.Println("Jogou no canal", i)
		}
	}()

	// fmt.Println(<-canal) // esvaziar canal

	// time.Sleep(time.Second * 2) 

	// for x := range canal {
	// 	fmt.Println("Recebeu do canal", x)
	// 	time.Sleep(time.Second)
	// }

	go worker(canal, 1)
	worker(canal, 2)
}

func worker(canal chan int, workerId int) { 
	for {
		fmt.Println("Recebeu do canal", <-canal, "no worker", workerId)
		time.Sleep(time.Second)
	}
}