package arcsort

import (
	"reflect"
	"testing"
)

func TestSortArcs(t *testing.T) {
	arcs := []Arc{
		{From: 3, To: 5},
		{From: 1, To: 2},
		{From: 2, To: 4},
		{From: 3, To: 1},
		{From: 2, To: 3},
	}

	sorted := SortArcs(arcs, 3)
	expected := []Arc{
		{From: 1, To: 2},
		{From: 2, To: 4},
		{From: 2, To: 3},
		{From: 3, To: 5},
		{From: 3, To: 1},
	}

	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, sorted)
	}
}
