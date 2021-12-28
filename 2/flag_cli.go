package main

/**
go run flag_cli
go run flag_cli -s -name Buttercup
go run flag_cli --spanish -name Buttercup
*/

import (
	"flag"
	"fmt"
)

// Creates a new variable from a flag
var name = flag.String("name", "World", "A name to say hello to.")

// New variable to store flag value
var spanish bool
var help bool

func init() {
	// Sets variable to flag value
	flag.BoolVar(&help, "help", false, "Get Help")
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
}

func main() {
	flag.Parse()

	if help == true {
		flag.VisitAll(func(flag *flag.Flag) {
			format := "\t-%s: %s (Default: '%s')\n"
			fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
		})
		return
	}

	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
