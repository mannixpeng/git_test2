package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var Data = make(map[string]int)
var flag uint32
func main() {
	go test1()
	go test2()
	select {

	}
}

func test1() {
	for {
		v := atomic.CompareAndSwapUint32(&flag, 0, 1)
		if v {
			for k, val := range Data {
				fmt.Println("key: ", k, " val: ", val)
			}
			atomic.StoreUint32(&flag, 0)
			time.Sleep(10 * time.Second)
		}
		fmt.Println(flag)
		time.Sleep(100* time.Microsecond)
	}

}

func test2() {
	var count int
	for {
		v := atomic.CompareAndSwapUint32(&flag, 0, 1)
		if v {
			for i := 0; i < 100; i++ {
				Data[fmt.Sprintf("%d次改变", count)] += i
				count++
			}
			atomic.StoreUint32(&flag, 0)
			fmt.Println("test2", flag)
			time.Sleep(4 * time.Second)
			continue
		}
		time.Sleep(10 * time.Microsecond)
	}

}
