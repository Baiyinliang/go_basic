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

//使用此方法会有数据复制，会有更大的空间上复制的开销
func (e Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

//使用此方法没有新的数据产生，所有的数据都是在同一个地址的
//func (e *Employee) String() string {
//	fmt.Printf("Address is %x",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
//}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "白小白", 21}
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
