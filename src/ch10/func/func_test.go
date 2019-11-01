package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
函数是一等公民
与其他主要编程语言的差异
1.可以有多个返回值
2.所有参数都是值传递：slice，map，channel会有传引用的错觉
slice本身对应的是一个数组，数据结构中包含了对应的数据结构的指针
3.函数可以作为变量的值
4.函数可以作为参数的返回值
*/
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/**
实现函数方法执行的时间差计算
*/
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func SlowFun(ot int) int {
	time.Sleep(time.Second * 1)
	return ot
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(SlowFun)
	t.Log(tsSF(20))
}
