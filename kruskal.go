package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	u, v string
	w    int
}

var parent map[string]string

func find(x string) string {
	if parent[x] != x {
		parent[x] = find(parent[x])
	}
	return parent[x]
}

func kruskal(edges []Edge) ([]Edge, int) {
	sort.Slice(edges, func(i, j int) bool { return edges[i].w < edges[j].w })

	parent = map[string]string{}
	for _, e := range edges {
		parent[e.u], parent[e.v] = e.u, e.v
	}

	mst, total := []Edge{}, 0
	for _, e := range edges {
		ru, rv := find(e.u), find(e.v)
		if ru != rv { // no cycle -> take it
			parent[ru] = rv
			mst = append(mst, e)
			total += e.w
		}
	}
	return mst, total
}

func main() {
	edges := []Edge{
		{"B", "C", 10},
		{"C", "A", 3},
		{"B", "D", 6},
		{"D", "C", 2},
		{"C", "E", 4},
		{"D", "E", 1},
	}

	mst, total := kruskal(edges)
	for _, e := range mst {
		fmt.Printf("%s - %s : %d\n", e.u, e.v, e.w)
	}
	fmt.Println("Total weight:", total)
}
