package treeentity

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var DEPTH_MIN = -1     //最小深度值
var DEPTH_MAX = 100000 //最大深度值

type TreeRepositoryI interface {
	AddNode(node TreeNodeI) (err error)
	UpdateNode(node TreeNodeI) (err error)
	UpdateBatchNode(nodes []TreeNodeI) (err error)
	GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error)
	GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error)
}

type BaseTreeNodeI interface {
	GetNodeID() (nodeID string)
	SetPath(path string)
	SetDepth(depth int)
	GetParent() (parent TreeNodeI, err error)
}

type TreeNodeI interface {
	BaseTreeNodeI
	GetPath() (path string)
	GetDepth() (depth int)
	SetParentID(parentId string)
}
type TreeNodes []TreeNodeI

type TreeNodesI interface {
	Len() (len int)
	GetByIndex(i int) (node BaseTreeNodeI)
}

type tree struct {
	node       TreeNodeI
	repository TreeRepositoryI
}

func NewTree(node TreeNodeI, repository TreeRepositoryI) (t *tree) {
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
	var childrenNodeList []TreeNodeI
	err = r.GetAllByPathPrefix(nodeOldPath, -1, &childrenNodeList)
	if err != nil {
		return err
	}
	// 更新子节点路径和深度值
	newChildren := make([]TreeNodeI, 0)
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
	var childrenNodeList []TreeNodeI
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
	nodes := make([]TreeNodeI, 0)
	err = r.GetAllByNodeIds(nodeIdList, &nodes)
	if err != nil {
		return err
	}
	minDepth := DEPTH_MIN
	if relativeDepth > 0 {
		minDepth = n.node.GetDepth() - relativeDepth
	}
	outNodes := make([]TreeNodeI, 0)
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

func ResetAllPath(nodes TreeNodesI) (err error) {
	count := nodes.Len()
	nodeIdParentIdMap := map[string]string{}
	for i := 0; i < count; i++ {
		node := nodes.GetByIndex(i)
		nodeId := node.GetNodeID()
		parent, err := node.GetParent()
		if err != nil {
			return err
		}
		parentId := parent.GetNodeID()
		if _, ok := nodeIdParentIdMap[nodeId]; ok {
			err := errors.Errorf("dumplicate nodeId:%s", nodeId)
			return err
		}
		nodeIdParentIdMap[nodeId] = parentId
	}
	// 循环处理数据,增加path和depth
	for i := 0; i < count; i++ {
		node := nodes.GetByIndex(i)
		parent, err := node.GetParent()
		if err != nil {
			return err
		}
		nodeId := node.GetNodeID()
		parentId := parent.GetNodeID()
		revNodeIdList := make([]string, 0)
		revNodeIdList = append(revNodeIdList, nodeId) // 由下到上收集节点ID
		for {
			emptyParentId := parentId == "" || parentId == "0" // int 0 will be "0" after fmt.Sprintf("%v",)
			if emptyParentId {
				break
			}
			revNodeIdList = append(revNodeIdList, parentId)
			newParentId, ok := nodeIdParentIdMap[parentId]
			if !ok {
				err = errors.Errorf("not found record by parentId:%s", parentId)
				return err
			}
			parentId = newParentId
		}
		var w bytes.Buffer
		l := len(revNodeIdList)
		for i := l - 1; i > -1; i-- {
			w.WriteString("/")
			w.WriteString(revNodeIdList[i])
		}
		path := w.String()
		depth := strings.Count(path, "/")
		node.SetDepth(depth)
		node.SetPath(path)
	}
	return nil
}

// ResetAllPath1 给所有数据，重置path和depth字段，方便批量数据导入 纯函数
func ResetAllPath1(data []map[string]interface{}, nodeIdKey string, parentIdKey string) (out []map[string]interface{}, err error) {
	nodeIdParentIdMap := map[string]string{}
	dataMap := map[string]map[string]interface{}{}
	for _, record := range data {
		nodeId := fmt.Sprintf("%v", record[nodeIdKey])
		parentId := fmt.Sprintf("%v", record[parentIdKey])
		if _, ok := nodeIdParentIdMap[nodeId]; ok {
			err := errors.Errorf("dumplicate %s:%s", nodeIdKey, nodeId)
			return nil, err
		}
		nodeIdParentIdMap[nodeId] = parentId
	}
	for _, record := range data {
		nodeId := fmt.Sprintf("%v", record[nodeIdKey])
		parentId := fmt.Sprintf("%v", record[parentIdKey])
		revNodeIdList := make([]string, 0)
		revNodeIdList = append(revNodeIdList, nodeId) // 由下到上收集节点ID
		for {
			emptyParentId := parentId == "" || parentId == "0" // int 0 will be "0" after fmt.Sprintf("%v",)
			if emptyParentId {
				break
			}
			revNodeIdList = append(revNodeIdList, parentId)
			newParentId, ok := nodeIdParentIdMap[parentId]
			if !ok {
				err = errors.Errorf("not found record by %s:%s", nodeIdKey, parentId)
				return nil, err
			}
			parentId = newParentId
		}
		var w bytes.Buffer
		l := len(revNodeIdList)
		for i := l - 1; i > -1; i-- {
			w.WriteString("/")
			w.WriteString(revNodeIdList[i])
		}
		path := w.String()
		depth := strings.Count(path, "/")
		record["path"] = path
		record["depth"] = depth
		dataMap[nodeId] = record
	}
	out = make([]map[string]interface{}, 0)
	for _, record := range dataMap {
		out = append(out, record)
	}
	return out, nil
}
