package tree

import (
	"encoding/json"
	"strconv"

	"github.com/suifengpiao14/model"
	"github.com/suifengpiao14/templatemap/provider"
)

type Tree interface {
	TableSQL() (sql string)
	AddNode(simpleNode *SimpleNodeModel) (err error)
	BatchAddNode(nodeList []*NodeModel) (err error)
	GetNode(nodeId string) (node *NodeModel, err error)
	GetSubTreeLimitDepth(parentPath string, depth int) (nodeList []*NodeModel, err error)
	GetSubTreeNodeCount(nodeId string) (count int, err error)
	MoveSubTree(newPath string, oldPath string) (err error)
	DeleteSubTree(nodePathPrefix string) (err error)
	GetProvider() (provider provider.ExecproviderInterface)
}

type tree struct {
	model.Model
}

//NewTree 生成一个tree实例
func NewTree(table string, provider provider.ExecproviderInterface) Tree {
	return &tree{
		Model: model.Model{
			Table:    table,
			Provider: provider,
		},
	}
}

func (t *tree) GetProvider() (provider provider.ExecproviderInterface) {
	return t.Provider
}

func (t *tree) TableSQL() (sql string) {
	return TableSQL(t.Table)
}
func (t *tree) AddNode(simpleNode *SimpleNodeModel) (err error) {
	sql := AddNode(t.Table, simpleNode)
	_, err = t.GetProvider().Exec(t.Identity, sql.SQL)
	return err
}

func (t *tree) BatchAddNode(nodeList []*NodeModel) (err error) {
	sql := BatchAddNode(t.Table, nodeList)
	_, err = t.GetProvider().Exec(t.Identity, sql.SQL)
	return err
}

func (t *tree) GetNode(nodeId string) (node *NodeModel, err error) {
	sql := GetNode(t.Table, nodeId)
	data, err := t.GetProvider().Exec(t.Identity, sql.SQL)
	if err != nil {
		return nil, err
	}
	node = &NodeModel{}
	err = json.Unmarshal([]byte(data), node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (t *tree) GetSubTreeLimitDepth(parentPath string, depth int) (nodeList []*NodeModel, err error) {
	sql := GetSubTreeLimitDepth(t.Table, parentPath, depth)
	data, err := t.GetProvider().Exec(t.Identity, sql.SQL)
	if err != nil {
		return nil, err
	}
	nodeList = make([]*NodeModel, 0)
	err = json.Unmarshal([]byte(data), &nodeList)
	if err != nil {
		return nil, err
	}
	return nodeList, nil
}

func (t *tree) GetSubTreeNodeCount(nodeId string) (count int, err error) {
	sql := GetSubTreeNodeCount(t.Table, nodeId)
	data, err := t.GetProvider().Exec(t.Identity, sql.SQL)
	if err != nil {
		return 0, err
	}
	count, err = strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (t *tree) MoveSubTree(newPath string, oldPath string) (err error) {
	sql := MoveSubTree(t.Table, newPath, oldPath)
	_, err = t.GetProvider().Exec(t.Identity, sql.SQL)
	if err != nil {
		return err
	}
	return nil
}
func (t *tree) DeleteSubTree(nodePathPrefix string) (err error) {
	sql := DeleteSubTree(t.Table, nodePathPrefix)
	_, err = t.GetProvider().Exec(t.Identity, sql.SQL)
	return err
}
