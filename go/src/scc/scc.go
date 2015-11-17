/*
	scc
	Kosaraju's algorithm for computing strongly connected components in O(n) time
*/
package scc

import (
	"fmt"
	"sort"
	"util"
)

type Edge struct {
	v1 int
	v2 int
}

type Vertex struct {
	label int
	seen  bool
	edges []*Edge
}

type Vertices []*Vertex

type Graph struct {
	E []*Edge
	V Vertices
}

type Kosaraju struct {
	t           int
	finishTimes map[int]int
	path        []int
	scc         map[int][]int
}

func NewKosaraju() *Kosaraju {

	// create a Kosaraju
	k := new(Kosaraju)
	k.t = 1
	k.finishTimes = make(map[int]int)
	k.path = make([]int, 0)
	k.scc = make(map[int][]int)
	return k
}

func NewGraph(t [][]int) *Graph {
	g := new(Graph)
	g.E = make([]*Edge, 0)
	g.V = make([]*Vertex, 0)

	lookup := make(map[int]*Vertex)
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			lookup[t[i][j]] = nil
		}
	}

	for k, _ := range lookup {
		v := new(Vertex)
		v.label = k
		v.edges = make([]*Edge, 0)
		g.V = append(g.V, v)
	}

	g.sortVertices()
	for i := 0; i < len(t); i++ {
		// find index of v1
		v1 := g.findVertexIndex(t[i][0])
		if v1 < 0 {
			fmt.Printf("Vertex not found: %d\n", t[i][0])
			continue
		}
		for j := 1; j < len(t[i]); j++ {
			v2 := g.findVertexIndex(t[i][j])
			if v2 < 0 {
				fmt.Printf("Vertex not found: %d\n", t[i][j])
				continue
			}
			e := new(Edge)
			e.v1 = v1
			e.v2 = v2
			g.V[v1].edges = append(g.V[v1].edges, e)
			g.E = append(g.E, e)
		}
	}
	return g
}

func (g *Graph) ReverseGraph() {

	edges := make([]*Edge, 0)

	for _, e := range g.E {
		v1 := e.v1
		v2 := e.v2
		ne := new(Edge)
		ne.v1 = v2
		ne.v2 = v1
		// remove from v1
		g.V[v1].removeEdge(e)
		// add to v2
		g.V[v2].addEdge(ne)
		edges = append(edges, ne)
	}
	g.E = edges
}

func (g *Graph) sortVertices() {
	sort.Sort(g.V)
}

func (g *Graph) ReverseSortedIndices() []int {
	rmap := make(map[int]int, 0)
	labels := make([]int, 0)
	for i, v := range g.V {
		rmap[v.label] = i
		labels = append(labels, v.label)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(labels)))
	rsi := make([]int, 0)
	for _, label := range labels {
		rsi = append(rsi, rmap[label])
	}
	return rsi
}

func (g *Graph) findVertexIndex(label int) int {
	i := sort.Search(len(g.V),
		func(i int) bool {
			return g.V[i].label >= label
		})
	if i < len(g.V) && g.V[i].label == label {
		return i
	}
	return -1
}

func (g *Graph) MarkUnseen() {
	for _, v := range g.V {
		v.seen = false
	}
}

func (g *Graph) Print() {

	fmt.Println("********************************************************************************")
	for i := 0; i < len(g.V); i++ {
		fmt.Printf("V[%d]: %d\n", i, g.V[i].label)
		for j := 0; j < len(g.V[i].edges); j++ {
			v1 := g.V[i].edges[j].v1
			v2 := g.V[i].edges[j].v2
			fmt.Printf("\tE[%d].v1[%d]: %d\n", j, v1, g.V[v1].label)
			fmt.Printf("\tE[%d].v2[%d]: %d\n", j, v2, g.V[v2].label)
		}
	}
	fmt.Println("E:")
	for _, e := range g.E {
		fmt.Printf(" [%d][%d]{%d, %d}, ", e.v1, e.v2, g.V[e.v1].label, g.V[e.v2].label)
	}
	fmt.Println()
	fmt.Println("********************************************************************************")
}

func (v *Vertex) addEdge(edge *Edge) {
	v.edges = append(v.edges, edge)
}

func (v *Vertex) removeEdge(edge *Edge) {
	edges := make([]*Edge, 0)
	found := false
	for _, e := range v.edges {
		if !found && e.v1 == edge.v1 && e.v2 == edge.v2 {
			found = true
			continue
		}
		edges = append(edges, e)
	}
	v.edges = edges
}

func (v *Vertex) Print() {
	fmt.Printf("\nVertex [%d]:\n", v.label)
	for _, e := range v.edges {
		e.Print()
	}
}

func (slice Vertices) Len() int {
	return len(slice)
}

func (slice Vertices) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice Vertices) Less(i, j int) bool {
	return slice[i].label < slice[j].label
}

func (e *Edge) Print() {
	fmt.Printf(" {%d, %d}", e.v1, e.v2)
}

func (k *Kosaraju) Dfs(g *Graph, src int) int {
	found := -1

	// Check i is valid
	if len(g.V) <= src {
		return -1
	}

	v := g.V[src]

	if v.seen {
		return -1
	} else {
		v.seen = true
		k.path = append(k.path, v.label)
		for _, e := range v.edges {
			found = k.Dfs(g, e.v2)
		}
		if k.t > 0 {
			k.finishTimes[v.label] = k.t
			k.t++
		}
	}

	return found
}

func (k *Kosaraju) TopSccs(filename string) []int {
	table := util.ReadIntsTable(filename)
	sccs := make([]int, 0)

	g := NewGraph(table)
	g.ReverseGraph()

	// set finishing times on
	k.t = 1
	rIndices := g.ReverseSortedIndices()
	for _, v := range rIndices {
		k.Dfs(g, v)
	}

	for _, v := range rIndices {
		label := g.V[v].label
		g.V[v].label = k.finishTimes[label]
	}
	g.ReverseGraph()
	// set finishing times off
	k.t = 0
	g.MarkUnseen()
	k.path = nil
	rIndices = g.ReverseSortedIndices()
	for _, v := range rIndices {
		k.Dfs(g, v)

		if len(k.path) > 0 {
			sccs = append(sccs, len(k.path))
			k.path = nil
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sccs)))
	return sccs
}
