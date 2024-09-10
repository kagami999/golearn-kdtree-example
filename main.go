package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/kdtree"
	"github.com/sjwhitworth/golearn/metrics/pairwise"
)

func main() {
	data := [][]float64{
		{2.1, 3.5},
		{4.7, 1.2},
		{3.6, 4.9},
		{6.2, 3.4},
	}
	tg := []float64{3.0, 4.0}
	tr := kdtree.New()
	e := tr.Build(data)
	if e != nil {
		fmt.Print(e)
		return
	}
	in, _, e := tr.Search(1,
		&pairwise.Euclidean{},
		tg)
	if e != nil {
		fmt.Print(e)
		return
	}
	fmt.Print(in) // index: [0] => {2.1, 3.5}
}
