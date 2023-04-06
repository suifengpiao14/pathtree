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
var PathSpec = "/"

var (
	ERROR_NODE_NOT_FOUND = errors.Errorf("node not found")
)

type TreeNodeI interface {
	GetNodeID() (nodeID string)
	SetPath(path string)
	SetDepth(depth int)
	GetParentID() (parentID string)
	GetParent() (parent TreeNodeI, err error)
	GetPath() (path string)
	GetDepth() (depth int)
	SetParentID(parentId string)
	IsRoot() (ok bool)
	AddChildren(node TreeNodeI)
	IncrChildrenCount(causeNode TreeNodeI)
	GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error)
	GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error)
}

// EmptyTreeNode 主要用于屏蔽不需要实现的接口，以及对已有的实现屏蔽新增方法，建议TreeNodeI 的实现继承EmptyTreeNode
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
func (etn *EmptyTreeNode) GetParentID() (parentID string) {
	err := errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetParentID")
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
func (etn *EmptyTreeNode) GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetAllByPathPrefix")
	panic(err)
}
func (etn *EmptyTreeNode) GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error) {
	err = errors.WithMessage(ERROR_NOT_IMPLEMENTED, "GetAllByNodeIds")
	panic(err)
}

type treeNodeIs []TreeNodeI

// ConvertToTreeNodes 具体实例数据转接口数据
func ConvertToTreeNodes(src interface{}) (treeNodes treeNodeIs) {
	srcRv := reflect.Indirect(reflect.ValueOf(src))
	l := srcRv.Len()
	treeNodes = make(treeNodeIs, l)
	dstRv := reflect.Indirect(reflect.ValueOf(treeNodes))
	for i := 0; i < l; i++ {
		rv := srcRv.Index(i)
		dstRv.Index(i).Set(rv)
	}
	return treeNodes
}

func (tns treeNodeIs) Convert(dst interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(dst))
	copy := rv
	c := len(tns)
	arr := make([]reflect.Value, c)
	for i := 0; i < c; i++ {
		arr[i] = reflect.ValueOf(tns[i])
	}
	copy = reflect.Append(copy, arr...) // 此处重新赋值，丢失引用,所以用copy
	rv.Set(copy)
}

