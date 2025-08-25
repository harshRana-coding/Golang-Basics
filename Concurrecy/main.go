package main

import (
	"fmt"
	"time"
)

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)
	// dones[0] = make(chan bool)
	go greet("Nice to meet you !",done)
	// dones[1] = make(chan bool)
	go greet("How are you ?",done)
	// dones[2] = make(chan bool)
	go slowGreet("Sorry for being late !",done)
	// dones[3] = make(chan bool)
	go greet("Whats new about you ?",done)

	// for _, done := range dones {
	// 	<- done
	// }
	for range done {
		// fmt.Print(doneChan)
	}
}

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	done <- true
	close(done)
}
