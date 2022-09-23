package district

import (
	"encoding/json"

	"gitea.programmerfamily.com/go/treeentity"
	"gitea.programmerfamily.com/go/treeentity/example/district/repository"
)

func Add(record DistrictKV) (err error) {
	r := repository.NewSqlRepository()
	nodeEntity := treeentity.NewNodeEntity(r)
	addData, err := nodeEntity.AddNode(record.NodeID, record.ParentID, record.Label)
	if err != nil {
		return err
	}
	record.Depth = addData.Depth
	record.Path = addData.Path
	b, err := json.Marshal(record)
	if err != nil {
		return err
	}
	err = r.AddNode(b)
	if err != nil {
		return err
	}
	return nil
}
