package main

import (
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

//func fib(n int64) int64 {
//	//helloworld
//	if n < 0 {
//		return 0
//	}
//	if n <= 2 {
//		return n
//	}
//	return fib(n-1) + fib(n-2)
//}
//
//func main() {
//	for i := 0; i < 10; i++ {
//		fmt.Println(i)
//		result := fib(int64(i))
//		fmt.Println(result)
//	}
//	fmt.Println("hello world")
//	time.Sleep(100 * time.Hour)
//	select {}
//	//var wg sync.WaitGroup
//	//for i := 0; i < 10000000; i++ {
//	//	wg.Add(1)
//	//	o := i
//	//	go func() {
//	//		defer wg.Done()
//	//		fmt.Println("tmpO: ", o)
//	//		rand.Seed(time.Now().UnixNano())
//	//		time.Sleep(500 * time.Second)
//	//		//time.Sleep(time.Duration(rand.Intn(50)) * time.Second)
//	//	}()
//	//	numGoroutine := runtime.NumGoroutine()
//	//	fmt.Printf("numGoroutine: %d\n", numGoroutine)
//	//}
//	//go func() {
//	//	for i := 0; i < 1000; i++ {
//	//		numGoroutine := runtime.NumGoroutine()
//	//		fmt.Printf("numGoroutine: %d\n", numGoroutine)
//	//		time.Sleep(100 * time.Millisecond)
//	//	}
//	//}()
//	//wg.Wait()
//}

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
	time.Sleep(1 * time.Second)
}

func demoFunc() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	num, err := strconv.ParseInt("1", 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
	return
	fmt.Println(len("99476850"))
	return
	fmt.Println(10 << 20)
	return
	var Ancestors string
	ancestors := strings.Split(Ancestors, ",")

	Ancestors = strings.Join(ancestors, ",")
	fmt.Println(Ancestors)

	return
	var f float64 = 1.11111111
	fmt.Sprintf("%.2f", f)

	return

	unix := time.Now().Unix()
	fmt.Println(unix)

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	in := time.Unix(unix, 0).In(location)
	fmt.Println(in)
	return
	agents := make([]int64, 0)
	agents = append(agents, 1, 3, 2, 5)
	sort.Slice(agents, func(i, j int) bool {
		return agents[i] > agents[j]
	})
	fmt.Println(agents)
	return
	defer ants.Release()
	runTimes := 10000

	// Use the common pool.
	var wg sync.WaitGroup
	//syncCalculateSum := func() {
	//	demoFunc()
	//	wg.Done()
	//}
	startTime := time.Now().Unix()
	for i := 0; i < runTimes; i++ {
		i := i
		wg.Add(1)
		_ = ants.Submit(func() {
			defer wg.Done()
			myFunc(int32(i))
		})
	}
	fmt.Println("1 goroutines: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("costTime: %d\n", time.Now().Unix()-startTime)
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	startTime = time.Now().Unix()
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		defer wg.Done()
		myFunc(i)
	}, ants.WithNonblocking(false), ants.WithExpiryDuration(1*time.Hour))
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		if err := p.Invoke(int32(i)); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("2 goroutines: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Printf("costTime: %d\n", time.Now().Unix()-startTime)
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)
}
