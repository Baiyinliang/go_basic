package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr[0], arr[1])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	//idx为key，e为val
	for idx, e := range arr {
		t.Log(idx, e)
	}

	//当不需要关系key的时候，可以用_代替
	for _, e := range arr {
		t.Log(e)
	}
}

//多维数组
func TestMultiWeftArray(t *testing.T) {
	arr := [2][5]int{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for id, e := range arr {
		for idx, a := range e {
			t.Log(id, idx, a)
		}
	}
}

/**
数组截取
arr[开始索引(包含):结束索引(不包含)]
*/
func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	arr_sec := arr[3:]
	t.Log(arr_sec)
}


