package main

/*
There are five philosophers dining together at the same round table.
There are five plates, one in front of each philosopher, and one fork between each plate,
five forks total. The dish they are eating requires them to use both forks,
one on their left side and the other on their right side.
Each philosopher thinks for a random interval and then eats for a while.
To eat, a philosopher must acquire both forks, one on the left side and the other
on the right side of the plate.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func philisopher(inx int, leftFork, rightFork chan bool) {
	for {
		fmt.Printf("Philosopher %v is thinking\n", inx)
		time.Sleep(time.Duration(rand.Intn(1000)))
		// taking fork off the table
		select {
		case <-leftFork:
			select {
			case <-rightFork:
				fmt.Printf("Philisopher %v is eating", inx)
				time.Sleep(time.Duration(rand.Intn(1000)))
				rightFork <- true
			default:
			}
			leftFork <- true
		}
	}
}

func main() {
	var forks [5]chan bool
	// Put all forks on the the table
	for i := range forks {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}
	go philisopher(1, forks[4], forks[0])
	go philisopher(2, forks[0], forks[1])
	go philisopher(3, forks[1], forks[2])
	go philisopher(4, forks[2], forks[3])
	go philisopher(5, forks[3], forks[4])
	select {}
}
