package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go run(&wg)
	}
	wg.Wait()
}

func run(inWg *sync.WaitGroup) {
	defer inWg.Done()
	var wg sync.WaitGroup
	closeCh := make(chan struct{})
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Printf("dial failed: %s\n", err)
		return
	}
	inCh := make(chan string)
	go func() {
		var data string
		data = data
		for {
			//fmt.Println("请数据需要发送的数据：")
			//_, err := fmt.Scanf("%s", &data)
			//if err != nil {
			//	panic(err)
			//}
			rand.Seed(time.Now().UnixNano())
			intn := rand.Intn(1000000)
			data := fmt.Sprintf("data: %d", intn)
			inCh <- data
			if data == "quit" {
				closeCh <- struct{}{}
				break
			}
		}
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err == io.EOF {
				fmt.Println(err)
				break
			}
			if err != nil {
				fmt.Printf("conn read failed: %s", err)
				break
			}

			fmt.Println(string(buffer[:n]))
		}
	}()
	for {
		select {
		case d := <-inCh:
			_, err = conn.Write([]byte(d))
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("write conn failed: %s", err)
				return
			}
		case <-closeCh:
			goto endLoop
		}
	}
endLoop:
	fmt.Println("client close")
	wg.Wait()
}
