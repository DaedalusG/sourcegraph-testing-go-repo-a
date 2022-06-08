package main

import (
	"fmt"

	gorepoa "github.com/sourcegraph-testing/go-repo-a"
)

func main() {
	fmt.Println("This is go-repo-A")

	people := []gorepoa.Person{
		{Name: "Alice"},
		{Name: "Sophie"},
		{Name: "Butter"},
		{Name: "Penelope"},
	}

	fmt.Println("People:")
	for _, p := range people {
		fmt.Printf("- %s\n", p.Name)
	}

	fmt.Println("Result of deep call:")
	result := gorepoa.FunctionInRepoA(people)
	fmt.Println(result)
}
