/*
	contractions
	Kargar's min cut algorithm
*/
package contraction

import (
	"fmt"
)

type Edge struct {
	v1 *Vertex
	v2 *Vertex
}

type Vertex struct {
	label int
	edges []*Edge
}

type Graph struct {
	V []*Vertex
	E []*Edge
}

func NewGraph(t [][]int) *Graph {
	g := new(Graph)

	// initialize vertices
	vertices := make([]*Vertex, len(t))
	for i := 0; i < len(t); i++ {
		vertices[i] = new(Vertex)
		vertices[i].label = t[i][0]
		vertices[i].edges = make([]*Edge, 0)
	}
	g.V = make([]*Vertex, len(vertices))
	g.V = vertices

	// initialize edges
	edges := make([]*Edge, 0)
	for i := 0; i < len(t); i++ {
		for j := 1; j < len(t[i]); j++ {
			e := new(Edge)
			e.v1 = g.V[i]
			e.v2 = g.V[t[i][j]-1]
			edges = append(edges, e)
			g.V[i].edges = append(g.V[i].edges, e)
		}
	}
	g.E = make([]*Edge, len(edges))
	g.E = edges
	return g
}

func (g *Graph) Contract() {

	g.E[0].v2.Print()

	//	e.v1.Print()
	//	for i := 0; i < len(g.E[0].v2.edges); i++ {
	//		g.E[0].v1.edges = append(g.E[0].v2.edges[i])
	//		g.E[0].v2.edges[i]
	//	}

}

func (g *Graph) GetVertex(i int) *Vertex {
	return g.V[i]
}

func (g *Graph) GetEdge(i int) *Edge {
	return g.E[i]
}

func (g *Graph) Print() {
	for _, v := range g.V {
		v.Print()
	}
	fmt.Println("All edges:")
	for _, e := range g.E {
		e.Print()
	}
	fmt.Println()
}

func (v *Vertex) Print() {
	fmt.Printf("Vertex: %d\n", v.label)
	for _, e := range v.edges {
		e.Print()
	}
	fmt.Println()
}

func (e *Edge) Print() {
	fmt.Printf(" {%d, %d}", e.v1.label, e.v2.label)
}
