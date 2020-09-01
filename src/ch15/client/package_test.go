package client

import (
	//"ch15/server"
	"testing"
)

/**
package
1.基本服用模块单元
 以首字母大写来表明可被包外代码访问，相当于PHP中的public
2.代码的package可以和所在的目录不一致
3.同一目录里的Go代码的package要保持一致
*/

func TestPackage(t *testing.T) {
	//t.Log(server.GetFibonacci(5))
}

/**
go中的默认的map不是线程安全的
*/
