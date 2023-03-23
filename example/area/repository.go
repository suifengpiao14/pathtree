package area

import "gitea.programmerfamily.com/go/treeentity"

type NodeRepository struct {
}

func (r *NodeRepository) AddNode(node treeentity.TreeNode) (err error) {
	return err
}
func (r *NodeRepository) UpdateNode(node treeentity.TreeNode) (err error) {
	return err
}
func (r *NodeRepository) UpdateBatchNode(nodes []treeentity.TreeNode) (err error) {
	return err
}
func (r *NodeRepository) GetAllByPathPrefix(pathPrefix string, depth int, nodes *[]treeentity.TreeNode) (err error) {
	return err
}
func (r *NodeRepository) GetAllNodeByNodeIds(nodeIds []string, nodes *[]treeentity.TreeNode) (err error) {
	return err
}
