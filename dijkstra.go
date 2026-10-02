package main

import "fmt"

func dijkstra(g map[string]map[string]int, s, e string) (int, []string) {
	dist := map[string]int{s: 0}
	prev := map[string]string{}
	done := map[string]bool{}
	for {
		u, best := "", 1<<30
		for n, d := range dist {
			if !done[n] && d < best {
				u, best = n, d
			}
		}
		if u == "" || u == e {
			break
		}
		done[u] = true
		for v, w := range g[u] {
			if d, ok := dist[v]; !ok || best+w < d {
				dist[v], prev[v] = best+w, u
			}
		}
	}
	path := []string{}
	for n := e; n != ""; n = prev[n] {
		path = append([]string{n}, path...)
	}
	return dist[e], path
}

func visitAll(g map[string]map[string]int, mid []string) (int, []string) {
	best, route := 1<<30, []string{}
	var gen func(int)
	gen = func(i int) {
		if i == len(mid) {
			full := append([]string{"C"}, append(append([]string{}, mid...), "E")...)
			t := 0
			for j := 0; j+1 < len(full); j++ {
				w, ok := g[full[j]][full[j+1]]
				if !ok {
					return // an edge doesn't exist, this route is invalid
				}
				t += w
			}
			if t < best {
				best, route = t, full
			}
			return
		}
		for j := i; j < len(mid); j++ {
			mid[i], mid[j] = mid[j], mid[i]
			gen(i + 1)
			mid[i], mid[j] = mid[j], mid[i]
		}
	}
	gen(0)
	return best, route
}

func main() {
	g := map[string]map[string]int{
		"C": {"A": 3, "F": 4},
		"A": {"B": 4},
		"B": {"F": 6, "E": 5},
		"F": {"G": 4, "B": 6},
		"G": {"E": 3},
		"E": {},
	}
	d, p := dijkstra(g, "C", "E")
	fmt.Println("Plain shortest:", p, "=", d, "(skips nodes!)")
	d, p = visitAll(g, []string{"A", "B", "F", "G"})
	fmt.Println("Visit all nodes:", p, "=", d)
}
