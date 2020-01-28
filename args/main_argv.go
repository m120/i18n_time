package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type result struct {
	//version   string `json:"version"` //no use
	Timezones []struct {
		//	REGION      string `json:"region"`       //no use
		//	ZONE1       string `json:"zone1"`        //no use
		//	ZONE2       string `json:"zone2"`        //no use
		//	CODE        string `json:"code"`         //no use
		TZ string `json:"tz"`
		//	COORDINATES string `json:"coordinates"`  //no use
	} `json:"timezones"`
}

// reference from: https://golang.hateblo.jp/entry/2018/10/22/080000
// https://program.sakaiboz.com/golang/standard-package/flag-commandline/
//https://qiita.com/nightswinger/items/8abc3ee7db41a3365784
func flagUsage() {
	usageText := `
  timezone i18n
  
  Usage:
  ----
  default(no arg):	local and GMT(Europe/London)
  i18n:				World Timezone List
  help:				this message
  `

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func localgmt() {
	now := time.Now()
	timeformat := time.RFC1123

	// Local
	fmt.Println(now.Format(timeformat), "\t:", now.Location())

	// GMT(UTC)
	gmt, _ := time.LoadLocation("Europe/London")
	nowgmt := now.In(gmt)
	fmt.Println(nowgmt.Format(timeformat), "\t:", gmt)

}

func main() {
	flag.Usage = flagUsage

	now := time.Now()
	timeformat := time.RFC1123

	if len(os.Args) > 1 {
		osargs := os.Args[1]

		switch osargs {
		case "i18n":
			// i18n: json get
			resp, err := http.Get("https://m120.github.io/timezone-json/timezone.json")
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
		case "help":
			flag.Usage()
		default:
			flag.Usage()
		}
	} else {
		localgmt()
	}
}
