package tree

type Tree interface {
	TableSQL() (sql string)
	AddNode(simpleNode *SimpleNodeModel) (sql string)
	BatchAddNode(nodeList []*NodeModel) (sql string)
	GetNode(nodeId string) (sql string)
	GetSubTreeLimitDepth(parentPath string, depth int) (sql string)
	GetSubTreeNodeCount(nodeId string) (sql string)
	MoveSubTree(newPath string, oldPath string) (sql string)
	DeleteSubTree(nodePathPrefix string) (sql string)
}

type tree struct {
	table string
}

//NewTree 生成一个tree实例
func NewTree(table string) Tree {
	return &tree{
		table: table,
	}
}
func (t *tree) TableSQL() (sql string) {
	return TableSQL(t.table)
}
func (t *tree) AddNode(simpleNode *SimpleNodeModel) (sql string) {
	return AddNode(t.table, simpleNode)
}

func (t *tree) BatchAddNode(nodeList []*NodeModel) (sql string) {
	return BatchAddNode(t.table, nodeList)
}

func (t *tree) GetNode(nodeId string) (sql string) {
	return GetNode(t.table, nodeId)
}

func (t *tree) GetSubTreeLimitDepth(parentPath string, depth int) (sql string) {
	return GetSubTreeLimitDepth(t.table, parentPath, depth)
}

func (t *tree) GetSubTreeNodeCount(nodeId string) (sql string) {
	return GetSubTreeNodeCount(t.table, nodeId)
}

func (t *tree) MoveSubTree(newPath string, oldPath string) (sql string) {
	return MoveSubTree(t.table, newPath, oldPath)
}
func (t *tree) DeleteSubTree(nodePathPrefix string) (sql string) {
	return DeleteSubTree(t.table, nodePathPrefix)
}
