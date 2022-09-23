package repository

import "encoding/json"

type sqlRepository struct{}

func NewSqlRepository() (rep *sqlRepository) {
	return &sqlRepository{}
}

type AddNodeInput struct {
	Code     string `json:"code"`
	Title    string `json:"title"`
	NodeID   string `json:"nodeId"`
	ParentID string `json:"parentId"`
	Label    string `json:"label"`
	Depth    int    `json:"depth,string"`
	Path     string `json:"path"`
}

func (r *sqlRepository) AddNode(data []byte) (err error) {
	input := AddNodeInput{}
	err = json.Unmarshal(data, &input)
	if err != nil {
		return err
	}
	entity := GenDistrictSQLInsertEntity{
		Code:       input.NodeID,
		Depth:      input.Depth,
		Label:      input.Label,
		ParentCode: input.ParentID,
		Path:       input.Path,
		Title:      input.Title,
	}
	entity.

	return err
}
func (r *sqlRepository) GetNode(nodeId string, output interface{}) (err error) {
	return err
}
func (r *sqlRepository) GetAllByPathPrefixWithDepth(parentPath string, depth int, output interface{}) (err error) {
	return err
}
func (r *sqlRepository) CountByPathPrefix(path string, output interface{}) (err error) {
	return err
}
func (r *sqlRepository) GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error) {
	return err
}
