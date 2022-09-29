package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var flag bool
	go func() {
		time.Sleep(time.Second * 15)
		cond.L.Lock()
		flag = true
		cond.Broadcast()
		cond.L.Unlock()
	}()

	fmt.Println("waiting")
	go func() {
		cond.L.Lock()
		for !flag {
			fmt.Println("waiting flag 1")
			cond.Wait()
		}
		fmt.Println("wait1")

		cond.L.Unlock()
	}()
	go func() {
		cond.L.Lock()
		for !flag {
			fmt.Println("waiting flag 2")
			cond.Wait()
		}
		fmt.Println("wait2")
		cond.L.Unlock()
	}()
	cond.L.Lock()
	fmt.Println("waiting flag 3")
	cond.Wait()
	cond.L.Unlock()
	fmt.Println("wait")
	fmt.Println("done")
	time.Sleep(1 * time.Second)
}
