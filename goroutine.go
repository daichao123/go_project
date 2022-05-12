package main

import (
	"fmt"
	"time"
)

func main() {
	chanA := make(chan int, 5)
	chanB := make(chan int, 4)
	chanB <- 1
	chanB <- 2
	chanB <- 3
	chanB <- 4
	chanC := make(chan int, 3)
	chanC <- 1
	chanC <- 2
	chanC <- 3

	go TaskA(chanA)
	go TaskB(chanB)
	go TaskC(chanC)

	time.Sleep(time.Second * 10)

	fmt.Println("test over !")
}

func TaskA(ch chan int) {
	for {
		fmt.Println("TaskA!!!")
		time.Sleep(time.Second * 1)
		ch <- 123
	}

}

func TaskB(ch chan int) {
	for {
		fmt.Println("TaskB!!!")
		<-ch
	}

}

func TaskC(ch chan int) {
	for {
		fmt.Println("TaskC!!!")
		<-ch
	}
}
