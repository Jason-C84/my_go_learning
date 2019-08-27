package constant_test

import "testing"

const(
	Mon = 1 + iota
	Thu
	Web
)
const (
	Readable = 1 << iota
	Writeable
	Execable
)

func TestConstantTry(t *testing.T){
	t.Log(Mon, Thu, Web)
}

func TestConstantTry1(t *testing.T){
	a := 7
	t.Log(a&Readable, a&Writeable, a&Execable)
	b := 1
	t.Log(b&Readable, b&Writeable, b&Execable)
}
