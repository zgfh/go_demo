package main

import (
	"log"
)

/*

所有的类型都实现了空interface
Comma-ok断言: value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。


参考：https://wiki.jikexueyuan.com/project/go-web-programming/02.6.html
*/

type Person interface {
	Say(work string)
}

func SayHi(person Person) {
	person.Say("hi")
}

func testInterfaceVariable() {
	// interface变量存储的类型
	type Element interface{}
	type List [] Element

	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = NewGirl("test")
	for index, element := range list {
		if value, ok := element.(int); ok {
			log.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			log.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			log.Printf("list[%d] is a Person and its value is %s\n", index, value)
		} else {
			log.Printf("list[%d] is of a different type\n", index)
		}
	}
}

func main() {
	girl := NewGirl("babi")
	SayHi(girl)

	testInterfaceVariable()

}
