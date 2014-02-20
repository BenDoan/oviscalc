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

	time1, _ := time.Parse("15:04", times[0])
	time2, _ := time.Parse("15:04", times[1])

	fmt.Println(time2.Sub(time1))
}
