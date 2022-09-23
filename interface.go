package treeentity

type RepositoryInterface interface {
	GetNode(nodeId string, output interface{}) (err error)
	GetAllByPathPrefixWithDepth(patthPrefix string, depth int, output interface{}) (err error)
	CountByPathPrefix(pathPrefix string, output interface{}) (err error)
	GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error)
}
