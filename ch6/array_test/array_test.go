package array_test

import "testing"

func TestArrayInit(t *testing.T){
	var arr []int;
	arr1 := [3]int{1,2,3}
	arr2 := [...]int{1,2,3,4}
	t.Log(arr)
	t.Log(arr1)
	t.Log(arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,3,4,5}
	for _,val := range arr3{
		t.Log(val)
	}
}

func TestArraySection(t *testing.T){
	arr := [...]int{1,2,3,4,5,6}
	arr_sec := arr[:4]
	arr_sec1 := arr[4:]
	arr_sec2 := arr[:]
	arr_sec3 := arr[1:3]
	t.Log(arr_sec)
	t.Log(arr_sec1)
	t.Log(arr_sec2)
	t.Log(arr_sec3)

}