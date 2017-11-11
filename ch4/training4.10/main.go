package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/masato-d/gopl/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	var inOneMonth []*github.Issue
	var inOneYear []*github.Issue
	var overOneYear []*github.Issue
	for _, item := range result.Items {
		switch {
		case time.Now().AddDate(0, -1, 0).Before(item.CreatedAt):
			inOneMonth = append(inOneMonth, item)
		case time.Now().AddDate(-1, 0, 0).Before(item.CreatedAt):
			inOneYear = append(inOneYear, item)
		default:
			overOneYear = append(overOneYear, item)
		}
	}

	fmt.Println("================== in a month ===================")
	for _, item := range inOneMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("")
	fmt.Println("================== in a year ===================")
	for _, item := range inOneYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	fmt.Println("")
	fmt.Println("================== over a year ===================")
	for _, item := range overOneYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}