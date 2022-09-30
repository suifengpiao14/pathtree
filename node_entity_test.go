package treeentity_test

import (
	"fmt"
	"testing"

	"gitea.programmerfamily.com/go/treeentity"
)

func TestGetNode(t *testing.T) {
	var repository treeentity.RepositoryInterface
	instance := treeentity.NewNodeEntity(repository)
	nodeId := "first"
	label := "test"
	parentId := ""
	out, err := instance.AddNode(nodeId, parentId, label)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func TestBuildTree(t *testing.T) {
	// jsonStr := `[
	// {"nodeId":"B","parentId":"A","name":"first"},
	// {"nodeId":"C","parentId":"B","name":"first-chirden"},
	// {"nodeId":"A","parentId":"","name":"root"},
	// {"nodeId":"D","parentId":"A","name":"second"}
	// ]`
	jsonStr2 := `[
	{"nodeId":4,"parentId":3,"name":"first"},
	{"nodeId":5,"parentId":4,"name":"first-chirden"},
	{"nodeId":1,"parentId":0,"name":"root"},
	{"nodeId":3,"parentId":1,"name":"second"}
	]`
	out, err := treeentity.BuildTree(jsonStr2, "nodeId", "parentId")
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
