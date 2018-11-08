/*
	Exercise 4.10: Modify issues to report the results in age categories, say less than a month old,
less than a year old, and more than a year old.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("now: ", time.Now())
	now := time.Now()

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("over a year")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 > 365 {
			fmt.Printf("%s #%-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("over a month;and withini a year")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 > 30 && now.Sub(item.CreatedAt).Hours()/24 < 365 {
			fmt.Printf("%s #%-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}

	fmt.Println("within a month")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 < 30 {
			fmt.Printf("%s #%-5d %9.9s %.55s\n",
				item.CreatedAt, item.Number, item.User.Login, item.Title)
		}
	}

}
