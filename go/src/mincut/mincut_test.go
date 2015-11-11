/*
	comparisons
	test for comparing pivot choices for the quicksort algorithm
*/
package mincut

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"mincut"
	"os"
	"testing"
	"util"
)

func TestMain(m *testing.M) {
	flag.Parse()
	status := m.Run()
	os.Exit(status)
}

func TestReduced(t *testing.T) {
	table := util.ReadTable("../../../data/reducedMinCut.txt")
	i, cut, mcut := 0, 0, 0
	for i < 100 {
		g := mincut.NewGraph(table)
		cut = g.RandomizedContraction()
		if i == 0 || cut < mcut {
			mcut = cut
		}
		i++
	}
	fmt.Printf("Min cut: %d\n", mcut)
	assert.Equal(t, mcut, 2, "Minimum cut must equal 2 for the given graph")
}
func TestFull(t *testing.T) {
	table := util.ReadTable("../../../data/kargerMinCut.txt")
	i, cut, mcut := 0, 0, 0
	for i < 100 {
		g := mincut.NewGraph(table)
		cut = g.RandomizedContraction()
		if i == 0 || cut < mcut {
			mcut = cut
		}
		i++
	}
	fmt.Printf("Min cut: %d\n", mcut)
	assert.Equal(t, mcut, 17, "Minimum cut must equal 17 for the given graph")
}

func TestTable(t *testing.T) {
	table := util.ReadTable("../../../data/kargerMinCut.txt")
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			assert.NotEmpty(t, table[i][j], fmt.Sprintf("Location[%d][%d]"))
		}
	}
}
