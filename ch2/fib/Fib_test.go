package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T){
	var a int = 1
	var b int = 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		temp := a
		a = b
		b =  temp + b
	}
	fmt.Println()
}

func TestExchange(t *testing.T){
	a := 1
	b := 2
	//tmp := a
	//a = b
	//b = tmp
	a, b = b, a
	t.Log(a, b)
}