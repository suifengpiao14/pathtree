package area

import "gitea.programmerfamily.com/go/treeentity"

type Node struct {
	AreaID   string `json:"areaId"`
	AreaName string `json:"areaName"`
	ParentID string `json:"parentId"`
}

type Nodes []Node

func (node *Node) GetNodeID() (nodeID string) {
	return nodeID
}
func (node *Node) SetPath(path string) {

}
func (node *Node) GetPath() (path string) {
	return path
}
func (node *Node) SetDepth(depth int) {

}
func (node *Node) GetDepth() (depth int) {
	return depth
}
func (node *Node) IsLeaf() (yes bool) {
	return yes
}
func (node *Node) SetParentID(parentId string) {

}
func (node *Node) GetParent() (parent Node) {
	return parent
}
func (node *Node) GetRepository() (r treeentity.Repository) {
	return r
}
