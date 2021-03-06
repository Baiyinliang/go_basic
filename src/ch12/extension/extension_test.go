package extension

import (
	"fmt"
	"testing"
)

//扩展与复用

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Chao")
}
