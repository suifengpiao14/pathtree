package treeentity

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	LABEL_LEAF                = "leaf"
	ERROR_NOT_FOUND           = "404:404000001:not found"
	ERROR_ADD_NODE_LABLE_LEAF = "403:403000002:Leaf node is not allowed to add child nodes"
	ERROR_ADD_NODE_EXISTS     = "400:400000003:node id exists"
)

var DEPTH_MAX = 100000 //最大深度值
// nodeModel 树结构模型(只能在当前包内使用,离开当前包无法使用)
type nodeEntity struct {
	NodeID      string `json:"nodeId"`
	ParentID    string `json:"parentId"`
	Label       string `json:"label"`
	Depth       int    `json:"depth,string"`
	Path        string `json:"path"`
	_repository RepositoryInterface
}

// NewNodeEntity 包外唯一获得nodeEntity 方法
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
	var parent nodeEntity
	nodeIdList := make([]string, 0)
	nodeIdList = append(nodeIdList, nodeId)
	if parentId != "" && parentId != "0" {
		nodeIdList = append(nodeIdList, parentId)
	}
	nodeList := make([]nodeEntity, 0)
	err = n._repository.GetAllNodeByNodeIds(nodeIdList, &nodeList)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	for _, node := range nodeList {
		switch node.NodeID {
		case nodeId:
			err = errors.Errorf("%s;nodeId:%s", ERROR_ADD_NODE_EXISTS, nodeId)
			return nil, err
		case parentId:
			parent = node
		}
	}
	if parent.Label == LABEL_LEAF {
		err = errors.Errorf("%s;nodeId:%s", ERROR_ADD_NODE_LABLE_LEAF, parent.NodeID)
		return nil, err
	}
	path := fmt.Sprintf("/%s", nodeId)
	var diffDepth int
	node := nodeEntity{ // 重新赋值，不影响外部变量
		NodeID:   nodeId,
		ParentID: parentId,
		Label:    label,
		Path:     path,
		Depth:    strings.Count(path, "/"),
	}
	out.Path, diffDepth = calPath(node, parent)
	out.Depth = diffDepth + node.Depth
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

func (n *nodeEntity) GetSubTree(parentId string, depth int, withOutSelf bool, out interface{}) (err error) {
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

type moveSubTreeOut struct {
	NodeUpdateData     moveSubTreeOutNodeUpdateData        `json:"nodeUpdateData"`
	ChildrenUpdateData []*moveSubTreeOutChildrenUpdateData `json:"childrenUpdateData"`
}

type moveSubTreeOutNodeUpdateData struct {
	moveSubTreeOutChildrenUpdateData
	NewParentId string `json:"newParentId"`
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
	nodeMap, err := _getAllNodeMap(r, nodeIdList)
	if err != nil {
		return nil, err
	}
	node := nodeMap[nodeId]
	parent := nodeMap[newParentId]
	nodeOldPath := node.Path
	nodeNewPath, diffDepth := calPath(*node, *parent)
	newDepth := diffDepth + node.Depth
	// 修改node 节点本身
	out.NodeUpdateData = moveSubTreeOutNodeUpdateData{
		NewParentId: newParentId,
	}
	out.NodeUpdateData.moveSubTreeOutChildrenUpdateData = moveSubTreeOutChildrenUpdateData{
		NodeID:   nodeId,
		NewPath:  nodeNewPath,
		NewDepth: newDepth,
	}
	// 获取所有子节点
	var childrenNodeList []*nodeEntity
	err = r.GetAllByPathPrefix(node.Path, -1, &childrenNodeList)
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
	err = r.GetAllByPathPrefix(node.Path, -1, &childrenNodeList)
	if err != nil {
		return nil, err
	}
	nodeIdList = make([]string, 0)
	for _, childern := range childrenNodeList {
		nodeIdList = append(nodeIdList, childern.NodeID)
	}

	return nodeIdList, nil
}

// _getNode 根据节点ID获取节点数据，找不到数据，抛出错误 ERROR_NOT_FOUND,也可以由provider 直接返回错误,仅内部逻辑调用repository使用，因为返回值明确为 nodeEntity，其它数据会丢失
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

// _getAllNodeMap 批量获取节点，并且转换为map格式，其中一个nodeId有缺失，即返回错误
func _getAllNodeMap(r RepositoryInterface, nodeIdList []string) (nodeMap map[string]*nodeEntity, err error) {
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

// calPath 计算节点迁移的新路径和深度
func calPath(node nodeEntity, newParent nodeEntity) (newPath string, diffDepth int) {
	newPath = fmt.Sprintf("%s%s", newParent.Path, node.Path)
	diffDepth = newParent.Depth - node.Depth + 1
	return newPath, diffDepth
}

// BatchAddPathAndDepth 给所有数据，增加path和depth字段，方便批量数据导入 纯函数
func BatchAddPathAndDepth(data []map[string]interface{}, nodeIdKey string, parentIdKey string) (out []map[string]interface{}, err error) {
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

// BuildTree convert Two-dimensional array to tree
func BuildTree(records []map[string]interface{}, nodeIdKey string, parentIdKey string) (tree []*map[string]interface{}, err error) {
	tree = make([]*map[string]interface{}, 0)
	treeNodeMap := make(map[string]*map[string]interface{})

	for _, record := range records {
		nodeId := fmt.Sprintf("%v", record[nodeIdKey])     //change to string
		parentId := fmt.Sprintf("%v", record[parentIdKey]) //change to string
		_, ok := treeNodeMap[parentId]
		emptyParentId := parentId == "" || parentId == "0" // int 0 will be "0" after fmt.Sprintf("%v",)
		if !emptyParentId && !ok {
			parent := &map[string]interface{}{
				nodeIdKey: parentId,
			}
			treeNodeMap[parentId] = parent
		}
		//node := &record
		node := &map[string]interface{}{}
		for key, val := range record { // 确保对入参不产生副作用,不循环引用
			(*node)[key] = val
		}
		if tmpNode, ok := treeNodeMap[nodeId]; ok {
			children, ok := (*tmpNode)["children"]
			if ok {
				(*node)["children"] = children
			}

		}
		treeNodeMap[nodeId] = node
		// 根节点收集
		if emptyParentId {
			tree = append(tree, node)
		} else {
			// 子节点收集
			children := make([]*map[string]interface{}, 0)
			childrenIterface, ok := (*(treeNodeMap[parentId]))["children"]
			if ok {
				children = childrenIterface.([]*map[string]interface{})
			}
			children = append(children, node)
			(*(treeNodeMap[parentId]))["children"] = children
		}

	}
	return tree, nil
}
