package entity

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/suifengpiao14/onebehaviorentity"
)

const (
	LABEL_LEAF = "leaf"
)

//nodeModel 树结构模型(只能在当前包内使用,离开当前包无法使用)
type nodeEntity struct {
	NodeID   string      `json:"nodeId"`
	Title    string      `json:"title"`
	ParentID string      `json:"parentId"`
	Label    string      `json:"label"`
	Depth    int         `json:"depth"`
	Order    int         `json:"order"`
	Path     string      `json:"path"`
	Kv       interface{} `json:"kv"`
	_do      func() (out interface{}, err error)
	onebehaviorentity.Onebehaviorentity
}

type nodeEntityFactory struct {
	_repository RepositoryInterface
}

//NewNodeEntity 包外唯一获得nodeEntity 方法
func NewNodeEntityFactory(repository RepositoryInterface) (factory *nodeEntityFactory) {
	factory = &nodeEntityFactory{}
	return factory
}

var addNodeInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=nodeId,dst=nodeId,required
fullname=title,dst=title,required
fullname=label,dst=label,required
fullname=parentId,dst=parentId
fullname=order,dst=order
`

func (f *nodeEntityFactory) BuildAddNodeBehavior() (node *nodeEntity, err error) {
	node = &nodeEntity{}
	node.Build(node, addNodeInputSchema, "", func() (out interface{}, err error) {
		err = addNodeEffect(f._repository, node)
		return nil, err
	})
	return node, err
}

var getNodeBehavior = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=nodeId,dst=nodeId,required
`

func (f *nodeEntityFactory) BuildGetNodeBehavior() (node *nodeEntity, err error) {
	node = &nodeEntity{}
	node.Build(node, getNodeBehavior, "", func() (out interface{}, err error) {
		f._repository.GetNode(node.NodeID)
		err = addNodeEffect(f._repository, node)
		return nil, err
	})
	return node, err
}

func (n *nodeEntity) BatchAddNode(jsonByte []byte) (err error) {
	return err
}
func (n *nodeEntity) GetNode(nodeId string, output interface{}) (err error) {
	if nodeId == "" {

	}

	return err
}
func (n *nodeEntity) GetSubTreeLimitDepth(parentPath string, depth int, output interface{}) (err error) {
	return err
}
func (n *nodeEntity) GetSubTreeNodeCount(nodeId string, output *int) (err error) {
	return err
}
func (n *nodeEntity) MoveSubTree(newParentId string, oldParentId string) (err error) {
	return err
}
func (n *nodeEntity) DeleteSubTree(nodePathPrefix string) (err error) {
	return err
}

//addNodeEffect 增加节点数据存储逻辑(非纯函数)
func addNodeEffect(r RepositoryInterface, n *nodeEntity) (err error) {
	var parent nodeEntity
	if n.ParentID != "" && n.ParentID != "0" {
		err = r.GetNode(n.ParentID, &parent)
		if err != nil {
			err = errors.WithStack(err)
			return err
		}
	}
	n, err = addNodePure(n, &parent, LABEL_LEAF)
	if err != nil {
		return err
	}
	err = r.AddNode(n)
	if err != nil {
		return err
	}
	return nil
}

//addNodePure 增加节点业务逻辑(纯函数)
func addNodePure(n *nodeEntity, parent *nodeEntity, labelLeaf string) (out *nodeEntity, err error) {
	n.Path = fmt.Sprintf("/%s", n.NodeID)
	out = n
	if parent != nil {
		if parent.Label == labelLeaf {
			err = errors.Errorf("node %s has label %s, can not add children", parent.NodeID, labelLeaf)
			return nil, err
		}
		out.Path = fmt.Sprintf("%s%s", parent.Path, n.Path)
	}
	out.Depth = strings.Count(strings.Trim(n.Path, "/"), "/") + 1
	return out, nil
}

func getNode(n *nodeEntity, nodeId string) (node *nodeEntity, err error) {
	node = &nodeEntity{}
	err = n._repository.GetNode(nodeId, node)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	return node, nil
}
