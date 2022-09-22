package entity

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/onebehaviorentity"
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
	_do         func() (out interface{}, err error)
	onebehaviorentity.Onebehaviorentity
}

//NewNodeEntity 包外唯一获得nodeEntity 方法
func NewNodeEntity(repository RepositoryInterface) (node *nodeEntity) {
	node = &nodeEntity{
		_repository: repository,
	}
	return node
}

var addNodeInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=nodeId,dst=nodeId,required
fullname=label,dst=label,required
fullname=parentId,dst=parentId
`

func (n *nodeEntity) AddNode(nodeId string, label string, parentId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{
		_repository: n._repository,
	}
	data := fmt.Sprintf(`{"nodeId":"%s","label":"%s","parentId":"%s"}`, nodeId, label, parentId)
	doEntity = node.Build(node, addNodeInputSchema, "", func() (out interface{}, err error) {
		out, err = addNodeEffect(node)
		return out, err
	}).In([]byte(data))
	return doEntity
}

// 虽然只有一个参数，但是也使用相同格式，主要方便统一验证，统一定制参数错误返回
var nodeIdInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=nodeId,dst=nodeId,required
`
var nodeOutSchema = `
version=http://json-schema.org/draft-07/schema#,id=out,direction=out
fullname=nodeId,dst=nodeId,required
fullname=parentId,dst=parentId,required
fullname=label,dst=label,required
fullname=depth,dst=depth,required
fullname=path,dst=path,required
`

func (n *nodeEntity) GetNode(nodeId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s"}`, nodeId)
	doEntity = node.Build(node, nodeIdInputSchema, nodeOutSchema, func() (out interface{}, err error) {
		err = node._repository.GetNode(node.NodeID, &out)
		return out, err
	}).In([]byte(data))
	return doEntity
}

var getAllParentInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=out,direction=out
fullname=nodeId,dst=nodeId,required
fullname=withOutSelf,dst=withOutSelf,required,enum=["true","false"]
`

func (n *nodeEntity) GetAllParent(nodeId string, withOutSelf bool) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s","withOutSelf":"%s"}`, nodeId, strconv.FormatBool(withOutSelf))
	doEntity = node.Build(node, nodeIdInputSchema, "", func() (out interface{}, err error) {
		err = getAllParentEffect(n._repository, nodeId, withOutSelf, out)
		return out, err
	}).In([]byte(data))
	return doEntity
}

var getSubTreeLimitDepthInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=parentId,dst=parentId,required
fullname=depth,dst=depth,required
fullname=withOutSelf,dst=withOutSelf,required,entm["true","false"]
`

func (n *nodeEntity) GetSubTreeLimitDepth(parentId string, depth int, withOutSelf bool) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"parentId":"%s","depth","%d","withOutSelf":"%s"}`, parentId, depth, strconv.FormatBool(withOutSelf))
	doEntity = node.Build(node, getSubTreeLimitDepthInputSchema, nodeOutSchema, func() (out interface{}, err error) {
		out, err = getSubTreeLimitDepthEffect(n._repository, parentId, depth, withOutSelf)
		return out, err
	}).In([]byte(data))

	return doEntity
}

var getSubTreeNodeCountInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=parentId,dst=parentId,required
fullname=withOutSelf,dst=withOutSelf,required,entm["true","false"]
`

func (n *nodeEntity) GetSubTreeNodeCount(nodeId string, withOutSelf bool) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s","withOutSelf":"%s"}`, nodeId, strconv.FormatBool(withOutSelf))
	doEntity = node.Build(node, getSubTreeNodeCountInputSchema, "", func() (out interface{}, err error) {
		out, err = getSubTreeNodeCountEffect(n._repository, nodeId, withOutSelf)
		return out, err
	}).In([]byte(data))
	return doEntity
}

var MoveSubTreeInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=nodeId,dst=nodeId,required
fullname=newParentId,dst=newParentId,required
`

func (n *nodeEntity) MoveSubTree(nodeId string, newParentId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s","newParentId":"%s"}`, nodeId, newParentId)
	doEntity = node.Build(node, MoveSubTreeInputSchema, "", func() (out interface{}, err error) {
		out, err = moveSubTreeEffect(n._repository, nodeId, newParentId)
		return out, err
	}).In([]byte(data))
	return doEntity
}

func (n *nodeEntity) DeleteTree(nodeId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s""}`, nodeId)
	doEntity = node.Build(node, nodeIdInputSchema, "", func() (out interface{}, err error) {
		out, err = deleteTreeEffect(n._repository, nodeId)
		return out, err
	}).In([]byte(data))
	return
}

type AddNodeOut struct {
	NodeID   string `json:"nodeId"`
	ParentID string `json:"parentId"`
	Label    string `json:"label"`
	Depth    int    `json:"depth"`
	Path     string `json:"path"`
}

//addNodeEffect 增加节点数据存储逻辑(非纯函数)
func addNodeEffect(n *nodeEntity) (out *AddNodeOut, err error) {
	out = &AddNodeOut{
		NodeID:   n.NodeID,
		ParentID: n.ParentID,
		Label:    n.Label,
	}
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
	}
	return node, nil
}

func getAllParentEffect(r RepositoryInterface, nodeId string, withOutSelf bool, out interface{}) (err error) {
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
func getSubTreeLimitDepthEffect(r RepositoryInterface, parentId string, depth int, withOutSelf bool) (out interface{}, err error) {
	parentNode, err := _getNode(r, parentId)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	maxDepth := DEPTH_MAX
	if depth > 0 {
		maxDepth = parentNode.Depth + depth
	}
	parentPath := parentNode.Path
	if withOutSelf {
		parentPath = fmt.Sprintf("%s/", parentPath)
	}
	err = r.GetTreeLimitDepth(parentPath, maxDepth, &out)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	return out, err
}

func getSubTreeNodeCountEffect(r RepositoryInterface, nodeId string, withOutSelf bool) (count interface{}, err error) {
	node, err := _getNode(r, nodeId)
	if err != nil {
		return 0, err
	}
	parentPath := node.Path
	if withOutSelf {
		parentPath = fmt.Sprintf("%s/", parentPath)
	}
	err = node._repository.GetTreeNodeCount(parentPath, &count)
	if err != nil {
		return nil, err
	}
	return count, nil
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

type MoveSubTreeOut struct {
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

func moveSubTreeEffect(r RepositoryInterface, nodeId string, newParentId string) (moveSubTreeOut *MoveSubTreeOut, err error) {
	moveSubTreeOut = &MoveSubTreeOut{
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
	moveSubTreeOut.NodeUpdateData = moveSubTreeOutNodeUpdateData{
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
		moveSubTreeOut.ChildrenUpdateData = append(moveSubTreeOut.ChildrenUpdateData, childrenUpdateData)
	}
	return moveSubTreeOut, nil
}

func deleteTreeEffect(r RepositoryInterface, nodeId string) (nodList []string, err error) {
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
	nodList = make([]string, 0)
	for _, childern := range childrenNodeList {
		nodList = append(nodList, childern.NodeID)
	}

	return nodList, nil
}

//calPath 计算节点迁移的新路径和深度
func calPath(node *nodeEntity, newParent *nodeEntity) (newPath string, diffDepth int) {
	newPath = fmt.Sprintf("%s%s", newParent.Path, node.Path)
	diffDepth = newParent.Depth - node.Depth + 1
	return newPath, diffDepth
}
