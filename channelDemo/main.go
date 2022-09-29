package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type UserTaskHandleData struct {
	TaskEventType int
	Msg           []byte
}

func main() {
	strs := "{\"TaskEventType\":4,\"Msg\":\"eyJQbGF0Zm9ybUlkIjoyLCJBcHBsaWNhdGlvbklkIjoxMDAyLCJDaGFubmVsSWQiOjAsIlVzZXJJZCI6OTg4OTY3NiwiRGVzYyI6IiIsIlBhcmVudElkIjowfQ==\"}\n"

	m := UserTaskHandleData{}
	err := json.Unmarshal([]byte(strs), &m)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m)
	m2 := make(map[string]interface{})
	err = json.Unmarshal(m.Msg, &m2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", m2)
	return
	cond := sync.NewCond(&sync.Mutex{})
	intC := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			go func(i int) {
				intC <- i
			}(i)
		}
		close(intC)
	}()
	go func() {
		for {
			n, closed := <-intC
			if closed {
				fmt.Println(n)
				cond.L.Lock()
				cond.Broadcast()
				cond.L.Unlock()
				break
			} else {
				fmt.Println(n)
			}
		}
	}()
	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()
}
