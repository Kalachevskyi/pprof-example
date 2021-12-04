package block

import (
	"runtime"
	"testing"
)

func BenchmarkFoo(b *testing.B) {
	runtime.SetBlockProfileRate(1)
	d := make(chan struct{})
	for n := 0; n < b.N; n++ {
		go sum([]int{7, 2, 8, -9, 4, 0}, d)
	}
	done(b.N, d)
}

//func TestO(t *testing.T) {
//	runtime.SetBlockProfileRate(1)
//	d := make(chan struct{})
//	for n := 0; n < 10000; n++ {
//		go sum([]int{7, 2, 8, -9, 4, 0}, d)
//	}
//	done(10000, d)
//}
