package string

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

/**
分割与连接字符串
*/
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))

	t.Log(reflect.TypeOf(parts)) //使用reflect.TypeOf()获取数据类型
}

/**
字符串与int之间的转换
*/
func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str " + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
