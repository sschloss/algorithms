/*
	mincuts
	Kargar's min cut algorithm
*/
package mincut

import (
	"fmt"
	"math/rand"
)

type Edge struct {
	v1 string
	v2 string
}

type Vertex struct {
	label string
	edges []*Edge
}

type Graph struct {
	V map[string]*Vertex
	E []*Edge
}

func NewGraph(t [][]string) *Graph {
	g := new(Graph)

	// initialize vertices
	g.V = make(map[string]*Vertex)
	for i := 0; i < len(t); i++ {
		if t[i][0] == "" {
			panic(fmt.Sprintf("Invalid table::[%d][%d]\n", i, 0))
		}
		g.V[t[i][0]] = new(Vertex)
		g.V[t[i][0]].label = string(t[i][0])
		g.V[t[i][0]].edges = make([]*Edge, 0)
	}

	// initialize edges
	g.E = make([]*Edge, 0)
	for i := 0; i < len(t); i++ {
		for j := 1; j < len(t[i]); j++ {
			e := new(Edge)
			e.v1 = t[i][0]
			if t[i][j] == "" {
				panic(fmt.Sprintf("Invalid table::[%d][%d]\n", i, j))
			}
			e.v2 = t[i][j]
			g.E = append(g.E, e)
			g.V[t[i][0]].edges = append(g.V[t[i][0]].edges, e)
		}
	}
	return g
}

func (g *Graph) RandomizedContraction() int {
	i := 0
	for len(g.V) > 2 {
		i = rand.Intn(len(g.E))
		g.Contract(i)
	}

	return int(len(g.E) / 2)
}

func (g *Graph) Contract(index int) int {

	if index >= len(g.E) {
		return 1
	}
	se := g.E[index]
	v1 := se.v1
	v2 := se.v2

	result := g.transferVertexEdges(v1, v2)
	if result == 0 {
		g.adjustVertexEdges(v1, v2)
		g.removeSelfLoops(v1)
		g.updateEdges(v1, v2)
		g.removeVertex(v2)
	}
	return 0
}

func (g *Graph) transferVertexEdges(v1 string, v2 string) int {
	vertex1 := g.getVertex(v1)
	vertex2 := g.getVertex(v2)
	if vertex1 == nil || vertex2 == nil {
		return 1
	}
	// move edges from v2 to v1
	edges := make([]*Edge, 0)
	for i := 0; i < len(g.V[v2].edges); i++ {
		if g.V[v2].edges[i].v2 == v1 {
			continue
		}
		e := new(Edge)
		e.v1 = v1
		e.v2 = g.V[v2].edges[i].v2
		edges = append(edges, e)
	}
	g.V[v1].edges = append(g.V[v1].edges, edges...)
	return 0
}

func (g *Graph) adjustVertexEdges(v1 string, v2 string) {
	// adjust v2 on the edges on each vertex
	for k, _ := range g.V {
		edges := make([]*Edge, 0)
		for i := 0; i < len(g.V[k].edges); i++ {
			e := new(Edge)
			if g.V[k].edges[i].v2 == v2 {
				e.v1 = g.V[k].edges[i].v1
				e.v2 = v1
			} else {
				e.v1 = g.V[k].edges[i].v1
				e.v2 = g.V[k].edges[i].v2
			}
			edges = append(edges, e)
		}
		g.V[k].edges = edges
	}
}

func (g *Graph) removeSelfLoops(v1 string) {
	// remove self loops
	edges := make([]*Edge, 0)
	for i := 0; i < len(g.V[v1].edges); i++ {
		if g.V[v1].edges[i].v2 == v1 {
			continue
		}
		e := new(Edge)
		e.v1 = g.V[v1].edges[i].v1
		e.v2 = g.V[v1].edges[i].v2
		edges = append(edges, e)
	}
	g.V[v1].edges = edges
}

func (g *Graph) updateEdges(v1 string, v2 string) {
	// update graph edges array
	edges1 := make([]*Edge, 0)
	for i := 0; i < len(g.E); i++ {
		e := new(Edge)
		if g.E[i].v1 == v2 {
			e.v1 = v1
			e.v2 = g.E[i].v2
		} else if g.E[i].v2 == v2 {
			e.v1 = g.E[i].v1
			e.v2 = v1
		} else {
			e.v1 = g.E[i].v1
			e.v2 = g.E[i].v2
		}
		edges1 = append(edges1, e)
	}
	g.E = edges1

	// remove self loops in edges
	edges2 := make([]*Edge, 0)
	for _, e := range g.E {
		if e.v1 != e.v2 {
			ne := new(Edge)
			ne.v1 = e.v1
			ne.v2 = e.v2
			edges2 = append(edges2, ne)
		}
	}
	g.E = edges2

}

func (g *Graph) removeVertex(v string) {
	// remove vertex
	delete(g.V, v)
}

func (g *Graph) getVertex(key string) *Vertex {
	v, ok := g.V[key]
	if !ok {
		return nil
	}
	return v
}

func (g *Graph) getEdge(i int) *Edge {
	return g.E[i]
}

func (g *Graph) Print() {

	fmt.Println("********************************************************************************")
	for _, value := range g.V {
		value.Print()
	}
	fmt.Println("All edges:")
	for _, e := range g.E {
		e.Print()
	}
	fmt.Println()
	fmt.Println("********************************************************************************")
}

func (v *Vertex) Print() {
	fmt.Printf("Vertex: %s\n", v.label)
	for _, e := range v.edges {
		e.Print()
	}
	fmt.Println()
}

func (e *Edge) Print() {
	fmt.Printf(" {%s, %s}", e.v1, e.v2)
}
