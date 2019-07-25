package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Invalid arguments...")
	}

	for arg := range args {
		fmt.Println(arg)
	}
}
