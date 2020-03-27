


### goswagger
通过工具goswagger可以自动化的生成api的定义，api的client，mode等
参考: https://goswagger.io/tutorial/todo-list.html

1. 通过 swagger.yml定义api
2. 生成api 
```bash
# 安装，这里采用docker，更多参考：https://goswagger.io/install.html
进入demo-goswagger当前目录
alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"

swagger validate ./swagger.yml
swagger generate server -A todo-list -f ./swagger.yml

```
3. restapi/configure_todo_list.go 里面定义真正的handle
