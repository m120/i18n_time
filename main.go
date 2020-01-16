package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type tzdata struct {
	REGION      string `json:"region"`
	ZONE1       string `json:"zone_1"`
	ZONE2       string `json:"zone_2"`
	CODE        string `json:"code"`
	TZ          string `json:"tz"`
	COORDINATES string `json:"coordinates"`
}

func main() {
	now := time.Now()
	timeformat := time.RFC1123

	// Local
	fmt.Println(now.Format(timeformat), "\t:", now.Location())

	// json read
	raw, err := ioutil.ReadFile("./tz.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var tzd []tzdata
	json.Unmarshal(raw, &tzd)

	for _, tzs := range tzd {
		//loc, _ := time.LoadLocation(tz.REGION + "/" + tz.ZONE1)
		loc, _ := time.LoadLocation(tzs.TZ)
		nowloc := now.In(loc)
		fmt.Println(nowloc.Format(timeformat), "\t:", loc)
	}
}
