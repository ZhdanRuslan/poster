package main

import (
	"fmt"
	"log"
	"poster/domain"
)

func main() {

	result, err := domain.SearchIssues()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title:" + result.Title + "\n" + "Genre:" + result.Genre)
}
