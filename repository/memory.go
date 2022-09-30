package repository

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type memoryRepository struct {
	data string
}

func NewMemoryRepository(data string) (rep *memoryRepository) {
	return &memoryRepository{
		data: data,
	}
}

func (r *memoryRepository) GetNode(nodeId string, output interface{}) (err error) {
	path := fmt.Sprintf(`#(nodeId=="%s").0`, nodeId)
	res := gjson.Get(r.data, path)
	if !res.Exists() {
		err = errors.Errorf("404:4044:not found %s", nodeId)
		return err
	}
	err = json.Unmarshal([]byte(res.String()), output)
	return nil
}
func (r *memoryRepository) GetAllByPathPrefix(patthPrefix string, depth int, output interface{}) (err error) {
	return nil
}
func (r *memoryRepository) GetAllNodeByNodeIds(nodeIds []string, output interface{}) (err error) {
	return nil
}
