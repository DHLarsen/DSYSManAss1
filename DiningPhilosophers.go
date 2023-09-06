package main

import (
	"fmt"
	"math/rand"
	"time"
)

var philChans [5]chan bool

var forkChans [5]chan bool

func philosopher(i int) {

	for {
		rFork := requestFork((i + 1) % 5)
		lFork := requestFork(i)

		if rFork != lFork {
			delay := rand.Intn(10)
			time.Sleep(time.Duration(delay) * time.Millisecond)
			rFork = requestFork((i + 1) % 5)
			lFork = requestFork(i)
		}
		if rFork && lFork {
			fmt.Println("Philosopher ", i, " eating")
			time.Sleep(1 * time.Second)
		} else {
			if rFork {
				releaseFork((i + 1) % 5)
			}
			if lFork {
				releaseFork(i)
			}
			time.Sleep(4 * time.Millisecond)
		}
	}
}

func requestFork(i int) bool {
	forkChans[i] <- true
	return <-forkChans[i]
}

func releaseFork(i int) {
	forkChans[i] <- false
}

func fork(i int) {
	var inUse = false

	for {
		if <-forkChans[i] { //requesting forks
			if inUse {
				forkChans[i] <- false
			} else {
				inUse = true
				forkChans[i] <- true
			}
		} else { //releasing forks
			inUse = false
		}
	}
}

func main() {
	// Creates 5 philosophers, 5 forks and channels for the right and the left fork
	for i := 0; i < 5; i++ {
		philChans[i] = make(chan bool)
		forkChans[i] = make(chan bool)
		go fork(i)
		go philosopher(i)
	}

	for {

	}
}
