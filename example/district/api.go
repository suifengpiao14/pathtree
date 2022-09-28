package district

import (
	"encoding/json"

	"gitea.programmerfamily.com/go/treeentity"
	"gitea.programmerfamily.com/go/treeentity/example/district/repository"
)

func Add(record DistrictKV) (err error) {
	r := repository.NewDoaRepository()
	nodeEntity := treeentity.NewNodeEntity(r)
	addData, err := nodeEntity.AddNode(record.Code, record.ParentCode, record.Label)
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
func GetByCodeWithChildren(code string) (out []*DistrictKV, err error) {
	r := repository.NewDoaRepository()
	nodeEntity := treeentity.NewNodeEntity(r)
	err = nodeEntity.GetSubTree(code, -1, true, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func GetParent(code string) (out []*DistrictKV, err error) {
	r := repository.NewDoaRepository()
	nodeEntity := treeentity.NewNodeEntity(r)
	err = nodeEntity.GetAllParent(code, false, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func MoveNode(code string, newParentCode string) (err error) {
	r := repository.NewDoaRepository()
	nodeEntity := treeentity.NewNodeEntity(r)
	moveData, err := nodeEntity.MoveSubTree(code, newParentCode)
	if err != nil {
		return err
	}
	return nil
}
