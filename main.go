package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pause() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func sendMsg(msg string) {
	pause()
	fmt.Println(msg)
}

func main() {
	sendMsg("Hello") //sync

	go sendMsg("Ella") //Async
	go sendMsg("No")   //Async
	go sendMsg("Te")   //Async
	go sendMsg("Ama")  //Async

	sendMsg("World!") //sync

	time.Sleep(2 * time.Second)

}
