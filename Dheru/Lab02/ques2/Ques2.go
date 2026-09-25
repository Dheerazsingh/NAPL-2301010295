package main

import "fmt"

func main() {
	fmt.Println("SLICE OPERATIONS")

	scores := []int{100, 200, 300}
	fmt.Println("Initial:", scores)

	scores = append(scores, 400)
	fmt.Println("After add:", scores)

	scores = append(scores[:0], scores[1:]...)
	fmt.Println("After remove index 0:", scores)

	scores[1] = 250
	fmt.Println("After update index 1:", scores)

	fmt.Println()
	fmt.Println("MAP OPERATIONS")

	subjectMarks := map[string]int{
		"Math":    85,
		"Science": 90,
	}
	fmt.Println("Initial:", subjectMarks)

	subjectMarks["English"] = 78
	fmt.Println("After insert English:", subjectMarks)

	delete(subjectMarks, "Science")
	fmt.Println("After delete Science:", subjectMarks)

	value, found := subjectMarks["Math"]
	fmt.Println("Lookup Math:", value, found)

	value, found = subjectMarks["Science"]
	fmt.Println("Lookup Science:", value, found)
}
