package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	First string `json:"firstName"`
	Last  string `json:"lastName"`
}

// These annotations are all legal. The Go parser will correctly handle this
// and the JSON encoder will be able to pick out which of those apply to it.
type RobustName struct {
	First string `json:"firstName" xml:"FirstName"`
	Last  string `json:"lastName,omitempty"`
	Other string `not,even.a=tag`
}

func main() {
	n := &Name{"Nacho", "Herrera"}
	data, _ := json.Marshal(n)
	fmt.Printf("%s\n", data)
}
