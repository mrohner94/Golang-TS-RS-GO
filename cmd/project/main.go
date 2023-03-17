package main

import (
	"fmt"
	"log"

	"github.com/mrohner94/compare/pkg/config"
)

func main() {
	fmt.Println("Hello")
	opts, err := config.GetOpts()
	if err != nil {
		log.Fatalf("unable to get options %v", err)
	}
	fmt.Printf("opts %+v", opts)
}
