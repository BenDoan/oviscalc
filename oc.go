package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"time"
)

func main() {
	usage := "Usage: oc add <time> <time>"
	args, _ := docopt.Parse(usage, nil, true, "", false)

	times := args["<time>"].([]string)

	time1, err := time.Parse("15:04", times[0])
	time2, err2 := time.Parse("15:04", times[1])

	if err == nil && err2 == nil {
		fmt.Println(time1.Sub(time2))
	} else {
		if err != nil {
			fmt.Println(err)
		}

		if err2 != nil {
			fmt.Println(err)
		}
	}

}
