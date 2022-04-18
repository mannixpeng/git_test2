package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"sync"
	"time"
)

//func main() {
//	// 务必先进行初始化
//	err := sentinel.InitDefault()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 配置一条限流规则
//	_, err = flow.LoadRules([]*flow.Rule{
//		{
//			Resource:                "some-test",
//			TokenCalculateStrategy:  flow.Direct,
//			ControlBehavior:         flow.Reject,
//			Threshold:               1,
//			StatIntervalInMs:        1000,
//		},
//	})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	ch := make(chan struct{})
//	for i := 0; i < 10; i++ {
//		go func() {
//			for {
//				// 埋点逻辑，埋点资源名为 some-test
//				e, b := sentinel.Entry("some-test")
//				if b != nil {
//					// 请求被拒绝，在此处进行处理
//					fmt.Println("请求被拒绝")
//					time.Sleep(time.Duration(rand.Uint64() % 10) * time.Millisecond)
//				} else {
//					// 请求允许通过，此处编写业务逻辑
//					fmt.Println(util.CurrentTimeMillis(), "Passed")
//					time.Sleep(time.Duration(rand.Uint64() % 10) * time.Millisecond)
//					fmt.Println("请求通过")
//					// 务必保证业务结束后调用 Exit
//					e.Exit()
//				}
//
//			}
//		}()
//	}
//	<-ch
//}

type RateLimiter struct {
	resources map[string]ratelimit.Limiter
	mu        *sync.RWMutex
}

// NewRateLimiter .
func NewRateLimiter() *RateLimiter {
	i := &RateLimiter{
		resources: make(map[string]ratelimit.Limiter),
		mu:        &sync.RWMutex{},
	}

	return i
}

// addResource creates a new rate limiter and adds it to the resources map,
// using the resource as the key
func (r *RateLimiter) addResource(resource string, per time.Duration, rate int) ratelimit.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter := ratelimit.New(rate, ratelimit.Per(per))

	r.resources[resource] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided resource if it exists.
// Otherwise calls addResource to add resource to the map
func (r *RateLimiter) GetLimiter(resource string, per time.Duration, rate int) ratelimit.Limiter {
	r.mu.Lock()
	defer r.mu.Unlock()
	limiter, exists := r.resources[resource]

	if !exists {
		return r.addResource(resource, per, rate)
	}
	return limiter
}

func main() {
	limiter := NewRateLimiter()
	tests := []struct {
		resource string
		period   time.Duration
		rate     int
	}{
		{
			resource: "127.0.0.1",
			period: 1 * time.Second,
			rate: 1,
		},
		{
			resource: "127.0.0.2",
			period: 1 * time.Second,
			rate: 1,
		},
		{
			resource: "127.0.0.3",
			period: 1 * time.Second,
			rate: 1,
		},
	}
	prev := time.Now()
	for i, t := range tests {
		getLimiter := limiter.GetLimiter(t.resource, t.period, t.rate)
		go func(i int) {
			for j := 0; j < 100; j++ {
				now := getLimiter.Take()
				fmt.Println(i, j, now.Sub(prev))
				prev = now
			}
		}(i)
	}
	time.Sleep(60 * time.Second)
}
