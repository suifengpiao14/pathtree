package treeentity

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var DEPTH_MIN = -1     //最小深度值
var DEPTH_MAX = 100000 //最大深度值

type TreeRepository interface {
	AddNode(node TreeNode) (err error)
	UpdateNode(node TreeNode) (err error)
	UpdateBatchNode(nodes []TreeNode) (err error)
	GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error)
	GetAllNodeByNodeIds(nodeIds []string, nodes interface{}) (err error)
}

type TreeNode interface {
	GetNodeID() (nodeID string)
	SetPath(path string)
	GetPath() (path string)
	SetDepth(depth int)
	GetDepth() (depth int)
	SetParentID(parentId string)
	GetParent() (parent TreeNode, err error)
}
type TreeNodes []TreeNode

type tree struct {
	node       TreeNode
	repository TreeRepository
}

func NewTree(node TreeNode, repository TreeRepository) (t *tree) {
	return &tree{
		node:       node,
		repository: repository,
	}
}

// AddNode 节点
func (t tree) AddNode() (err error) {

	n := t
	path := fmt.Sprintf("/%s", n.node.GetNodeID())
	n.node.SetPath(path)
	var diffDepth int
	path, diffDepth, err = t.calPathAndDepth()
	if err != nil {
		return err
	}
	depth := diffDepth + n.node.GetDepth()
	n.node.SetPath(path)
	n.node.SetDepth(depth)
	err = n.repository.AddNode(n.node)
	return err
}

func (t tree) MoveChildren(newParentId string) (err error) {
	node := t
	r := node.repository
	nodeOldPath := node.node.GetPath()
	nodeNewPath, diffDepth, err := t.calPathAndDepth()
	if err != nil {
		return err
	}
	newDepth := diffDepth + node.node.GetDepth()
	// 修改node 节点本身
	node.node.SetParentID(newParentId)
	node.node.SetPath(nodeNewPath)
	node.node.SetDepth(newDepth)
	err = r.UpdateNode(node.node)
	if err != nil {
		return err
	}
	// 获取所有子节点
	var childrenNodeList []TreeNode
	err = r.GetAllByPathPrefix(nodeOldPath, -1, &childrenNodeList)
	if err != nil {
		return err
	}
	// 更新子节点路径和深度值
	newChildren := make([]TreeNode, 0)
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

func (t tree) DeleteWithChildren() (nodeIdList []string, err error) {
	node := t
	r := node.repository
	// 获取所有子节点
	var childrenNodeList []TreeNode
	err = r.GetAllByPathPrefix(node.node.GetPath(), -1, &childrenNodeList)
	if err != nil {
		return nil, err
	}
	nodeIdList = make([]string, 0)
	for _, childern := range childrenNodeList {
		nodeIdList = append(nodeIdList, childern.GetNodeID())
	}

	return nodeIdList, nil
}

// GetParents 获取节点的所有父节点
func (t tree) GetParents(relativeDepth int, withOutSelf bool, out interface{}) (err error) {
	n := t
	r := n.repository
	nodeIdList := strings.Split(n.node.GetPath(), "/")
	if len(nodeIdList) == 0 {
		return nil
	}
	if !withOutSelf {
		nodeIdList = nodeIdList[:len(nodeIdList)-1]
	}
	if len(nodeIdList) == 0 {
		return nil
	}
	nodes := make([]TreeNode, 0)
	err = r.GetAllNodeByNodeIds(nodeIdList, &nodes)
	if err != nil {
		return err
	}
	minDepth := DEPTH_MIN
	if relativeDepth > 0 {
		minDepth = n.node.GetDepth() - relativeDepth
	}
	outNodes := make([]TreeNode, 0)
	for _, node := range nodes {
		if node.GetDepth() <= minDepth {
			continue
		}
		outNodes = append(outNodes, node)
	}
	rv := reflect.Indirect(reflect.ValueOf(out))
	if !rv.CanSet() {
		err = errors.Errorf("return variable out must be canSet")
		return err
	}
	rv.Set(reflect.Indirect(reflect.ValueOf(outNodes)))
	return nil
}

func (t tree) GetChildren(relativeDepth int, withOutSelf bool, out interface{}) (err error) {
	n := t.node
	r := t.repository
	maxDepth := DEPTH_MAX
	if relativeDepth > 0 {
		maxDepth = n.GetDepth() + relativeDepth
	}
	parentPath := n.GetPath()
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

// calPath 计算节点迁移的新路径和深度
func (t tree) calPathAndDepth() (newPath string, diffDepth int, err error) {
	n := t
	parent, err := n.node.GetParent()
	if err != nil {
		return newPath, diffDepth, err
	}
	newPath = fmt.Sprintf("%s%s", parent.GetPath(), n.node.GetPath())
	diffDepth = parent.GetDepth() - n.node.GetDepth() + 1
	return newPath, diffDepth, nil
}
