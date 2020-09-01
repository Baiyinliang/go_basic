package remote_package

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

/**
go中的默认的map不是线程安全的
*/

func TestPackage(t *testing.T) {
	var num1 float64 = 3.1
	var num2 int = 2
	d1 := decimal.NewFromFloat(num1).Mul(decimal.NewFromFloat(float64(num2)))
	fmt.Println(d1)
}

/**
1.通过go get获取远程依赖
2.go get -u 强制从网络更新远程依赖
3.注意代码在github上的组织形式，以适应go get
4.直接以代码路径开始，不要有src
*/
