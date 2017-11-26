package main

import (
	"fmt"
)

var prereqs = map[string][]string {
	"algorithms": {"data structures"},
	"calculus": {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures": {"discrete math"},
	"databases": {"data structures"},
	"discrete math": {"intro to programming"},
	"formal languages": {"discrete math"},
	"networks": {"operating systems"},
	"operating systems": {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(items map[string][]string) {
		for item, required := range items {
			if !seen[item] {
				seen[item] = true
				for _, ri := range required {
					if _, ok := items[ri]; ok {
						set := map[string][]string{
							ri: items[ri],
						}
						visitAll(set)
					} else {
						if !contains(order, ri) {
							order = append(order, ri)
						}
					}
				}
				if !contains(order, item) {
					order = append(order, item)
				}
			}
		}
	}

	visitAll(m)
	return order
}

func contains(s []string, search string) bool {
	for _, str := range s {
		if str == search {
			return true
		}
	}
	return false
}