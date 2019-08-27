package extention

import (
	"fmt"
	"testing"
)

type Pet struct{

}

func (p *Pet) Speak(){
	fmt.Println("...")
}

func (p *Pet) SpeakTo(name string){
	fmt.Println(" ", name)
}

type Dog struct {
	p *Pet
}

func (d *Dog) speak(){
	fmt.Println("wang!")
}
func (d *Dog) speakTo(name string) {
	d.speak()
	fmt.Println(" ", name)
}

func TestClient2(t *testing.T){
	p := new(Pet)
	p.Speak()
	p.SpeakTo("www")
	d := new(Dog)
	d.speak()
	d.speakTo("hhhh")
	d.p.Speak()
	d.p.SpeakTo("ppp")
}

type Programmer interface {
	writeHandle() string
}

type GoProgrammer struct {

}

type JavaProgrammer struct {

}

func (p *GoProgrammer) writeHandle() string{
	return "fmt.printf(\"Hello World!\")"
}

func (p *JavaProgrammer) writeHandle() string{
	return "fmt.printf(\"Hello World\")"
}

func WriteProgram(p Programmer){
	fmt.Printf("%T, %v\n", p, p.writeHandle())
}
func TestClient3(t *testing.T){
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)

	WriteProgram(goProg)
	WriteProgram(javaProg)
}