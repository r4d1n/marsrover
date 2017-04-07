package main

import (
	"log"
	"os"

	"github.com/r4d1n/marsrover"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please Specify a Rover Name")
	}
	client := marsrover.NewClient("DEMO_KEY", "")
	m, err := client.GetManifest(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	log.Println(m)
}
