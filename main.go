package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	timeformat := time.RFC1123

	// Local
	fmt.Println(now.Format(timeformat), "\t:", now.Location())

	for _, l := range []string{"Pacific/Auckland", "Asia/Tokyo"} {
		loc, _ := time.LoadLocation(l)
		nowloc := now.In(loc)
		fmt.Println(nowloc.Format(timeformat), "\t:", loc)
	}
}
