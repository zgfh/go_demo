package main

import (
	"fmt"
	"sync"
	"time"
)

/*
参考: https://studygolang.com/articles/11038?fr=sidebar
1. sync.Mutex: 一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
2. sync.RWMutex: 读写锁与互斥锁最大的不同就是可以分别对 读、写 进行锁定。一般用在大量读操作、少量写操作的情况：
同时只能有一个 goroutine 能够获得写锁定。
同时可以有任意多个 gorouinte 获得读锁定。
同时只能存在写锁定或读锁定（读和写互斥）。

读锁定（RLock），对读操作进行锁定
读解锁（RUnlock），对读锁定进行解锁
写锁定（Lock），对写操作进行锁定
写解锁（Unlock），对写锁定进行解锁
3. WaitGroup 用于等待一组 goroutine 结束
4. sync.Once 对象可以使得函数多次调用只执行一次
5. Cond 实现一个条件变量，即等待或宣布事件发生的 goroutines 的会合点。
func NewCond(l Locker) *Cond
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()
6. sync.Pool 可以作为临时对象的保存和复用的集合。

*/

type Locker interface {
	Lock()
	Unlock()
}

// 如果没有锁，并发的时候就会造成错误的变量处理，如下，会出现 a 的值是相同的情况
func demo_nolock() {
	var a = 0
	var wg sync.WaitGroup

	fmt.Printf("a 的值可能有相同的\n")
	// 启动 100 个协程，需要足够大
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}

	wg.Wait()
	fmt.Printf("goroutine a 不一定是1000: a=%d\n",  a)
	// 等待 1s 结束主程序
	// 确保所有协程执行完
	time.Sleep(time.Second*3)
}

func demo_lock() {
	var a = 0
	var wg sync.WaitGroup

	fmt.Printf("a 的值可能有相同的\n")
	// 启动 100 个协程，需要足够大
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			lock.Lock()
			defer lock.Unlock()
			a += 1
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}

	wg.Wait()
	fmt.Printf("goroutine a 一定是1000: a=%d\n",  a)
	// 等待 1s 结束主程序
	// 确保所有协程执行完
	time.Sleep(time.Second*3)
}


func demo_doonce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func main() {
	fmt.Printf("")
	demo_nolock()
	demo_lock()
	demo_doonce()
}
