package arcsort

type Arc struct {
	From int
	To   int
}

func SortArcs(arcs []Arc, n int) []Arc {
	m := len(arcs)

	freq := make([]int, n+1) // Индексы от 1 до n.
	for _, arc := range arcs {
		freq[arc.From]++
	}

	pos := make([]int, n+1)
	pos[1] = 0
	for i := 2; i <= n; i++ {
		pos[i] = pos[i-1] + freq[i-1]
	}

	sorted := make([]Arc, m)

	tempPos := make([]int, len(pos))
	copy(tempPos, pos)
	for _, arc := range arcs {
		index := tempPos[arc.From]
		sorted[index] = arc
		tempPos[arc.From]++
	}

	return sorted
}
