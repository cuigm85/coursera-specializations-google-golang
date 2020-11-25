package main

import (
	"fmt"
	"sync"
    "time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS  *chan int
	rightCS *chan int
	times   int
	host    *chan int
	id      int
}

func (p *Philo) eat(wg *sync.WaitGroup) {
	for {
		if p.times == 3 {
			wg.Done()
			return
		}
		<-*p.host

		var handOne int
		var handTwo int
		// random pick up
		select {
		case handOne = <-*p.leftCS:
		case handOne = <-*p.rightCS:
		default:
			handOne = -1
		}
        if handOne == -1 {
			*p.host <- 1
            continue
        }
        if handOne == p.id - 1 { // handOne is left
		    select {
		    case handTwo = <-*p.rightCS:
		    default:
			    handTwo = -1
		    }
            if handTwo == -1 {
                *p.leftCS <- handOne
			    *p.host <- 1
                continue
            }
        } else {
		    select {
		    case handTwo = <-*p.leftCS:
		    default:
			    handTwo = -1
		    }
            if handTwo == -1 {
                *p.rightCS <- handOne
			    *p.host <- 1
                continue
            }
        }


		p.times += 1
		fmt.Printf("philosopher %d start eating %d time(s)\n", p.id, p.times)
        time.Sleep(10 * time.Millisecond)
        if handOne == p.id - 1 { // handOne is left
		    *p.rightCS <-handTwo
		    *p.leftCS <- handOne
        } else {
		    *p.leftCS <-handTwo
		    *p.rightCS <-handOne
        }
		fmt.Printf("philosopher %d finishing eating %d time(s)\n", p.id, p.times)
		*p.host <- 1
	}
}

func main() {
	const philosNo = 5
	host := make(chan int, 2)
	var CSticks [philosNo]chan int
	for i := range CSticks {
		CSticks[i] = make(chan int, 1)
	}

	philos := make([]*Philo, philosNo)
	for i := 0; i < philosNo; i++ {
		philos[i] = &Philo{&CSticks[i], &CSticks[(i+1)%philosNo], 0, &host, i + 1}
	}

	var wg sync.WaitGroup
	wg.Add(philosNo)
	for i := 0; i < philosNo; i++ {
		go philos[i].eat(&wg)
	}

	// host put CSticks on the table
	for i := range CSticks {
		CSticks[i] <- i
	}

	// buffer size of 2 make sure that no more than 2 philosophers to eat concurrently.
	for i := 0; i < 2; i++ {
		host <- 1
	}
	wg.Wait()
	for i := 0; i < 2; i++ {
		<-host
	}
}
