/*
	Counting inversions
*/
package inversions

import (
	"inversions"
	"testing"
	"util"
)

func TestCountInversions(t *testing.T) {

	ex1 := util.ReadInts("../../../data/IntegerArray.txt")
	count := inversions.CountInversions(ex1)
	if count != 2407905288 {
		t.Error("Fail: %q", count)
	}
}
