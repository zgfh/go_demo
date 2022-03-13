package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

/*
pprof 是go的性能分析工具

代码开启方式
```
引入 _ "net/http/pprof" 即可(利用init, 对http 增加了/debug/pprof/ api)

```


# 1.内存分析: 可以分析出具体哪里占用了内存

需要下载: https://graphviz.org/download/
go tool pprof -http=:8080 http://localhost:9090/debug/pprof/heap

如果是生产环境，也可以下载后本地分析:
1. wget -O profile_name-heap.heap http://localhost:9090/debug/pprof/heap
2. go tool pprof -http=:8080 profile_name-heap.heap

对比两个时刻的内存变化
1. 首先:wget -O profile_name-heap-1.heap http://localhost:9090/debug/pprof/heap
1. 一段时间后，再此执行: wget -O profile_name-heap-2.heap http://localhost:9090/debug/pprof/heap
2. 对比两个文件: go tool pprof -http=:8080 --base profile_name-heap-1.heap profile_name-heap-2.heap


https://www.robustperception.io/optimising-prometheus-2-6-0-memory-usage-with-pprof

*/

var test_heap *[]int

func test_heap1(w http.ResponseWriter, r *http.Request) {
	container := make([]int, 8)

	log.Println("> loop.")
	// slice会动态扩容，用它来做堆内存的申请
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, time.Now().Nanosecond())
	}
	w.Write([]byte(fmt.Sprintf("array: %v", len(container))))

	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}

func test_heap2(w http.ResponseWriter, r *http.Request) {
	container := make([]int, 32*1000*10)

	log.Println("> loop.")
	for i := 0; i < len(container); i++ {
		container[i] = time.Now().Nanosecond()
	}
	// 已知扩大 test_heap,每次10M多
	if test_heap != nil {
		*test_heap = append(*test_heap, container...)
	} else {
		test_heap = &container
	}

	w.Write([]byte(fmt.Sprintf("array: %v", len(*test_heap))))

	log.Println("< loop.")
	// container在f函数执行完毕后不再使用
}

func main() {
	fmt.Println(os.TempDir())

	http.HandleFunc("/", test_heap2)

	ip := "0.0.0.0:6060"
	fmt.Println("listen: %v", ip)
	fmt.Println("see pprof: http://127.0.0.1:6060/debug/pprof/")
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
