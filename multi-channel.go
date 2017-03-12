package main

import (
    "fmt"
	"time"
)

func put(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			time.Sleep(500 * time.Millisecond)
			out <- n
		}
		close(out)
	}()
	return out
}

func putIt() <-chan int {
    out := make(chan int)
    go func() {
        for i:=0;i<10000; i++ {
            out <- i
        }
        close(out)
    }()
    return out
}

func get(in <-chan int, num int) {
	for n := range in {
		fmt.Printf("Channel %v: I got %v\n", num, n)
	}
}

func main() {
	// out := put(1,2,3,4,5,6,7,8,9,10)
	out := putIt()
	go get(out, 1)
	get(out, 2)
}
