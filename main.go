package gorepoa

import gorepob "github.com/sourcegraph-testing/go-repo-b"

type Person struct {
	Name string
}

func FunctionInRepoA(people []Person) string {
	var names []string

	for _, p := range people {
		names = append(names, p.Name)
	}

	return gorepob.FunctionInRepoB(names)
}

func privateFunction(num int) string {
	return fmt.Sprintf("this is a private Function. num=%d", num)
}
