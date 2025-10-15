package conc

import (
	"fmt"
	"sync"
)

var counter int

func RaceConditionRun() {

	fmt.Println("---------------------")
	fmt.Println("Counting without mutex")
	counter = 0
	var wg sync.WaitGroup

	increment := func() {
		for i := 0; i < 1000000; i++ {
			counter++
		}
		wg.Done()
	}

	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}

var mu sync.Mutex

func AvoidRaceConditionRun() {
	fmt.Println("---------------------")
	fmt.Println("Counting with a mutex")
	counter = 0

	var wg sync.WaitGroup

	increment := func() {
		for i := 0; i < 1000000; i++ {
			//critical section
			mu.Lock()
			counter++
			mu.Unlock()
		}
		wg.Done()
	}

	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
