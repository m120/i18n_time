package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//fmt.Println(now)
	//
	//	const layout = "Now, Monday Jan 02 15:04:05 JST 2006"
	//	fmt.Println(now.Format(layout))

	t := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	const layout = "Mon Jan 02 15:04:05 2006 "
	//const layout = "2006 01 02 15:04:05"
	//time_p, err := time.Parse(layout, t)
	//fmt.Println("time:", time_p)

	//fmt.Println(now.Location(), ":\t", t.Format(layout))
	timeformat := time.RFC1123
	//fmt.Println(now.Location(), ":\t", t.Format(time.RFC1123))
	fmt.Println(now.Location(), ":\t", t.Format(timeformat))

	//t := time.Date(2001, time.November, 10, 23, 59, 58, 0, time.UTC)
	//t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	//fmt.Println(t)

	//loc, _ := time.LoadLocation("Asia/Tokyo")
	loc, _ := time.LoadLocation("Pacific/Auckland")
	nowloc := time.Now().In(loc)
	//fmt.Println("Location : ", loc, " Time : ", now_loc)
	fmt.Println(loc, ":\t", nowloc.Format(timeformat))
}
