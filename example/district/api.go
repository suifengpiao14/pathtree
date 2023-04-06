package district

import (
	"gitea.programmerfamily.com/go/pathtree"
)

func Add(record District) (err error) {
	nodeEntity := pathtree.NewTreeNode(&record)
	err = nodeEntity.AddNode()
	if err != nil {
		return err
	}
	return nil
}
func GetByCodeWithChildren(code string) (out []*District, err error) {
	record := District{
		Code: code,
	}
	nodeEntity := pathtree.NewTreeNode(&record)
	err = nodeEntity.GetChildren(-1, true, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func GetParent(code string) (out []*District, err error) {
	record := District{
		Code: code,
	}
	nodeEntity := pathtree.NewTreeNode(&record)
	err = nodeEntity.GetParents(-1, false, &out)
	if err != nil {
		return out, err
	}
	return out, nil
}

func MoveNode(code string, newParentCode string) (err error) {
	record := District{
		Code: code,
	}
	nodeEntity := pathtree.NewTreeNode(&record)
	out := make([]District, 0)
	err = nodeEntity.MoveChildren(newParentCode, out)
	if err != nil {
		return err
	}
	return nil
}
