package treeentity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Repository interface {
	AddNode(node Node) (err error)
	UpdateNode(node Node) (err error)
	UpdateBatchNode(nodes []Node) (err error)
	GetAllByPathPrefix(pathPrefix string, depth int, nodes *[]Node) (err error)
	GetAllNodeByNodeIds(nodeIds []string, nodes *[]Node) (err error)
}

type Node interface {
	GetNodeID() (nodeID string)
	SetPath(path string)
	GetPath() (path string)
	SetDepth(depth int)
	GetDepth() (depth int)
	IsLeaf() (yes bool)
	SetParentID(parentId string)
	GetParent() (parent Node)
	GetRepository() (r Repository)
}
type Nodes []Node

// Add 节点
func Add(n Node) (err error) {
	parent := n.GetParent()
	if parent != nil && !parent.IsLeaf() {
		err = errors.Errorf("%s;nodeId:%s", ERROR_ADD_NODE_LABLE_LEAF, parent.GetNodeID())
		return err
	}
	path := fmt.Sprintf("/%s", n.GetNodeID())
	n.SetPath(path)
	var diffDepth int
	path, diffDepth = calPathAndDepth(n)
	depth := diffDepth + n.GetDepth()
	n.SetPath(path)
	n.SetDepth(depth)
	err = n.GetRepository().AddNode(n)
	return err
}

// GetAllParent 获取节点的所有父节点
func GetAllParent(n Node, withOutSelf bool, out *[]Node) (err error) {
	r := n.GetRepository()
	nodeIdList := strings.Split(n.GetPath(), "/")
	if len(nodeIdList) == 0 {
		return nil
	}
	if !withOutSelf {
		nodeIdList = nodeIdList[:len(nodeIdList)-1]
	}
	if len(nodeIdList) == 0 {
		return nil
	}
	err = r.GetAllNodeByNodeIds(nodeIdList, out)
	if err != nil {
		return err
	}
	return nil
}

func GetSubTree(n Node, depth int, withOutSelf bool, out *[]Node) (err error) {
	r := n.GetRepository()
	p := n.GetParent()
	maxDepth := DEPTH_MAX
	if depth > 0 {
		maxDepth = p.GetDepth() + depth
	}
	parentPath := p.GetPath()
	if !withOutSelf {
		parentPath = fmt.Sprintf("%s/", parentPath)
	}
	err = r.GetAllByPathPrefix(parentPath, maxDepth, out)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

func MoveSubTree(node Node, newParentId string) (err error) {
	r := node.GetRepository()
	nodeOldPath := node.GetPath()
	nodeNewPath, diffDepth := calPathAndDepth(node)
	newDepth := diffDepth + node.GetDepth()
	// 修改node 节点本身
	node.SetParentID(newParentId)
	node.SetPath(nodeNewPath)
	node.SetDepth(newDepth)
	err = r.UpdateNode(node)
	if err != nil {
		return err
	}
	// 获取所有子节点
	var childrenNodeList []Node
	err = r.GetAllByPathPrefix(nodeOldPath, -1, &childrenNodeList)
	if err != nil {
		return err
	}
	// 更新子节点路径和深度值
	newChildren := make([]Node, 0)
	for _, children := range childrenNodeList {
		newPath := strings.Replace(children.GetPath(), nodeOldPath, nodeNewPath, 1)
		newDepth := children.GetDepth() + diffDepth
		children.SetPath(newPath)
		children.SetDepth(newDepth)
		newChildren = append(newChildren, children)
	}
	err = r.UpdateBatchNode(newChildren)
	return err
}

func DeleteTree(node Node) (nodeIdList []string, err error) {
	r := node.GetRepository()
	// 获取所有子节点
	var childrenNodeList []Node
	err = r.GetAllByPathPrefix(node.GetPath(), -1, &childrenNodeList)
	if err != nil {
		return nil, err
	}
	nodeIdList = make([]string, 0)
	for _, childern := range childrenNodeList {
		nodeIdList = append(nodeIdList, childern.GetNodeID())
	}

	return nodeIdList, nil
}

// calPath 计算节点迁移的新路径和深度
func calPathAndDepth(n Node) (newPath string, diffDepth int) {
	parent := n.GetParent()
	// if parent == nil {
	// 	return n.GetPath(), 0
	// }
	newPath = fmt.Sprintf("%s%s", parent.GetPath(), n.GetPath())
	diffDepth = parent.GetDepth() - n.GetDepth() + 1
	return newPath, diffDepth
}
