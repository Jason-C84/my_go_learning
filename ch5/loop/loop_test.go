package loop

import "testing"

func TestWhileLoop(t *testing.T){
	n := 0
	for n < 5 {

		n++
		t.Log(n)
	}
}
