package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type result struct {
	version   string `json:"version"`
	Timezones []struct {
		REGION      string `json:"region"`
		ZONE1       string `json:"zone_1"`
		ZONE2       string `json:"zone_2"`
		CODE        string `json:"code"`
		TZ          string `json:"tz"`
		COORDINATES string `json:"coordinates"`
	} `json:"timezones"`
}

func main() {
	now := time.Now()
	timeformat := time.RFC1123

	// Local
	fmt.Println(now.Format(timeformat), "\t:", now.Location())

	// GMT(UTC)
	gmt, _ := time.LoadLocation("Europe/London")
	nowgmt := now.In(gmt)
	fmt.Println(nowgmt.Format(timeformat), "\t:", gmt)
	fmt.Println("-----------------------------------------------------")

	// i18n: json read
	raw, err := ioutil.ReadFile("./tz.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bytes := []byte(raw)
	var tzd result
	json.Unmarshal(bytes, &tzd)

	for _, tzs := range tzd.Timezones {
		loc, _ := time.LoadLocation(tzs.TZ)
		nowloc := now.In(loc)
		fmt.Println(nowloc.Format(timeformat), "\t:", loc)
	}
}
