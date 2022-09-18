package entity

import (
	"github.com/suifengpiao14/onebehaviorentity"
)

//nodeModel 树结构模型(只能在当前包内使用,离开当前包无法使用)
type nodeEntity struct {
	NodeID      string `json:"nodeId"`
	Title       string `json:"title"`
	ParentID    string `json:"parentId"`
	Label       string `json:"label"`
	Depth       int    `json:"depth"`
	Order       int    `json:"order"`
	Path        string `json:"path"`
	Ext         string `json:"ext"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
	_do         func() (out interface{}, err error)
	_repository RepositoryInterface
	onebehaviorentity.Onebehaviorentity
}

//NewNodeEntity 包外唯一获得nodeEntity 方法
func NewNodeEntity(repository RepositoryInterface) (node *nodeEntity) {
	node = &nodeEntity{
		_repository: repository,
	}
	return node
}

func (n *nodeEntity) Do() (out interface{}, err error) {
	out, err = n._do()
	return out, err
}

var addNodeInputSchema = `
version=http://json-schema.org/draft-07/schema#,id=input,direction=in
fullname=nodeId,dst=nodeId,required
fullname=title,dst=title,required
fullname=parentId,dst=parentId
fullname=label,dst=label
fullname=order,dst=order
`

func (n *nodeEntity) AddNode(jsonByte []byte) (err error) {
	n.Build(n, addNodeInputSchema, "", func() (out interface{}, err error) {
		err = n._repository.AddNode(n)
		return nil, err
	})
	n.In(jsonByte)
	err = n.Out(nil)
	if err != nil {
		return err
	}
	return nil
}

func (n *nodeEntity) BatchAddNode(jsonByte []byte) (err error) {
	return err
}
func (n *nodeEntity) GetNode(nodeId string, output interface{}) (err error) {
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
