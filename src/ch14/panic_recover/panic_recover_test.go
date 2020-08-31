package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

/**
panic
1.panic用于不可恢复的错误
2.panic推出前会执行defer指定的内容
os.Exit
1.os.Exit退出时不会调用defer指定的函数
2.os.Exit退出时不输出当前调用栈信息
*/

func TestPanicVsExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally")
	//}()
	//可以通过recover处理异常数据，但是recover很危险
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong"))
	//os.Exit(-1)
}

/**
当心! recover成为恶魔
1.可能是系统资源消耗完毕，导致异常无法抛出，使用recover容易形成僵尸服务进程，导致health check失效
2."Let it Crash!"往往是我们恢复不确定性错误的最好方法
*/
