package main

var philChans [5]chan bool

var forkChans [5]chan bool

func philosopher(i int) {
	rFork := requestFork(i)
	lFork := requestFork((i - 1) % 5)

	if rFork != lFork {
		if lFork {
			rFork = requestFork(i)
		} else {
			lFork = requestFork((i - 1) % 5)
		}
	}
}

func requestFork(i int) bool {
	forkChans[i] <- true
	return <-forkChans[i]
}

func fork(i int) {
	var inUse = false

	for {
		<-forkChans[i] //waits for signal
		if inUse {
			forkChans[i] <- false
		} else {
			inUse = true
			forkChans[i] <- true
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
