package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type result struct {
	version   string `json:"version"`
	Timezones []struct {
		//	REGION      string `json:"region"`       //no use
		//	ZONE1       string `json:"zone1"`        //no use
		//	ZONE2       string `json:"zone2"`        //no use
		//	CODE        string `json:"code"`         //no use
		TZ string `json:"tz"`
		//	COORDINATES string `json:"coordinates"`  //no use
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
	fmt.Printf("%v\n", "--------------------------------------------------------")

	// i18n: json get
	resp, err := http.Get("https://m120.github.io/testsite/tz.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	respbody, err := ioutil.ReadAll(resp.Body)
	var tzd result
	json.Unmarshal(respbody, &tzd)

	for _, tzs := range tzd.Timezones {
		loc, _ := time.LoadLocation(tzs.TZ)
		nowloc := now.In(loc)
		fmt.Println(nowloc.Format(timeformat), "\t:", loc)
	}
}
