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
	e2 := new(Employee) //注意这里反悔的引用/指针。相当于 e := &Employee{}
	e2.Id = "2"         //与其他主要编程语言的差异：通过实例的指针访问成员不需要使用->
	e2.Name = "Bai"
	e2.Age = 23
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

/**
行为（方法）定义，与其他主要编程语言的差异
*/
//第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制
//使用此方法会有数据复制，会有更大的空间上复制的开销
//func (e Employee) String() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) //比较两者的区别
//	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
//使用此方法没有新的数据产生，所有的数据都是在同一个地址的
func (e *Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name)) //比较两者的区别
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "白小白", 21}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
