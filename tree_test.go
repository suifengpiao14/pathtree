package tree

import (
	"fmt"
	"testing"
)

func GetTree() Tree {
	return NewTree("attribute_tree", nil)
}

func TestTreeTreeSQL(t *testing.T) {
	sql := GetTree().TableSQL()
	fmt.Println(sql)
}
