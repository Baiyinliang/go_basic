package type_test

import (
	"testing"
)

type MyInt int64

/**
类型转化
1.GO语言不支持隐式类型转换
2.别名和原有类型也不能进行隐式类型转换
*/

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

/**
指针类型
与其他主要编程语言的差异
1.不支持指针运算
2.string是值类型，其默认的初始值为空字符串，而不是nil
*/
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //a对应的指针
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	//判断空字符串的方式
	if s == "" {

	}
}
