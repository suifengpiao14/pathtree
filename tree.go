package pathtree

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var DEPTH_MIN = -1     //最小深度值
var DEPTH_MAX = 100000 //最大深度值

var (
	ERROR_NODE_NOT_FOUND = errors.Errorf("node not found")
)

type TreeRepositoryI interface {
	AddNode(node TreeNodeI) (err error)
	UpdateNode(node TreeNodeI) (err error)
	UpdateBatchNode(nodes TreeNodeIs) (err error)
	GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error)
	GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error)
}

type EmptyTreeRpository struct{}

func (r *EmptyTreeRpository) AddNode(node TreeNodeI) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "AddNode")
	panic(err)
}
func (r *EmptyTreeRpository) UpdateNode(node TreeNodeI) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "UpdateNode")
	panic(err)
}
func (r *EmptyTreeRpository) UpdateBatchNode(nodes TreeNodeIs) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "UpdateBatchNode")
	panic(err)
}
func (r *EmptyTreeRpository) GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetAllByPathPrefix")
	panic(err)
}
func (r *EmptyTreeRpository) GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetAllByNodeIds")
	panic(err)
}

type TreeNodeI interface {
	GetNodeID() (nodeID string)
	SetPath(path string)
	SetDepth(depth int)
	GetParent() (parent TreeNodeI, err error)
	GetPath() (path string)
	GetDepth() (depth int)
	SetParentID(parentId string)
	IsRoot() (ok bool)
	AddChildren(node TreeNodeI)
	IncrChildrenCount(causeNode TreeNodeI)
}

//EmptyTreeNode 主要用于屏蔽不需要实现的接口，以及对已有的实现屏蔽新增方法，建议TreeNodeI 的实现继承EmptyTreeNode
type EmptyTreeNode struct {
}

var ERROR_NOT_IMPLEMENTED = errors.New("not implemented")

func (etn *EmptyTreeNode) GetNodeID() (nodeID string) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetNodeID")
	panic(err)
}
func (etn *EmptyTreeNode) SetPath(path string) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "SetPath")
	panic(err)
}
func (etn *EmptyTreeNode) SetDepth(depth int) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "SetDepth")
	panic(err)
}
func (etn *EmptyTreeNode) GetParent() (parent TreeNodeI, err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetParent")
	panic(err)
}
func (etn *EmptyTreeNode) GetPath() (path string) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetPath")
	panic(err)
}
func (etn *EmptyTreeNode) GetDepth() (depth int) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetDepth")
	panic(err)
}
func (etn *EmptyTreeNode) SetParentID(parentId string) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "SetParentID")
	panic(err)
}
func (etn *EmptyTreeNode) IsRoot() (ok bool) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "IsRoot")
	panic(err)
}
func (etn *EmptyTreeNode) AddChildren(node TreeNodeI) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "AddChildren")
	panic(err)
}
func (etn *EmptyTreeNode) IncrChildrenCount(causeNode TreeNodeI) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "IncrChildrenCount")
	panic(err)
}

type TreeNodeIs []TreeNodeI

//ConvertToTreeNodes 具体实例数据转接口数据
func ConvertToTreeNodes(src interface{}) (treeNodes TreeNodeIs) {
	srcRv := reflect.Indirect(reflect.ValueOf(src))
	l := srcRv.Len()
	treeNodes = make(TreeNodeIs, l)
	dstRv := reflect.Indirect(reflect.ValueOf(treeNodes))
	for i := 0; i < l; i++ {
		rv := srcRv.Index(i)
		dstRv.Index(i).Set(rv)
	}
	return treeNodes
}

func (tns TreeNodeIs) Convert(dst interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(dst))
	copy := rv
	c := len(tns)
	arr := make([]reflect.Value, c)
	for _, node := range tns {
		nodeRv := reflect.ValueOf(node)
		arr = append(arr, nodeRv)
	}
	copy = reflect.Append(copy, arr...) // 此处重新赋值，丢失引用,所以用copy
	rv.Set(copy)
}

