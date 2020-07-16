package error

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/**
Go的错误机制与其他主要编程语言的差异

1.没有异常机制

2.error类型实现了error接口
type error interface{
	Error() string
}

3.可以通过errors.New来快速创建错误实例
errors.New("n must be in the range [0,1]")
*/

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not Larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	GetFibonacci2("地方")
	//if v, err := GetFibonacci(11); err != nil {
	//	if err == LessThanTwoError {
	//		fmt.Println("It is less 2")
	//	}
	//	t.Error(err)
	//} else {
	//	t.Log(v)
	//}
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)

}
