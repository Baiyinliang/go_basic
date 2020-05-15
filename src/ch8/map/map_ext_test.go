package map_ext

import (
	"testing"
)

/**
map与工厂模式
*/
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

/**
实现set
Go的内置集合没有Set实现，可以map[type]bool
1.元素的唯一性
2.基本操作
1) 添加元素
2) 判断元素是否存在
3) 删除元素
4) 元素个数
*/
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet)) //获取map长度
	delete(mySet, 1)  //删除指定map的键
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
