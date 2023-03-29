package main

import (
	"fmt"
	"log"
	"os"
	"tgb/ch4/github"
)

func main() {
	filtersList := os.Args[1:]
	if len(filtersList) == 0 {
		filtersList = []string{"repo:golang/go"}
	}

	result, err := github.SearchIssues(filtersList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d тем:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
