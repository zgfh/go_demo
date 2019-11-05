package v1

import "fmt"

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

// ... 不定长变参
func A(a ...int) {
	fmt.Println(a)
}

func DemoFunc() {
	fmt.Println("max: %v", max(1, 2))
	A(1, 2, 3, 4, 5)

	// 将一个函数赋值一个变量，该变量是函数类型
	a := func() {
		fmt.Println("匿名函数")
	}
	// 调用
	a()

}
