package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// function to add an array of numbers.
func sum(s []int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	//runtime.SetBlockProfileRate(1)
	go foo()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo() {
	for i := 0; i < 100000; i++ {
		go sum([]int{7, 2, 8, -9, 4, 0})
	}
}
