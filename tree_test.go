package tree

import (
	"fmt"
	"testing"
)

func TestTreeSQL(t *testing.T) {
	out := TreeSQL("tree_relation")
	fmt.Println(out)

}
