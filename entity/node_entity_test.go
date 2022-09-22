package entity_test

import (
	"fmt"
	"testing"

	"gitea.programmerfamily.com/go/treemodel/entity"
)

func TestGetNode(t *testing.T) {
	var repository entity.RepositoryInterface
	instance := entity.NewNodeEntity(repository)
	nodeId := "first"
	label := "test"
	parentId := ""
	var out entity.AddNodeOut
	err := instance.AddNode(nodeId, label, parentId).Do().Out(&out).Error()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
