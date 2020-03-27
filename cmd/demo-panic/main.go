package main

import (
	"errors"
	"fmt"
)
/**
https://blog.golang.org/defer-panic-and-recover

 */

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f:", r)
		}
	}()
	//log.Panicf("error","demo panic")
	panic(errors.New("demo panic"))
}
