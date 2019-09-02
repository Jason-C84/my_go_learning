package ch20

import (
	"testing"
	"time"
	"fmt"
	"context"
)

func isCancelled(cancelChan chan struct{}) bool{
	select{
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}){
	cancelChan<- struct{}{}
	fmt.Println("cancel_1")
}

func cancel_2(cancelChan chan struct{}){
	fmt.Println("cancel_2")
	close(cancelChan)
}

func TestTaskCancel(t *testing.T){
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond*5)
			}
			fmt.Println(i, "Cancelled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	//cancel_1(cancelChan)
	time.Sleep(time.Second + 1)
}

func isCancelled1(ctx context.Context) bool{
	select {
		case <- ctx.Done():
			return true
		default:
			return false
	}

}

func TestTaskCancelContext(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled1(ctx) {
					break
				}
				time.Sleep(time.Millisecond*5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	//cancel_1(cancelChan)
	time.Sleep(time.Second + 1)
}