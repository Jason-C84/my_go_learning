package _interface

import "testing"

type programmer interface {
	WriteHandler() string
}

type GoProgrammer struct{

}

func (gp *GoProgrammer) WriteHandler() string{
	return "fmt.println(\"Hello World!\")"
}

func TestClient(t *testing.T){
	p := new(GoProgrammer)
	t.Log(p.WriteHandler())
}