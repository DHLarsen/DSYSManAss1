package main

import (
	"fmt"
	"math/rand"
	"time"
)

var philChans [5]chan bool

var forkChans [5]chan bool // fork chains

/*
The program can end up in a loop, where the same philosopher keeps picking up the same fork, a specific philosopher never picks up any forks, two philosophers are the only ones eating, or something reminiscent of this
This happens when print statements are called frequently
To see the code run practically run forever, comment out the active print statements, and use the commented out % print statement in philosopher
With this print statement, we have tested our code with over 30.000 eats per philosopher
*/

func philosopher(i int) {
	var eats int

	for {
		delay := rand.Intn(10) + 2
		time.Sleep(time.Duration(delay) * time.Millisecond)

		rFork := requestFork((i + 1) % 5)
		lFork := requestFork(i)

		if rFork && lFork {
			eats++
			fmt.Println("Philosopher ", i, " eating. Total eats: ", eats)
			/*
				if eats%100 == 0 { // use this instead of the other print statements to see the code run practically forever
					fmt.Println("Philosopher ", i, " eating. Total eats: ", eats)
				}
			*/
			time.Sleep(time.Duration(rand.Intn(2)+2) * time.Millisecond) // time it takes for philosopher to eat
			fmt.Println("Philosopher ", i, " thinking")
		}
		if rFork {
			releaseFork((i + 1) % 5)
		}
		if lFork {
			releaseFork(i)
		}
		time.Sleep(time.Duration(rand.Intn(10)+2) * time.Millisecond)
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
