package main

import (
	"fmt"

	"github.com/DutkinTE/algorithm/arcsort"
)

func main() {
	arcs := []arcsort.Arc{
		{From: 3, To: 5},
		{From: 1, To: 2},
		{From: 2, To: 4},
		{From: 3, To: 1},
		{From: 2, To: 3},
	}

	fmt.Println("До сортировки:")
	for _, arc := range arcs {
		fmt.Printf("(%d, %d) ", arc.From, arc.To)
	}
	fmt.Println()

	sortedArcs := arcsort.SortArcs(arcs, 3)

	fmt.Println("После сортировки по начальной вершине:")
	for _, arc := range sortedArcs {
		fmt.Printf("(%d, %d) ", arc.From, arc.To)
	}
	fmt.Println()
}
