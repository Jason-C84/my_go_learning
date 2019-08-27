package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a,b,c)
}

func TestPoint(t *testing.T) {
	var a int = 1
	var aptr = &a
	aptr = aptr+1
	t.Logf("%T %T", a, aptr)


}
