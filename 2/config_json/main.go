package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, _ := os.Open("config.json")
	defer file.Close()

	conf := Configuration{}
	err := json.NewDecoder(file).Decode(&conf)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Path: %s", conf.Path)
}
