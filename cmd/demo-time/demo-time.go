package main

import (
	"fmt"
	"time"
)




func main() {
	now:=time.Now()
	feature:=now.Add(10*time.Minute)
	fmt.Printf("%v before %v : %v",now.String(),feature.String(),now.Before(feature))
}
