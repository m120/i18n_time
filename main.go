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

var now = time.Now()
var timeformat = time.RFC1123
var timejsonurl = "https://m120.github.io/timezone-json/timezone.json"

type result struct {
	Timezones []struct {
		TZ string `json:"tz"`
	} `json:"timezones"`
}

func flagUsage() {
	usageText := `
  i18n timezone
  
  Usage:
  -------------------------------------------
  - Default(No ARG): "Local & GMT"
  $ go run main.go

  - i18n: World Timezone List
  $ go run main.go i18n
	
  - "{TZ}": Specified Timezone 
  $ go run main.go "{TZ}"
	
    - Ex: "America/Chicago"
      $ go run main.go America/Chicago

  - help: This message. :-)
  $ go run main.go help
	
  `
	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func localtime() {
	fmt.Println(now.Format(timeformat), "\t:", now.Location())
}

func loadlocation(tz string) {
	loc, _ := time.LoadLocation(tz)
	nowloc := now.In(loc)
	fmt.Println(nowloc.Format(timeformat), "\t:", loc)
}

func tz(x string) {
	switch x {
	case "localgmt":
		// Local
		localtime()

		// UTC(GMT)
		loadlocation("Etc/UTC")
	case "i18n":
		// i18n: json get
		resp, err := http.Get(timejsonurl)
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
			loadlocation(tzs.TZ)
		}
	default:
		localtime()
		loadlocation(x)
	}
}

func main() {
	flag.Usage = flagUsage
	if len(os.Args) > 1 {
		osargs := os.Args[1]
		switch osargs {
		case "i18n":
			tz("i18n")
		case "help":
			flag.Usage()
		default:
			tz(osargs)
		}
	} else {
		tz("localgmt")
	}
}
