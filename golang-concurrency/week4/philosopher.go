package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	times           int
	host            *chan int
	id              int
}

func (p *Philo) eat(wg *sync.WaitGroup) {
	for {
		if p.times == 3 {
			fmt.Printf("finishing eating %d\n", p.id)
			wg.Done()
			return
		}

		<-*p.host
		p.leftCS.Lock()
		p.rightCS.Lock()

		p.times += 1
		fmt.Printf("start eating philosopher %d the %d times\n", p.id, p.times)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		*p.host <- 1
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	host := make(chan int, 2)
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], 0, &host, i + 1}
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg)
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
