package empty_interface

import (
	"fmt"
	"testing"
)

/**
空接口与断言
1.空接口可以表示任何类型
2.通过断言来将空接口转换为制定类型
v,ok := p.(int) //ok=true 时为转换成功
*/

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if s, ok := p.(string); ok {
	//	fmt.Println("string", s)
	//	return
	//}
	//fmt.Println("UnKnow Type")
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("UnKnow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

/**
Go 接口最佳实践
倾向于使用小的接口定义，很多接口只包含一个方法
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

-------------------------------------------

较大的接口定义，可以由多个小接口定义组合而成

type ReadWriter interface {
	Reader
	Writer
}

-----------------------------------

只依赖于必要功能的最小接口
func StoreData(reader Reader) error {
	///业务逻辑
}

*/
