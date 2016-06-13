package main // prereqs记录了每个课程的前置课程
import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
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
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				// visitAll(m[item])
				mm := make(map[string]bool, len(m[item]))
				for _, p := range m[item] {
					mm[p] = true
				}
				visitAll(mm)
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool, len(m))
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}
