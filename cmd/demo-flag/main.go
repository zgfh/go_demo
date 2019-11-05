package main

import (
	goflag "flag"
	"fmt"
	flag "github.com/spf13/pflag"
)
/*
命令解析工具,目前一般用 cobra
pflag支持更多的特性
参考：https://www.jianshu.com/p/f9cf46a4de0e
*/

var flagvar int
func init() {
	flag.IntVar(&flagvar, "flagname", 11, "help message for flagname")
}
func main() {
	var ip *int = flag.Int("ip", 1234, "help message for flagname")
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.Parse()

	flag.BoolP("verbose", "v", false, "verbose output")
	flag.String("coolflag", "yeaah", "it's really cool flag")
	flag.Int("usefulflag", 777, "sometimes it's very useful")

	flag.CommandLine.MarkDeprecated("usefulflag", "please use --good-flag instead")
	flag.CommandLine.SortFlags=false
	flag.PrintDefaults()


	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)
}
