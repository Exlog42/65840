//go:build ignore

package main

import (
	"fmt"
	"sync"
)

// 最佳实践，把共享数据和它的专属锁，打包放进一个结构体里
type SafeCounter struct {
	count int
	mu    sync.Mutex
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()         //先上锁
	defer c.mu.Unlock() //defer解锁，即便发生panic也会解锁
	c.count++
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func routine(wg *sync.WaitGroup, c *SafeCounter) {
	defer wg.Done()
	c.Increment()
}

func main() {
	counter := SafeCounter{}
	// 开启1000个协程增加计数
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go routine(&wg, &counter)
	}
	wg.Wait()
	fmt.Println("最终的计数器结果：", counter.Get())
}
