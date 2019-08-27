package _map

import (
	"testing"
)

func TestMapWithValue(t *testing.T){
	m1 := map[int]func(op int)int{}
	m1[1] = func (op int)int {return op}
	m1[2] = func (op int)int {return op*op}
	m1[3] = func (op int)int {return op*op*op}

	t.Log(m1[1](3), m1[2](3),m1[3](3))

}