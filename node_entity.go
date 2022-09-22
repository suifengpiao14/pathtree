package treeentity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	LABEL_LEAF                   = "leaf"
	ERROR_NOT_FOUND              = "404:404000001:not found"
	ERROR_ADD_NODE_TO_LABLE_LEAF = "403:403000002:Leaf node is not allowed to add child nodes"
)

var DEPTH_MAX = 100000 //最大深度值
//nodeModel 树结构模型(只能在当前包内使用,离开当前包无法使用)
type nodeEntity struct {
	NodeID      string `json:"nodeId"`
	ParentID    string `json:"parentId"`
	Label       string `json:"label"`
	Depth       int    `json:"depth"`
	Path        string `json:"path"`
	_repository RepositoryInterface
}

//NewNodeEntity 包外唯一获得nodeEntity 方法
func NewNodeEntity(repository RepositoryInterface) (node *nodeEntity) {
	node = &nodeEntity{
		_repository: repository,
	}
	return node
}

type AddNodeOut struct {
	Depth int    `json:"depth"`
	Path  string `json:"path"`
}

func (n *nodeEntity) AddNode(nodeId string, parentId string, label string) (out *AddNodeOut, err error) {
	out = &AddNodeOut{}
	var parent *nodeEntity
	if n.ParentID != "" && n.ParentID != "0" {
		err = n._repository.GetNode(n.ParentID, parent)
		if err != nil {
			err = errors.WithStack(err)
			return nil, err
		}
	}
	if parent != nil {
		if parent.Label == LABEL_LEAF {
			err = errors.Errorf("%s;nodeId:%s", ERROR_ADD_NODE_TO_LABLE_LEAF, parent.NodeID)
			return nil, err
		}
		var diffDepth int
		out.Path, diffDepth = calPath(n, parent)
		out.Depth = diffDepth + n.Depth
	}

	return out, nil
}

func (n *nodeEntity) GetNode(nodeId string, out interface{}) (err error) {
	err = n._repository.GetNode(nodeId, out)
	return err
}

