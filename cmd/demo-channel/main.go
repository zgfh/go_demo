package main

import (
	"fmt"
	"time"
)

/*
https://colobu.com/2016/04/14/Golang-Channels/
https://i6448038.github.io/2019/03/23/go-select-principle/

定义: 箭头的指向就是数据的流向,总是优先和最左边的类型结合
   双向,缓存大小为100: ch := make(chan int,100)
   发送给管道int数据: ch := make(chan<- int); ch <- 3
   从管道接受int数据: ch := make(<-chan int); x:= <-ch

使用:
   发送: ch <- 3
   接收(ok 为false 表示关闭了，ok可省略):  x, ok := <-ch
   关闭: close(ch)
   循环: for i := range c {fmt.Println(i)};
   select: 遍历所有case的channel，查看其是否可读或者可写,如果其中的channel可读或者可写，则解锁所有channel，并返回对应的channel数据
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				fmt.Println("quit")
				return
			}
		}
   超时: 通过 time.After 实现超时功能


*/

func demo() {
	c1 := make(chan string, 1)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- fmt.Sprintf("result %v", i)
		}
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
		close(c1)
	}()

	for {
		select {
		case res, ok := <-c1:
			if ok == false {
				fmt.Println("finish")
				return
			}
			fmt.Println(res)
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1")
		}

	}
}

func kick(c chan<- interface{}) {
	select {
	case c <- nil:
		fmt.Println("xxx")
	default:
		fmt.Println("default")
	}
}

func main() {
	ch:=make(chan interface{}, 1)
	kick(ch)
	kick(ch)
	<-ch
	kick(ch)
	fmt.Println("demo 1")
	demo()

	fmt.Println("demo 2")
	c := make(chan int, 10)
	c <- 1
	c <- 2
	close(c)
	for i := range c {
		fmt.Println(i)
	}

}
