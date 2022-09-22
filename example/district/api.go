package district

import (
	"gitea.programmerfamily.com/go/treeentity"
	"gitea.programmerfamily.com/go/treeentity/example/district/repository"
)

func Add(record DistrictKV) (err error) {
	r := repository.NewMemeryRepository()
	nodeEntity := treeentity.NewNodeEntity(r)
	addData, err := nodeEntity.AddNode(record.NodeID, record.ParentID, record.Label)
	if err != nil {
		return err
	}
	record.Depth = addData.Depth
	record.Path = addData.Path
	err = r.AddNode(record)
	if err != nil {
		return err
	}
	return nil
}
