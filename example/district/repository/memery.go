package repository

type memeryRepository struct{}

func NewMemeryRepository() (rep *memeryRepository) {
	return &memeryRepository{}
}

func (r *memeryRepository) AddNode(node interface{}) (err error) {
	return err
}
func (r *memeryRepository) GetNode(nodeId string, output interface{}) (err error) {
	return err
}
func (r *memeryRepository) GetTreeLimitDepth(parentPath string, depth int, output interface{}) (err error) {
	return err
}
func (r *memeryRepository) GetTreeNodeCount(path string, output interface{}) (err error) {
	return err
}
func (r *memeryRepository) GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error) {
	return err
}
