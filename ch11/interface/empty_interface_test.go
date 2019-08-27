package _interface

import (
	"fmt"
	"testing"
)

func DoStringTest(p interface{}){
	/*
	if i, ok := p.(int); ok {
		fmt.Println("Integer ", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String ", s)
		return
	}
	fmt.Println("unkown")
*/
	switch  v :=p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string ", v)
	default:
		fmt.Println("unknown")
	}
}

func TestDoString(t *testing.T){
	DoStringTest(10)
	DoStringTest("1234")
	DoStringTest('d')
}
