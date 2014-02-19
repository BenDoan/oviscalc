package main

import (
	"fmt"
	"github.com/docopt/docopt.go"
	"time"
)

func main() {
	usage := "Usage: oc add <time> <duration>"
	args, _ := docopt.Parse(usage, nil, true, "", false)
	fmt.Println(args)

	fmt.Println(time.Parse("15:04", "2:03"))
}
