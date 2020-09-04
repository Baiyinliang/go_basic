package groutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}

func TestGo(t *testing.T) {
	fmt.Println("所有协程已开始执行")
	go Add(1, 1)
	fmt.Println("所有协程已执行完毕")
}

func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}

//chanel 双向管道通信
//goroutine 运行再相同的地址空间，访问共享内存必须做好同步，Go 提供了一种机制，用于协程之间进行数据通信『channels』这个类似于 Unix shell 中的双向管道。

//由 make 创建，创建时需要指定类型
//可以接收、发送指定类型数据
//使用 <- 操作符进行接收、发送数据，channels 接收和发送数据都是阻塞的

func TestChannel(t *testing.T) {
	a := []int{7, 2, 8, -9, 4, 0}

	ch := make(chan int)
	go multi(a[:len(a)/2], ch) // 开启一个协程 并用 ch 作为通信通道
	go sum(a[len(a)/2:], ch)

	fmt.Println(a[:len(a)/2])
	fmt.Println(a[len(a)/2:])

	fmt.Println("ch 中没有数据我将阻塞起来等数据")
	x, y := <-ch, <-ch // 接收 ch 通道中的值
	fmt.Println(x, y, x+y)
}

// Sleep 需要 time 包
func sum(a []int, ch chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println("我将 10s 后计算完成再将数据 push 到通道 ch")
	time.Sleep(time.Second * 10)
	ch <- sum
}

func multi(a []int, ch chan int) {
	sum := 1
	for _, v := range a {
		sum *= v
	}
	fmt.Println("我将 5s 后计算完成再将数据 push 到通道 ch")
	time.Sleep(time.Second * 5)
	ch <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestABC(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
