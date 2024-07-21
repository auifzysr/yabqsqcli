package main

import (
	"log"

	"github.com/auifzysr/yabqsqcli/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
