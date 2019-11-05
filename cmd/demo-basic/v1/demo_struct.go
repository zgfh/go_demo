package v1

import "fmt"

type Books struct {
	Name    string
	Title   string
	Author  string
	Subject string
	Book_id int
}

func (b *Books) ToStr() string {
	return b.Name + b.Title
}

func (b *Books) Test() {
	if b == nil {
		fmt.Println("b not init,but can call his func")
	} else {
		fmt.Println("b is inited  ")
	}
}

func StructMain() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */
	var Book4 *Books
	Book1.Test()
	Book4.Test()
	Book3 := Books{
		Title: "test",
	}

	/* book 1 描述 */
	Book1.Title = "Go 语言"
	Book1.Author = "www.runoob.com"
	Book1.Subject = "Go 语言教程"
	Book1.Book_id = 6495407

	/* book 2 描述 */
	Book2.Title = "Python 教程"
	Book2.Author = "www.runoob.com"
	Book2.Subject = "Python 语言教程"
	Book2.Book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1  : %s\n", Book1.ToStr())
	fmt.Printf("Book 1 Author : %s\n", Book1.Author)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 Title : %s\n", Book2.ToStr())
	fmt.Printf("Book 2 Book_id : %d\n", Book2.Book_id)

	fmt.Println("Book 3 Title: %s", Book3.Title)
}
