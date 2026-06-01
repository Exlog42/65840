//go:build ignore

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(cxt context.Context, id int, jobs <-chan int, res chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-cxt.Done():
			fmt.Printf("worker(%d)收到撤退信号(%s)\n", id, cxt.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				// 通道被关闭，此时管道会一直被读取到0值
				fmt.Println("工作完毕")
				return
			}
			fmt.Printf("worker(%d)开始处理任务(%d)\n", id, job)
			time.Sleep(3 * time.Second)

			// 嵌套select，在交付任务之前再次检查任务是否被取消
			select {
			case <-cxt.Done():
				fmt.Printf("worker(%d)收到撤退信号(%s)\n", id, cxt.Err())
				return
			case res <- fmt.Sprintf("worker(%d)完成了工作", id):
			}

		}
	}
}

func main() {
	// 创建一个带有3s超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // 虽然cancel会在3s之后自动调用，但是养成好习惯

	const numJobs = 10
	const numWorkers = 3
	// 创建分配工作和收集结果的管道
	jobs := make(chan int, numJobs)
	res := make(chan string, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, res, &wg)
	}

	// 把nunJobs个任务塞进工作channel
	for i := range numJobs {
		jobs <- (i + 1)
	}
	close(jobs)

	// 等待所有worker完成工作，另起一个协程，防止主协程阻塞
	go func() {
		wg.Wait()
		close(res)
	}()

	for resi := range res {
		fmt.Println(resi)
	}

	fmt.Println("整个进程结束")
}
