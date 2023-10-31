package main

import (
	"fmt"
	"log"

	"github.com/brandenc40/safer"
)

func main() {
	client := safer.NewClient()

	snapshot, err := client.GetCompanyByMCMX("133655")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%#v", snapshot)
}
