package tree

import (
	"fmt"
	"testing"
)

func TestTreeSQL(t *testing.T) {
	out := TreeSQL("tree_relation")
	fmt.Println(out)
}

func GetTree() Tree {
	return NewTree("attribute_tree")
}

func TestTreeTreeSQL(t *testing.T) {
	sql := GetTree().TableSQL()
	fmt.Println(sql)
}
