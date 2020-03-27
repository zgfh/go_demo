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


二十五个关键字:
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var

var和const参考2.2Go语言基础里面的变量和常量申明
package和import已经有过短暂的接触
func 用于定义函数和方法
return 用于从函数返回
defer 用于类似析构函数
go 用于并发
select 用于选择不同类型的通讯
interface 用于定义接口，参考2.6小节
struct 用于定义抽象数据类型，参考2.5小节
break、case、continue、for、fallthrough、else、if、switch、goto、default这些参考2.3流程介绍里面
chan用于channel通讯
type用于声明自定义类型
map用于声明map类型数据
range用于读取slice、map、channel数据
https://wiki.jikexueyuan.com/project/go-web-programming/02.8.html



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
