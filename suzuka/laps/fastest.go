package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
[
  {
    "number": "53",
    "Timings": [
      {
        "driverId": "michael_schumacher",
        "position": "1",
        "time": "1:27.911"
      }
		]
	}
]
*/

type lap struct {
	Number  string
	Timings []Timing
}
type Timing struct {
	DriverId string
	Position string
	Time     string
}

func main() {
	files := make([]string, 0)
	for i, a := range os.Args {
		if i == 0 {
			continue
		}
		files = append(files, a)
	}
	if len(files) == 0 {
		log.Fatal("pass json files of lap data")
	}
	for _, file := range files {
		var (
			fastestLap    lap
			fastestTiming Timing
		)
		j, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("%s was not readable", file)
		}
		laps := make([]lap, 0)
		err = json.Unmarshal(j, &laps)
		if err != nil {
			log.Fatalf("%s did not contain an array of lap data", file)
		}
		var fastest = 999 * time.Minute
		for _, l := range laps {
			for _, t := range l.Timings {
				var tmp, tmp2 []string
				tmp = strings.Split(t.Time, ":")
				if len(tmp) != 2 {
					log.Fatalf("%s did not have mm:ss", t.Time)
				}
				mm, err := strconv.Atoi(tmp[0])
				if err != nil {
					log.Fatalf("non-integer minutes: %s", tmp[0])
				}
				tmp2 = strings.Split(tmp[1], ".")
				if len(tmp) != 2 {
					log.Fatalf("%s did not have ss.mmm", tmp[1])
				}
				ss, err := strconv.Atoi(tmp2[0])
				if err != nil {
					log.Fatalf("non-integer seconds: %s", tmp2[0])
				}
				ms, err := strconv.Atoi(tmp2[1])
				if err != nil {
					log.Fatalf("non-integer milliseconds: %s", tmp2[1])
				}

				var d time.Duration
				d = time.Duration(mm)*time.Minute + time.Duration(ss)*time.Second + time.Duration(ms)*time.Millisecond
				if d < fastest {
					fastest = d
					fastestLap = l
					fastestTiming = t
				}
			}
		}
		fmt.Printf(`{
	year:"%s",
	timeDisplay:"%s",
	timeMilliseconds: %v,
	lapNumber: "%s",
	driver:"%s"
}`+"\n", file, fastestTiming.Time, (fastest.Nanoseconds() / 10E6), fastestLap.Number, fastestTiming.DriverId)
	}
}
