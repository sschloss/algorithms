/*
	scc_test
	Kosaraju's algorithm for computing strongly connected components in O(n) time
*/
package scc

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"scc"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	status := m.Run()
	os.Exit(status)
}

func TestSCC(tst *testing.T) {
	fmt.Println("TestSCC")
	filename := "../../../data/SCC.txt"
	k := scc.NewKosaraju()
	expected := []int{434821, 968, 459, 313, 211}
	result := k.TopSccs(filename)
	if len(expected) < len(result) {
		for i := 0; i < len(expected); i++ {
			msg := fmt.Sprintf("Expected %d\nGot %d\n", expected[i], result[i])
			assert.Equal(tst, expected[i], result[i], msg)
		}
	} else {
		tst.Fail()
	}

}
