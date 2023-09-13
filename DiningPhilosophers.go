package main

import (
	"fmt"
	"math/rand"
	"time"
)

var philChans [5]chan bool

var forkChansIn [5]chan bool // fork chains

var forkChansOut [5]chan bool // fork chains

/*
Code written by tpep, dhla and habr

Our code does not deadlock, since philosophers no matter what release their forks after a random delay
Random delays ensure no sequencialisation
We also seperated fork channels into input and output to protect against race conditions
(we used to have only one channel per fork, and then race conditions could cause channels to deadlock)
Our code is tested sucessfully with over 600.000 eats per philosopher
*/

func philosopher(i int) {
	var eats int

	for {
		delay := rand.Intn(2) + 1
		time.Sleep(time.Duration(delay) * time.Millisecond)

		rFork := requestFork((i + 1) % 5)
		lFork := requestFork(i)

		if rFork && lFork {
			eats++
			fmt.Println("Philosopher ", i, " eating. Total eats: ", eats)

			time.Sleep(time.Duration(rand.Intn(5)+5) * time.Millisecond) // time it takes for philosopher to eat
			fmt.Println("Philosopher ", i, " thinking")
		}
		if rFork {
			releaseFork((i + 1) % 5)
		}
		if lFork {
			releaseFork(i)
		}
	}
}

func requestFork(i int) bool {
	forkChansIn[i] <- true
	return <-forkChansOut[i]
}

func releaseFork(i int) {
	forkChansIn[i] <- false
}

func fork(i int) {
	var inUse = false

	for {
		if <-forkChansIn[i] { //requesting forks
			if inUse {
				forkChansOut[i] <- false
			} else {
				inUse = true
				forkChansOut[i] <- true
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
		forkChansIn[i] = make(chan bool)
		forkChansOut[i] = make(chan bool)
		go fork(i)
		go philosopher(i)
	}

	for {

	}
}