func (tns treeNodeIs) ResetAllPath() (err error) {
	count := len(tns)
	// 循环处理数据,增加path和depth
	for i := 0; i < count; i++ {
		treeNode := NewTreeNode(tns[i])
		err = treeNode.ResetPath()
		if err != nil {
			return err
		}
	}
	return nil
}
func (tns *treeNodeIs) FormatToTree() (trees treeNodeIs) {
	trees = make(treeNodeIs, 0)
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

func (tns *treeNodeIs) CountChildren() {
	for _, node := range *tns {
		parent, _ := node.GetParent()
		if parent != nil {
			parent.IncrChildrenCount(node)
		}
	}
}

type treeNode struct {
	nodeI TreeNodeI
}

func NewTreeNode(nodeI TreeNodeI) (t *treeNode) {
	return &treeNode{
		nodeI: nodeI,
	}
}

// AddNode 设置节点的路径和深度
func (t treeNode) AddNode() (err error) {

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
	return err
}

func (t treeNode) MoveChildren(newParentId string, out interface{}) (err error) {
	node := t
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
	// 获取所有子节点
	var childrenNodeList treeNodeIs
	err = node.nodeI.GetAllByPathPrefix(nodeOldPath, -1, &childrenNodeList)
	if err != nil {
		return err
	}
	// 更新子节点路径和深度值
	newChildren := make(treeNodeIs, 0)
	for _, children := range childrenNodeList {
		newPath := strings.Replace(children.GetPath(), nodeOldPath, nodeNewPath, 1)
		newDepth := children.GetDepth() + diffDepth
		children.SetPath(newPath)
		children.SetDepth(newDepth)
		newChildren = append(newChildren, children)
	}
	nodes := make(treeNodeIs, 0)
	nodes = append(nodes, node.nodeI)
	nodes = append(nodes, newChildren...)
	nodes.Convert(out)
	return nil
}

func (t treeNode) DeleteWithChildren() (nodeIdList []string, err error) {
	node := t
	// 获取所有子节点
	var childrenNodeList treeNodeIs
	err = node.nodeI.GetAllByPathPrefix(node.nodeI.GetPath(), -1, &childrenNodeList)
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
func (t treeNode) GetParents(relativeDepth int, withOutSelf bool, out interface{}) (err error) {
	n := t
	nodeIdList := strings.Split(n.nodeI.GetPath(), PathSpec)
	if len(nodeIdList) == 0 {
		return nil
	}
	if !withOutSelf {
		nodeIdList = nodeIdList[:len(nodeIdList)-1]
	}
	if len(nodeIdList) == 0 {
		return nil
	}
	nodes := make(treeNodeIs, 0)
	err = n.nodeI.GetAllByNodeIds(nodeIdList, &nodes)
	if err != nil {
		return err
	}
	minDepth := DEPTH_MIN
	if relativeDepth > 0 {
		minDepth = n.nodeI.GetDepth() - relativeDepth
	}
	outNodes := make(treeNodeIs, 0)
	for _, node := range nodes {
		if node.GetDepth() <= minDepth {
			continue
		}
		outNodes = append(outNodes, node)
	}
	outNodes.Convert(out)
	return nil
}

func (t treeNode) GetChildren(relativeDepth int, withOutSelf bool, out interface{}) (err error) {
	n := t.nodeI
	maxDepth := DEPTH_MAX
	if relativeDepth > 0 {
		maxDepth = n.GetDepth() + relativeDepth
	}
	parentPath := n.GetPath()
	if !withOutSelf {
		parentPath = fmt.Sprintf("%s/", parentPath)
	}
	err = n.GetAllByPathPrefix(parentPath, maxDepth, out)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

// ResetPath 重置路径和深度 通过parentID 递归生成父节点路径,避免父节点路径错误时,不会将错误放大,另外path 为冗余字断,主要用于优化查询,不做模型计算
func (t treeNode) ResetPath() (err error) {
	node := t.nodeI
	revNodeIdList := make([]string, 0)
	parent := node
	for {
		parentId := ""
		isRoot := true
		if !IsNil(parent) {
			parentId = parent.GetNodeID()
			isRoot = parent.IsRoot()
		}
		revNodeIdList = append(revNodeIdList, parentId)
		if isRoot {
			break
		}
		treeNode := NewTreeNode(parent)
		parentNodes := treeNodeIs{}
		parent, err = getParent(*treeNode, &parentNodes)
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
		w.WriteString(PathSpec)
		w.WriteString(nodeId)
	}
	path := w.String()
	depth := strings.Count(path, PathSpec)
	node.SetDepth(depth)
	node.SetPath(path)
	return nil
}

// calPath 计算节点迁移的新路径和深度
func (t treeNode) calPathAndDepth() (newPath string, diffDepth int, err error) {
	n := t
	parent, err := n.nodeI.GetParent()
	if err != nil {
		return newPath, diffDepth, err
	}
	newPath = fmt.Sprintf("%s%s", parent.GetPath(), n.nodeI.GetPath())
	diffDepth = parent.GetDepth() - n.nodeI.GetDepth() + 1
	return newPath, diffDepth, nil
}

// getParent 通过parentID 获取父类,使用path作为性能优化,主要用于设置路径情景
func getParent(node treeNode, cacheNodes *treeNodeIs) (parent TreeNodeI, err error) {
	parentID := node.nodeI.GetParentID()
	for _, parent := range *cacheNodes {
		if parent.GetNodeID() == parentID {
			return parent, nil
		}
	}
	parent, err = node.nodeI.GetParent()
	if err != nil {
		return nil, err
	}
	path := parent.GetPath()
	idList := strings.Split(path, PathSpec)
	if len(idList) > 0 {
		nodes := treeNodeIs{}
		err = node.nodeI.GetAllByNodeIds(idList, &nodes)
		if err != nil {
			return nil, err
		}
		*cacheNodes = append(*cacheNodes, nodes...)
	}
	return parent, err
}

func IsNil(v interface{}) bool {
	valueOf := reflect.ValueOf(v)
	k := valueOf.Kind()
	switch k {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return valueOf.IsNil()
	default:
		return v == nil
	}
}
