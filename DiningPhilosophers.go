package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock1 sync.Mutex

var lock2 sync.Mutex

var philChans [5]chan bool

var forkChans [5]chan bool // fork chains

func philosopher(i int) {
	var eats int

	for {
		delay := rand.Intn(10)
		time.Sleep(time.Duration(delay) * time.Millisecond)

		rFork := requestFork((i + 1) % 5)
		lFork := requestFork(i)

		if rFork {
			fmt.Println("Philosopher ", i, " picked up the right fork")
		}
		if lFork {
			fmt.Println("Philosopher ", i, " picked up the left fork")
		}

		if rFork && lFork {
			eats++
			fmt.Println("Philosopher ", i, " eating. Total eats: ", eats)
			//time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second) // time it takes for philosopher to eat
			//time.Sleep(100 * time.Millisecond)
			fmt.Println("Philosopher ", i, " thinking")
		}
		lock1.Lock()
		if rFork {
			fmt.Println("Philosopher ", i, " released the right fork")
			releaseFork((i + 1) % 5)
		}
		if lFork {
			fmt.Println("Philosopher ", i, " released the left fork")
			releaseFork(i)
		}
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		lock1.Unlock()
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
