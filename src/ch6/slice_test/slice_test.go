package slice_test

import "testing"

/**
切片slice
切片声明方式[]type,len,cap
切片主要有长度和容量两部分组成
长度是指当前数组已经拥有的长度
容量是指切片包含的最大长度
*/
func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) //声明一个类型为整型，长度为3，容量为5的切片
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

/**
切片中容量的增长会随着长度的增长而增长，当容量无法存储现有的长度时，
容量会以切片长度*2的形式增长
*/
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s := append(s, i)
		t.Log(len(s), cap(s))
	}
}

/**
切片的存储空间是共享的，他们共享了同一个后端数组，一旦修改了后面的数据，原有的数据也会发生变化。
 */
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}


/**
数组VS切片
1.容量是否可伸缩
答：数组的容量是不可伸缩的。
2.是否可以进行比较
答：切片不可进行比较，数组可以
 */
func TestSliceComparing(t *testing.T)  {
	//a:=[]int{1,2,3,4}
	//b:=[]int{1,2,3,4}
	//
	//if a== b{
	//	t.Log("equal")
	//}
}