package main

import (
	"log"
)

/*

所有的类型都实现了空interface
Comma-ok断言: value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。


参考：https://wiki.jikexueyuan.com/project/go-web-programming/02.6.html
*/
type Person struct {
	name string
}

func (p *Person) printName() {
	log.Printf("name: %s", p.name)
}

func (p *Person) New(name string) ( *Person) {
	return &Person{name: name}
}

var person1 Person

func main() {
	person1.printName()
	person2 := person1.New("alan")
	person2.printName()

}
