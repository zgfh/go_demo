package main

import (
	"fmt"
	"reflect"
)

/*


参考：
https://blog.golang.org/laws-of-reflection
https://golang.org/pkg/reflect/
https://wiki.jikexueyuan.com/project/go-web-programming/02.6.html
 */
func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	//
	editE := p.Elem()
	fmt.Println("kind is float64: ", editE.Kind() == reflect.Float64)
	fmt.Println("settability of v:", editE.CanSet())
	editE.SetFloat(7.1)



	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
