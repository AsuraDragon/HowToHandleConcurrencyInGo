package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func sendMsg(msg string) {
	defer wg.Done()
	pause()
	fmt.Println(msg)
}

func main() {
	wg.Add(4)
	defer wg.Wait()
	go func(msg string) {
		defer wg.Done()
		pause()
		fmt.Println(msg)
	}("Ella")

	go sendMsg("No")  //Async
	go sendMsg("Te")  //Async
	go sendMsg("Ama") //Async

	time.Sleep(2 * time.Second)

}
