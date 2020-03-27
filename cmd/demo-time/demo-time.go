package main

import (
	"fmt"
	"time"
)




func main() {
	now:=time.Now()
	fmt.Printf("now: %v \n",now)
	feature:=now.Add(10*time.Minute)
	fmt.Printf("%v before %v : %v",now.String(),feature.String(),now.Before(feature))
}
