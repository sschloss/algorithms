/*
	comparisons
	test for comparing pivot choices for the quicksort algorithm
*/
package contraction

import (
	"contraction"
	"flag"
	"fmt"
	"os"
	"testing"
	"util"
)

func TestMain(m *testing.M) {
	flag.Parse()

	fmt.Println("Hello")

	table := util.ReadTable("../../../data/reducedMinCut.txt")
	g := contraction.NewGraph(table)
	//	g.Print()
	g.Contract()
	//	g.Print()
	//	status := m.Run()
	//	os.Exit(status)
	os.Exit(0)
}

func TestFull(t *testing.T) {
	table := util.ReadTable("../../../data/kargerMinCut.txt")

	g := contraction.NewGraph(table)

	//	g.Print()
	for i := 0; i < len(g.V); i++ {
		g.V[i].Print()
	}
	//	e := g.GetEdge(0)
	//	e.Print()

}
