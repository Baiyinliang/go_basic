package loop_test

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	//等同于while(n < 5)
	for n < 5 {
		t.Log(n)
		n++
	}
}

