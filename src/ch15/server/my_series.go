package server

import "fmt"

func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

/**
init方法
1.在main被执行钱，所有依赖的package的init方法都会被执行
2.不同包的init函数按照包导入的依赖关系决定执行顺序，golang会自己帮我们处理依赖顺序
3.每个包可以有多个init函数
4.包的每个源文件也可以有多个init函数，这点比较特殊
*/

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
