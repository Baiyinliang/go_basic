package operator_test

import "testing"

const (
	Readable = 1 << iota
	Wirtable
	Executable
)


/**
用==比较数组
1.相同纬度切含有相同个属元素的数组才可以比较
2.每个元素都相同的才相等
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 3}
	//c := [...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

/**
位运算符
与其他编程语言的主要差异
&^按位置零
只要右边的位操作为1，那么结果就是1，如果右边的位操作为0，那么左边是什么就是什么
 */
func TestBitClear(t *testing.T) {
	a := 7 //0111

	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Wirtable == Wirtable, a&Executable == Executable)
}
