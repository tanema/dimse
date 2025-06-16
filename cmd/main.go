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
	checkErr("echo", client.Echo())
	fmt.Println("echo successfully returned")
	client.Close()
}
