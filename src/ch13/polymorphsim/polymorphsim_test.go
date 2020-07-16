package polymorphsim

import (
	"fmt"
	"testing"
)

//接口与多态

type Code string
type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld()) //%T输出实例或者传入的参数类型
}

func TestPolymorphism(t *testing.T) {
	//Programmer是一个interface，对应的类型只能是指针类型
	goProg := &GoProgrammer{}       //
	javaProg := new(JavaProgrammer) //指针类型的实例引用必须用new或者&取址符来处理
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}
