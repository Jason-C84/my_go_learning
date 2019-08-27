package condition

import "testing"

func TestIfMultSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}
}

func TestSwitchMult(t *testing.T){
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3 :
			t.Log("Odd")
		default:
			t.Log("number is not 0-3")
		}

	}

}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch  {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1 :
			t.Log("Odd")
		default:
			t.Log("unknown")
		}

	}
}