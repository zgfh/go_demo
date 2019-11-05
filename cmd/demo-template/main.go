package main

import (
	"fmt"
	"html/template"
	"os"
)

/*
参考: https://www.kancloud.cn/cserli/golang/531904

go template:

模版语法:

{{.}} 表示当前对象，如user对象
{{.FieldName}} 表示对象的某个字段
{{range …}}{{end}} go中for…range语法类似，循环
{{with …}}{{end}} 当前对象的值，上下文
{{if …}}{{else}}{{end}} go中的if-else语法类似，条件选择,if 后只能出现bool值，不支持bool表达式

{{. | html}} //当前对象作为参数，传入内置函数html中
{{with $x := "output"}}{{$x}}{{end}} 定义变量
{{Say `Func`}} 模版函数
模版嵌套

例子:
 hello {{.UserName}}!
    {{range .Emails}}
        an email {{.}}
    {{end}}
    {{with .Friends}}
        {{range .}}
            my friend name is {{.Fname}}
        {{end}}
    {{end}}

*/
func Say(args ...interface{}) string {
	return fmt.Sprintf("%s %v", "Hello", args[0])
}

func main() {
	type person struct {
		Id      int
		Name    string
		Country string
	}

	data := person{Id: 1001, Name: "ma yun", Country: "China"}

	fmt.Println("data = ", data)

	tmpl := template.New("tmpl1")
	tmpl = tmpl.Funcs(template.FuncMap{"Say": Say})

	//tmpl, err := template.ParseFiles("./tmpl1.html")
	//template.Must: 检测模板是否正确：大括号是否匹配，注释是否正确关闭，变量是否正确,如果错误，panic
	template.Must(tmpl.Parse("Hello {{.Name}}, Welcome to go programming...\n {{Say .Name}} \n {{`Pipe` | Say}}"))

	tmpl.Execute(os.Stdout, data)
}
