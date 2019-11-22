package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "白小白", 21}
	e1 := Employee{Name: "Evan", Age: 12}
	e2 := new(Employee) //指针类型
	e2.Id = "2"
	e2.Name = "Bai"
	e2.Age = 23
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//func (e *Employee) String() string {
//	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "白小白", 21}
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
