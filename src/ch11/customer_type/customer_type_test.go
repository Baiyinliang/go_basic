package customer_type

import (
	"fmt"
	"testing"
	"time"
)

/**
自定义类型
可以定义数据为指定的类型
*/

type IntCov func(op int) int

/**
实现函数方法执行的时间差计算
*/
func timeSpent(inner IntCov) IntCov {
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
	tsSF := timeSpent(SlowFun)
	t.Log(tsSF(20))
}
