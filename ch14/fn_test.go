package ch1

import "testing"
import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

func returnMultiValues ()(int, int){
	return rand.Intn(10), rand.Intn(5)
}

func timeSpent(inner func(op int) int) func(op int) int{
	return func(n int) int{
		start := time.Now()
		ret := inner(n)
		fmt.Println("sleep 10")
		fmt.Println("time spent :", time.Since(start).Seconds())
		return ret
	}
}

func TestFn(t *testing.T){
	a, b := returnMultiValues()
	t.Log(a, b)
	t.Log("=======================")
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op+20
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T){
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(11,22,33,44))
}

func Clear(){
	fmt.Println("clear clear")
}

func TestDefer(t *testing.T){
	defer Clear()
	fmt.Println("start")
	//panic("err")
}

type employee struct{
	id int
	name string
	age string
}

func (e employee) String() string{
	return strconv.Itoa(e.id)+"_"+e.name+ "_"+e.age
}

func TestCreateStruct(t *testing.T){
	e := employee{12,"Bob","35"}
	t.Log(e)
	t.Log(e.String())
	e1 := employee{id: 13, name: "John", age: "54"}
	t.Log(e1)
	t.Log(e1.String())
	e2 := new(employee)
	e2.id = 14
	e2.name = "Lucy"
	e2.age = "28"
	t.Log(e2)
	t.Log(e2.String())
}