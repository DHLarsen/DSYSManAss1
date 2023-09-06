package main

import "fmt"

func philosopher(i int) {

}

func fork(i int) {

}

func main() {
	fmt.Println("Hello, World!")

	var chans [5]chan bool
	for i := 0; i < 5; i++ {
		go philosopher(i)
		go fork(i)
		chans[i] = make(chan bool)
	}

}
