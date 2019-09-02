package ch18

import (
	"time"
	"fmt"
	"testing"
)

func service() string{
	time.Sleep(time.Millisecond*50)
	return "Done"
}

func AsynService() chan string{
	retCh := make(chan string, 1)

	go func(){
		ret := service()
		fmt.Println("return result:")

		retCh <- ret
		fmt.Println("service exited")
	}()

	return retCh
}

func TestSelect(t *testing.T){
	select{
	case  ret :=  <-AsynService():
		t.Log(ret)
	case <-time.After(time.Millisecond*10):
		t.Error("time out")

	}
}