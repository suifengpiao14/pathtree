package entity

type RepositoryInterface interface {
	AddNode(node *nodeEntity) (err error)
	BatchAddNode(nodeList []*nodeEntity) (err error)
	GetNode(nodeId string, output *nodeEntity) (err error)
	GetSubTreeLimitDepth(parentPath string, depth int, output []*nodeEntity) (err error)
	GetSubTreeNodeCount(nodeId string, output int) (err error)
	MoveSubTree(newPath string, oldPath string) (err error)
	DeleteSubTree(nodePathPrefix string) (err error)
}
