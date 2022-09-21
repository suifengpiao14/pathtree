package entity

type RepositoryInterface interface {
	AddNode(data []byte) (err error)
	MoveSubTree(newPath string, oldPath string) (err error)
	UpdateParent(nodeId string, newParentId string, newPath string, newDepth int) (err error)
	UpdatePath(oldPathPrefix string, newPathPrefix string, diffDepth int) (err error)

	DeleteTree(pathPrefix string) (err error)

	GetNode(nodeId string, output interface{}) (err error)
	GetTreeLimitDepth(parentPath string, depth int, output interface{}) (err error)
	GetTreeNodeCount(path string, output interface{}) (err error)
	GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error)
}
