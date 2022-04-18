package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{})
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	go func() {
		t := time.NewTicker(1 * time.Second)
		t1 := time.NewTicker(2 * time.Second)
		t2 := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-t.C:
				fmt.Println("t: ", time.Now().Format("2006-01-02 15:04:05"))
				c <- struct{}{}
			case <-t1.C:
				fmt.Println("t1: ", time.Now().Format("2006-01-02 15:04:05"))
				c1 <- struct{}{}
			case <-t2.C:
				fmt.Println("t2: ", time.Now().Format("2006-01-02 15:04:05"))
				c2 <- struct{}{}
			}
		}
	}()
	for {
		select {
		case <-c:
			go func() {
				fmt.Println("c: ", time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(10 * time.Second)
			}()

		case <-c1:
			 go func() {
				 fmt.Println("c1: ", time.Now().Format("2006-01-02 15:04:05"))
				 time.Sleep(20 * time.Second)
			 }()

		case <-c2:
			go func() {
				fmt.Println("c2: ", time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(30 * time.Second)
			}()

		}
	}
}
