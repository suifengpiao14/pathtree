package entity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/onebehaviorentity"
)

const (
	LABEL_LEAF                   = "leaf"
	ERROR_NOT_FOUND              = "404:404000001:not found"
	ERROR_ADD_NODE_TO_LABLE_LEAF = "403:403000002:Leaf node is not allowed to add child nodes"
)

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
fullname=order,dst=order
`

func (n *nodeEntity) AddNode(input []byte) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{
		_repository: n._repository,
	}
	doEntity = node.Build(node, addNodeInputSchema, "", func() (out interface{}, err error) {
		err = addNodeEffect(node)
		return nil, err
	}).In(input)
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
fullname=order,dst=order,required
fullname=path,dst=path,required
`

func (n *nodeEntity) GetNode(nodeId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s"}`, nodeId)
	doEntity = node.Build(node, nodeIdInputSchema, nodeOutSchema, func() (out interface{}, err error) {
		err = node._repository.GetNode(node.NodeID, &out)
		if err != nil {
			return nil, err
		}
		return out, nil
	}).In([]byte(data))
	return doEntity
}

var getSubTreeLimitDepth = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=parentId,dst=parentId,required
`

func (n *nodeEntity) GetSubTreeLimitDepth(parentId string, depth int) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"parentId":"%s"}`, parentId)
	doEntity = node.Build(node, getSubTreeLimitDepth, nodeOutSchema, func() (out interface{}, err error) {
		out, err = getSubTreeLimitDepthEffect(n._repository, parentId, depth)
		if err != nil {
			return nil, err
		}
		return out, nil
	}).In([]byte(data))

	return doEntity
}

//
func (n *nodeEntity) GetSubTreeNodeCount(nodeId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s"}`, nodeId)
	doEntity = node.Build(node, nodeIdInputSchema, "", func() (out interface{}, err error) {
		out, err = getSubTreeNodeCountEffect(n._repository, nodeId)
		if err != nil {
			return nil, err
		}
		return out, nil
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
		err = moveSubTreeEffect(n._repository, nodeId, newParentId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}).In([]byte(data))
	return doEntity
}
func (n *nodeEntity) DeleteSubTree(nodePathPrefix string) (err error) {
	return err
}

func (n *nodeEntity) DeleteTree(nodeId string) (doEntity onebehaviorentity.OnebehaviorentityDoInterface) {
	node := &nodeEntity{}
	data := fmt.Sprintf(`{"nodeId":"%s""}`, nodeId)
	doEntity = node.Build(node, nodeIdInputSchema, "", func() (out interface{}, err error) {
		err = deleteTreeEffect(n._repository, nodeId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}).In([]byte(data))
	return
}

//addNodeEffect 增加节点数据存储逻辑(非纯函数)
func addNodeEffect(n *nodeEntity) (err error) {
	var parent *nodeEntity
	if n.ParentID != "" && n.ParentID != "0" {
		err = n._repository.GetNode(n.ParentID, parent)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	if parent != nil {
		if parent.Label == LABEL_LEAF {
			err = errors.Errorf("%s;nodeId:%s", ERROR_ADD_NODE_TO_LABLE_LEAF, parent.NodeID)
			return err
		}
		var diffDepth int
		n.Path, diffDepth = calPath(n, parent)
		n.Depth = diffDepth + n.Depth

	}

	n, err = addNodePure(n, parent, LABEL_LEAF)
	if err != nil {
		return err
	}
	err = n._repository.AddNode(n)
	if err != nil {
		return err
	}
	return nil
}

//addNodePure 增加节点业务逻辑(独立成纯函数，方便测试)
func addNodePure(n *nodeEntity, parent *nodeEntity, labelLeaf string) (out *nodeEntity, err error) {
	n.Path = fmt.Sprintf("/%s", n.NodeID)
	out = n

	out.Depth = strings.Count(strings.Trim(n.Path, "/"), "/") + 1
	return out, nil
}

//getNode 根据节点ID获取节点数据，找不到数据，抛出错误 ERROR_NOT_FOUND,也可以由provider 直接返回错误
func getNode(r RepositoryInterface, nodeId string) (node *nodeEntity, err error) {
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

func getSubTreeLimitDepthEffect(r RepositoryInterface, parentId string, depth int) (out interface{}, err error) {
	parentNode, err := getNode(r, parentId)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	maxDepth := parentNode.Depth + depth
	err = r.GetTreeLimitDepth(parentNode.Path, maxDepth, &out)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	return out, err
}

func getSubTreeNodeCountEffect(r RepositoryInterface, nodeId string) (count interface{}, err error) {
	node, err := getNode(r, nodeId)
	if err != nil {
		return 0, err
	}
	err = node._repository.GetTreeNodeCount(node.Path, &count)
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

func moveSubTreeEffect(r RepositoryInterface, nodeId string, newParentId string) (err error) {
	nodeIdList := []string{nodeId, newParentId}
	nodeMap, err := getAllNodeMap(r, nodeIdList)
	if err != nil {
		return err
	}
	node := nodeMap[nodeId]
	parent := nodeMap[newParentId]
	newPath, diffDepth := calPath(node, parent)
	newDepth := diffDepth + node.Depth
	// 修改node 节点本身
	err = r.UpdateParent(nodeId, newParentId, newPath, newDepth)
	if err != nil {
		return err
	}
	// 更新子节点路径
	err = r.UpdatePath(node.Path, newPath, diffDepth)
	if err != nil {
		return err
	}
	return nil
}

func deleteTreeEffect(r RepositoryInterface, nodeId string) (err error) {
	node, err := getNode(r, nodeId)
	if err != nil {
		return err
	}
	err = r.DeleteTree(node.Path)
	if err != nil {
		return err
	}
	return nil
}

//calPath 计算节点迁移的新路径和深度
func calPath(node *nodeEntity, newParent *nodeEntity) (newPath string, diffDepth int) {
	newPath = fmt.Sprintf("%s%s", newParent.Path, node.Path)
	diffDepth = newParent.Depth - node.Depth + 1
	return newPath, diffDepth
}
