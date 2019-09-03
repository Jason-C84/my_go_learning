package until_anyone_reply

import (
	"fmt"
	"time"
	"testing"
	"runtime"
)

func runTask(id int) string{
	time.Sleep(time.Millisecond*10)
	return fmt.Sprintf("The result is from task: %d \n", id)
}

func firstResponse() string {
	numOnRunner := 10
	ch := make(chan string, numOnRunner)
	//ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(id int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	return <- ch
}

func AllResponse() string {
	numOnRunner := 10
	ch := make(chan string, numOnRunner)
	//ch := make(chan string)
	for i := 0; i < 10; i++ {
		go func(id int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	finalResp := ""
	for i := 0; i < numOnRunner; i++ {
		finalResp += <-ch
	}

	return finalResp
}

func TestFirstResponse(t *testing.T){
	t.Logf("Before Response: %d\n", runtime.NumGoroutine())
	//t.Log(firstResponse())
	t.Log(AllResponse())
	time.Sleep(time.Second*1)
	t.Logf("After Response: %d\n", runtime.NumGoroutine())

}

