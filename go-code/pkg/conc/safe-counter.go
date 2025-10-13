package conc

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	Mu    sync.Mutex
	WG    sync.WaitGroup
	Val   int
	Print bool
}

func NewSafeCounter() *SafeCounter {
	sc := SafeCounter{}
	sc.Mu = sync.Mutex{}
	sc.Val = 0
	sc.Print = true
	return &sc
}

func (sc *SafeCounter) increment() {
	sc.Mu.Lock()
	sc.Val++
	if sc.Print {
		fmt.Printf("\rCounter value = %d%%", sc.Val)
		time.Sleep(500 * time.Millisecond)

	}
	sc.Mu.Unlock()
}
func (sc *SafeCounter) decrement() {
	sc.Mu.Lock()
	sc.Val--
	if sc.Print {
		fmt.Printf("\rCounter value = %d%%", sc.Val)
		time.Sleep(500 * time.Millisecond)
	}
	sc.Mu.Unlock()
}

func (sc *SafeCounter) Increment(iter int) {
	for i := 0; i < iter; i++ {
		sc.increment()
	}
	sc.WG.Done()
}
func (sc *SafeCounter) Decrement(iter int) {
	for i := 0; i < iter; i++ {
		sc.decrement()
	}
	sc.WG.Done()
}

func CountUpAndDownTest(iter int) {

	sc := NewSafeCounter()

	sc.WG.Add(2)
	go sc.Increment(iter)
	go sc.Decrement(iter)

	sc.WG.Wait()
	fmt.Println("\nFinal counter value =", sc.Val)

}
