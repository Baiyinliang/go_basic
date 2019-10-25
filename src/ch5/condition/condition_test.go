package condition_test

import "testing"

/**
if条件
1.condition表达式结果必须为布尔值
2.支持变量赋值
*/
func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

/**
switch条件
与其他主要编程语言的差异
1.条件表达式不限制为常量或整数；
2.单个case中，可以出现多个结果选项，使用逗号隔开；
3.与C语言等规则相反，Go语言不需要用break来明确退出一个case；
4.可以不设定switch之后的条件表达式，在此种条件下，整个switch结构与多个if...else...的逻辑作用等同
*/
func TestSwitchMultiSec(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
