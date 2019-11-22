package main

import "fmt"
import "sort"

type Train struct {
	Train   string
	Minutes int
	Score   int
}

func main() {
	// slice of structs we'll return
	targets := []Train{}

	// create one and add it
	var a Train
	a.Train = "mont"
	a.Minutes = 5
	a.Score = 0
	targets = append(targets, a)

	fmt.Println(targets)

	// create another and add it
	var b Train
	b.Train = "mont"
	b.Minutes = 1
	b.Score = 0
	targets = append(targets, b)

	fmt.Println(targets)

	// create another and add it
	var c Train
	c.Train = "mont"
	c.Minutes = 3
	c.Score = 0
	targets = append(targets, c)

	fmt.Println(targets)

	// sort the slice by Minutes
	fmt.Println("\nsorting")
	sort.Slice(targets, func(i, j int) bool { return targets[i].Minutes < targets[j].Minutes })

	fmt.Println(targets)

}