func (n *nodeEntity) GetAllParent(nodeId string, withOutSelf bool, out interface{}) (err error) {
	r := n._repository
	node, err := _getNode(r, nodeId)
	if err != nil {
		return err
	}
	nodeIdList := strings.Split(node.Path, "/")
	if len(nodeIdList) == 0 {
		return nil
	}
	if withOutSelf {
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

func (n *nodeEntity) GetSubTreeLimitDepth(parentId string, depth int, withOutSelf bool, out interface{}) (err error) {
	r := n._repository
	parentNode, err := _getNode(r, parentId)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	maxDepth := DEPTH_MAX
	if depth > 0 {
		maxDepth = parentNode.Depth + depth
	}
	parentPath := parentNode.Path
	if withOutSelf {
		parentPath = fmt.Sprintf("%s/", parentPath)
	}
	err = r.GetTreeLimitDepth(parentPath, maxDepth, out)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

func (n *nodeEntity) GetSubTreeNodeCount(nodeId string, withOutSelf bool, count *int) (err error) {
	r := n._repository
	node, err := _getNode(r, nodeId)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	parentPath := node.Path
	if withOutSelf {
		parentPath = fmt.Sprintf("%s/", parentPath)
	}
	err = node._repository.GetTreeNodeCount(parentPath, count)
	if err != nil {
		err = errors.WithStack(err)
		return err
	}
	return nil
}

type moveSubTreeOut struct {
	NodeUpdateData     moveSubTreeOutNodeUpdateData        `json:"nodeUpdateData"`
	ChildrenUpdateData []*moveSubTreeOutChildrenUpdateData `json:"childrenUpdateData"`
}

type moveSubTreeOutNodeUpdateData struct {
	NodeID      string `json:"nodeId"`
	NewParentId string `json:"newParentId"`
	NewPath     string `json:"newPath"`
	NewDepth    int    `json:"newDepth,string"`
}

type moveSubTreeOutChildrenUpdateData struct {
	NodeID   string `json:"nodeId"`
	NewPath  string `json:"newPath"`
	NewDepth int    `json:"newDepth,string"`
}

func (n *nodeEntity) MoveSubTree(nodeId string, newParentId string) (out *moveSubTreeOut, err error) {
	r := n._repository
	out = &moveSubTreeOut{
		NodeUpdateData:     moveSubTreeOutNodeUpdateData{},
		ChildrenUpdateData: make([]*moveSubTreeOutChildrenUpdateData, 0),
	}
	nodeIdList := []string{nodeId, newParentId}
	nodeMap, err := getAllNodeMap(r, nodeIdList)
	if err != nil {
		return nil, err
	}
	node := nodeMap[nodeId]
	parent := nodeMap[newParentId]
	nodeOldPath := node.Path
	nodeNewPath, diffDepth := calPath(node, parent)
	newDepth := diffDepth + node.Depth
	// 修改node 节点本身
	out.NodeUpdateData = moveSubTreeOutNodeUpdateData{
		NodeID:      nodeId,
		NewParentId: newParentId,
		NewPath:     nodeNewPath,
		NewDepth:    newDepth,
	}
	// 获取所有子节点
	var childrenNodeList []*nodeEntity
	err = r.GetTreeLimitDepth(node.Path, -1, &childrenNodeList)
	if err != nil {
		return nil, err
	}
	// 更新子节点路径和深度值
	for _, children := range childrenNodeList {
		newPath := strings.Replace(children.Path, nodeOldPath, nodeNewPath, 1)
		newDepth := children.Depth + diffDepth
		childrenUpdateData := &moveSubTreeOutChildrenUpdateData{
			NodeID:   children.NodeID,
			NewPath:  newPath,
			NewDepth: newDepth,
		}
		out.ChildrenUpdateData = append(out.ChildrenUpdateData, childrenUpdateData)
	}
	return out, nil
}

func (n *nodeEntity) DeleteTree(nodeId string) (nodeIdList []string, err error) {
	r := n._repository
	// 获取节点
	var node nodeEntity
	err = r.GetNode(nodeId, &node)
	if err != nil {
		return nil, err
	}
	// 获取所有子节点
	var childrenNodeList []*nodeEntity
	err = r.GetTreeLimitDepth(node.Path, -1, &childrenNodeList)
	if err != nil {
		return nil, err
	}
	nodeIdList = make([]string, 0)
	for _, childern := range childrenNodeList {
		nodeIdList = append(nodeIdList, childern.NodeID)
	}

	return nodeIdList, nil
}

//_getNode 根据节点ID获取节点数据，找不到数据，抛出错误 ERROR_NOT_FOUND,也可以由provider 直接返回错误,仅内部逻辑调用repository使用，因为返回值明确为 nodeEntity，其它数据会丢失
func _getNode(r RepositoryInterface, nodeId string) (node *nodeEntity, err error) {
	node = &nodeEntity{}
	err = r.GetNode(nodeId, node)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	if node.NodeID == "" {
		err = errors.Errorf("%s;nodeId:%s", ERROR_NOT_FOUND, nodeId)
		return nil, err
	}
	return node, nil
}

//getAllNodeMap 批量获取节点，并且转换为map格式，其中一个nodeId有缺失，即返回错误
func getAllNodeMap(r RepositoryInterface, nodeIdList []string) (nodeMap map[string]*nodeEntity, err error) {
	nodeList := make([]*nodeEntity, 0)
	err = r.GetAllNodeByNodeIds(nodeIdList, nodeList)
	if err != nil {
		return nil, err
	}
	nodeMap = make(map[string]*nodeEntity, 0)
	for _, node := range nodeList {
		nodeMap[node.NodeID] = node
	}

	//validate
	for _, nodeId := range nodeIdList {
		_, ok := nodeMap[nodeId]
		if !ok {
			err = errors.Errorf("%s;nodeId:%s", ERROR_NOT_FOUND, nodeId)
			return nil, err
		}
	}
	return nodeMap, nil
}

//calPath 计算节点迁移的新路径和深度
func calPath(node *nodeEntity, newParent *nodeEntity) (newPath string, diffDepth int) {
	newPath = fmt.Sprintf("%s%s", newParent.Path, node.Path)
	diffDepth = newParent.Depth - node.Depth + 1
	return newPath, diffDepth
}
