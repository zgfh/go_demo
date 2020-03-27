package main

/**
type xx func()  :xx(a)，a强制转换为xx类型	 https://golangtc.com/t/53b8c2e8320b522f430000c4
*/

//定义结构体
type demeStruct struct {
	test string
}

type demoInterFace interface {
	test()
}

//重新定义关键字
type Mystring string

func (m *Mystring) test() {

}


//重新定义函数 https://golangtc.com/t/53b8c2e8320b522f430000c4
type MyFunction func() string


func main() {

}
