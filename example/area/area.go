package area

import (
	"strconv"

	"gitea.programmerfamily.com/go/treeentity"
)

type AreaRecordI interface {
	treeentity.TreeRepository
	GetByAreaID(areaID string) (areaRecord *AreaRecord, err error)
	GetByLevel(depth string) (areaRecord AreaRecords, err error)
	GetByKeyWord(keyword string, depth string) (areaRecord AreaRecords, err error)
}

type AreaRecord struct {
	AreaID   string `json:"areaId"`
	AreaName string `json:"areaName"`
	ParentID string `json:"parentId"`
	Path     string `json:"path"`
	Depth    string `json:"depth"`
}

const (
	LEVEL_CITY = "3"
)

type AreaRecords []AreaRecord

func (node *AreaRecord) GetNodeID() (nodeID string) {
	nodeID = node.AreaID
	return nodeID
}
func (node *AreaRecord) SetPath(path string) {
	node.Path = path

}
func (node *AreaRecord) GetPath() (path string) {
	return node.Path
}
func (node *AreaRecord) SetDepth(depth int) {
	node.Depth = strconv.Itoa(depth)

}
func (node *AreaRecord) GetDepth() (depth int) {
	depth, _ = strconv.Atoi(node.Depth)
	return depth
}

func (node *AreaRecord) SetParentID(parentId string) {
	node.ParentID = parentId

}
func (node *AreaRecord) GetParent() (parent treeentity.TreeNode, err error) {
	parentArea, err := node.GetRepository().GetByAreaID(node.ParentID)
	if err != nil {
		return nil, err
	}
	return parentArea, nil
}

func (node *AreaRecord) GetRepository() (r AreaRecordI) {
	return r
}
