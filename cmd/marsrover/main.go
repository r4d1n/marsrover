package main

import (
	"log"
	"os"
	"strconv"

	"github.com/r4d1n/marsrover"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please Specify a Rover Name")
	}
	client := marsrover.NewClient(os.Getenv("NASA_API_KEY"))
	if len(os.Args) < 3 {
		m, err := client.GetManifest(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		log.Println(m)
	} else {
		sol, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		m, err := client.GetImagesBySol(os.Args[1], sol)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("completed request:\n", m)
	}
}
