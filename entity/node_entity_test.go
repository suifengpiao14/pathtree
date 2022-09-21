package entity_test

import (
	"testing"

	"gitea.programmerfamily.com/go/treemodel/entity"
)

func TestGetNode(t *testing.T) {
	var repository entity.RepositoryInterface
	instance := entity.NewNodeEntity(repository)
	data := `
	{"nodeId":"first"}
	
	`
	err := instance.AddNode([]byte(data)).Do().Out(nil).Error()
	if err != nil {
		panic(err)
	}
}
