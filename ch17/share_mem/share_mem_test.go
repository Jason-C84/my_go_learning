package share_mem

import (
	"testing"
	"time"
	"sync"
	"fmt"
)

func TestCounter(t *testing.T){
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter ++
		}()
	}

	time.Sleep(1 + time.Second)
	t.Log(counter)
}

func TestCounterSafe(t *testing.T){
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(){
			defer func() {
				mut.Unlock()
			}()

			mut.Lock()

			counter++
		}()

	}

	time.Sleep(1 + time.Second)
	t.Log(counter)
}

func TestCounterWaitGroup(t *testing.T){
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(){
			defer func(){
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log(counter)
}

func service(){
	fmt.Println("do something")
	return
}


