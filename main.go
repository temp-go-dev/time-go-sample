package main

import (
	"fmt"

	com "github.com/temp-go-dev/time-go-sample/common"
)

func main() {

	// date := "2019-12-22 23:01:00"
	// layout := com.Timeformat("yyyy-MM-dd HH:mm:ss")

	// t, _ := time.Parse(layout, date)
	// com.SetNowForDebug(t)

	fmt.Println(com.GetNow())
	// fmt.Println(com.GetNowString(layout))
}
