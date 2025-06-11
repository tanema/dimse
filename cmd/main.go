package main

import (
	"fmt"
	"log"

	"github.com/tanema/dimse"
)

func checkErr(scope string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", scope, err)
	}
	log.Println(scope)
}

func main() {
	client, err := dimse.Connect("www.dicomserver.co.uk:104")
	if err != nil {
		log.Fatalf("connection err: %v", err)
	}
	for {
		select {
		case err := <-client.Errors():
			fmt.Println("err", err)
		case evt := <-client.Events():
			fmt.Println("evt", evt)
		}
	}
}
