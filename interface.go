package treeentity

type RepositoryInterface interface {
	GetNode(nodeId string, output interface{}) (err error)
	GetAllByPathPrefix(pathPrefix string, depth int, output interface{}) (err error)
	GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error)
}
