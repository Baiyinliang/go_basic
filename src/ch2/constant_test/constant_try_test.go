package constant_test

import "testing"

//连续常量的命名方式
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Wirtable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)

}

func TestConstantTry1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Wirtable == Wirtable, a&Executable == Executable)
}
