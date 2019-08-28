package ch1

import (
	"testing"
	"errors"
)

var LessThanError = errors.New("n should be less 100")
var LargerThanError = errors.New("n should be larger than 2")

func GetFibonacci(n int) ([]int, error){
	if n > 100 {
		return nil, LessThanError
	}
	if n < 2 {
		return nil, LargerThanError
	}
	fibList := []int{1,1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2] + fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {

	if v, err := GetFibonacci(10); err != nil{
		if err == LessThanError{
			t.Log("it is less")
		}
		t.Error(err)
	}else {
		t.Log(v)
	}

}