func (tns TreeNodeIs) ResetAllPath() (err error) {
	count := len(tns)
	// 循环处理数据,增加path和depth
	for i := 0; i < count; i++ {
		node := tns[i]
		revNodeIdList := make([]string, 0)
		parent := node
		for {
			parentId := ""
			isRoot := true
			if parent != nil {
				parentId = parent.GetNodeID()
				isRoot = parent.IsRoot()
			}
			revNodeIdList = append(revNodeIdList, parentId)
			if isRoot {
				break
			}

			parent, err = parent.GetParent()
			if errors.Is(err, ERROR_NODE_NOT_FOUND) {
				err = nil
			}
			if err != nil {
				return err
			}
		}
		var w bytes.Buffer
		l := len(revNodeIdList)
		for i := l - 1; i > -1; i-- {
			nodeId := revNodeIdList[i]
			if nodeId == "" {
				continue
			}
			w.WriteString("/")
			w.WriteString(nodeId)
		}
		path := w.String()
		depth := strings.Count(path, "/")
		node.SetDepth(depth)
		node.SetPath(path)
	}
	return nil
}
func (tns *TreeNodeIs) FormatToTree() (trees TreeNodeIs) {
	trees = make(TreeNodeIs, 0)
	for _, node := range *tns {
		if node.IsRoot() {
			trees = append(trees, node)
			continue
		}
		parent, _ := node.GetParent()
		if parent != nil {
			parent.AddChildren(node)
		}
	}
	return trees
}

func (tns *TreeNodeIs) CountChildren() {
	for _, node := range *tns {
		parent, _ := node.GetParent()
		if parent != nil {
			parent.IncrChildrenCount(node)
		}
	}
}

type tree struct {
	nodeI      TreeNodeI
	repository TreeRepositoryI
}

func NewTree(nodeI TreeNodeI, repository TreeRepositoryI) (t *tree) {
	return &tree{
		nodeI:      nodeI,
		repository: repository,
	}
}

// AddNode 节点
func (t tree) AddNode() (err error) {

	n := t
	path := fmt.Sprintf("/%s", n.nodeI.GetNodeID())
	n.nodeI.SetPath(path)
	var diffDepth int
	path, diffDepth, err = t.calPathAndDepth()
	if err != nil {
		return err
	}
	depth := diffDepth + n.nodeI.GetDepth()
	n.nodeI.SetPath(path)
	n.nodeI.SetDepth(depth)
	err = n.repository.AddNode(n.nodeI)
	return err
}

func (t tree) MoveChildren(newParentId string) (err error) {
	node := t
	r := node.repository
	nodeOldPath := node.nodeI.GetPath()
	nodeNewPath, diffDepth, err := t.calPathAndDepth()
	if err != nil {
		return err
	}
	newDepth := diffDepth + node.nodeI.GetDepth()
	// 修改node 节点本身
	node.nodeI.SetParentID(newParentId)
	node.nodeI.SetPath(nodeNewPath)
	node.nodeI.SetDepth(newDepth)
	err = r.UpdateNode(node.nodeI)
	if err != nil {
		return err
	}
	// 获取所有子节点
	var childrenNodeList TreeNodeIs
	err = r.GetAllByPathPrefix(nodeOldPath, -1, &childrenNodeList)
	if err != nil {
		return err
	}
	// 更新子节点路径和深度值
	newChildren := make(TreeNodeIs, 0)
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
	var childrenNodeList TreeNodeIs
	err = r.GetAllByPathPrefix(node.nodeI.GetPath(), -1, &childrenNodeList)
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
	nodeIdList := strings.Split(n.nodeI.GetPath(), "/")
	if len(nodeIdList) == 0 {
		return nil
	}
	if !withOutSelf {
		nodeIdList = nodeIdList[:len(nodeIdList)-1]
	}
	if len(nodeIdList) == 0 {
		return nil
	}
	nodes := make(TreeNodeIs, 0)
	err = r.GetAllByNodeIds(nodeIdList, &nodes)
	if err != nil {
		return err
	}
	minDepth := DEPTH_MIN
	if relativeDepth > 0 {
		minDepth = n.nodeI.GetDepth() - relativeDepth
	}
	outNodes := make(TreeNodeIs, 0)
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
	n := t.nodeI
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
	parent, err := n.nodeI.GetParent()
	if err != nil {
		return newPath, diffDepth, err
	}
	newPath = fmt.Sprintf("%s%s", parent.GetPath(), n.nodeI.GetPath())
	diffDepth = parent.GetDepth() - n.nodeI.GetDepth() + 1
	return newPath, diffDepth, nil
}
