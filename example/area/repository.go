package area

import (
	"gitea.programmerfamily.com/go/treeentity"
	"github.com/pkg/errors"
)

type areaRecordRepository struct {
}

func (r *areaRecordRepository) AddNode(node treeentity.TreeNode) (err error) {
	err = errors.Errorf("Not implemented")
	return err
}
func (r *areaRecordRepository) UpdateNode(node treeentity.TreeNode) (err error) {
	err = errors.Errorf("Not implemented")
	return err
}
func (r *areaRecordRepository) UpdateBatchNode(nodes []treeentity.TreeNode) (err error) {
	err = errors.Errorf("Not implemented")
	return err
}
func (r *areaRecordRepository) GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error) {
	return nil
}
func (r *areaRecordRepository) GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error) {
	return nil
}
func (r *areaRecordRepository) GetByAreaID(areaID string) (areaRecord *AreaRecord, err error) {
	return nil, nil
}
func (r *areaRecordRepository) GetByLevel(depth string) (areaRecord AreaRecords, err error) {
	return nil, nil
}
func (r *areaRecordRepository) GetByKeyWord(keyword string, depth string) (areaRecord AreaRecords, err error) {
	return nil, nil
}
