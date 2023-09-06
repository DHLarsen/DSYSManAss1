package main

import "fmt"

var leftChans [5]chan bool
var rightChans [5]chan bool

func status(string status){
	fmt.Println(status)
}

func philosopher(i int) {
	var leftChan = leftChans[i];
	var rightChan = rightChans[i];
	
	if ()

}

func fork(i int) {

}
func create(){
	// Creates 5 philosophers, 5 forks and channels for the right and the left fork
	for i := 0; i < 5; i++ {
		leftChans[i] = make(chan bool)
		rightChans[i] = make(chan bool)
		go philosopher(i)
		go fork(i)
	}
}

func main() {
	fmt.Println("Hello, World!")
	create()

	

}
