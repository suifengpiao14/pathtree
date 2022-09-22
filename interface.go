package treeentity

type RepositoryInterface interface {
	GetNode(nodeId string, output interface{}) (err error)
	GetTreeLimitDepth(parentPath string, depth int, output interface{}) (err error)
	GetTreeNodeCount(path string, output interface{}) (err error)
	GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error)
}
