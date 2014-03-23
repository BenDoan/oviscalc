package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"time"
)

func main() {
	usage := "usage: oc [--zulu] [diff <time> <time>]"
	args, _ := docopt.Parse(usage, nil, true, "", false)

	formatString := ""
	if args["--zulu"].(bool) {
		formatString = time.RFC3339
	} else {
		formatString = "15:04"
	}

	if args["diff"].(bool) {
		Diff(args["<time>"].([]string), formatString)
	}
}

func Diff(times []string, formatString string) {
	time1, err := time.Parse(formatString, times[0])
	time2, err2 := time.Parse(formatString, times[1])

	if err == nil && err2 == nil {
		fmt.Println(TimeAbs(time1.Sub(time2)))
	} else {
		if err != nil {
			fmt.Println(err)
		}

		if err2 != nil {
			fmt.Println(err)
		}
	}
}

func TimeAbs(dur time.Duration) time.Duration {
	if dur > 0 {
		return dur
	} else {
		return -1 * dur
	}
}
