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
