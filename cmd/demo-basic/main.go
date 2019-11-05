package main

import (
	"fmt"
	"os"
)
import "github.com/zgfh/go-demo/cmd/demo-basic/v1"

/*
访问权限:
  因为Go是以大小写来区分是公有还是私有，但都是针对包级别的，所以在包内所有的都能访问，而方法绑定本身只能绑定包内的类型，
  所以方法可以访问接收者所有成员。如果是包外调用某类型的方法，则需要看方法名是大写还是小写，大写能被包外访问，小写只能被包内访问。

https://segmentfault.com/a/1190000011446643
*/

func B(per *v1.Books) {
	per.Name = "1"
	fmt.Println("B", *per)
}
func demo指针() {

	// 指针
	a := v1.Books{Name: "1"}
	fmt.Println("A", a)
	B(&a)
	fmt.Println("A", a)
}

func demoSwitch() {
	fmt.Println("demoSwitch")
	a := 1
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("defualt")
	}
	fmt.Println("end--demoSwitch")
}

func demoFor() {
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func main() {
	demoSwitch()
	v1.DemoFunc()
	v1.StructMain()
	demo指针()
	fmt.Println(os.TempDir())
	demoFor()
	//
}
