package repository

import (
	"fmt"
)

type memoryRepository struct {
	nodeIdKey   string
	parentIdKey string
	data        map[string]map[string]interface{}
}

func NewMemoryRepository(data []map[string]interface{}, nodeIdKey string, parentIdKey string) (rep *memoryRepository) {
	rep = &memoryRepository{
		data:        make(map[string]map[string]interface{}, 0),
		nodeIdKey:   nodeIdKey,
		parentIdKey: parentIdKey,
	}
	for _, record := range data {
		nodeId := fmt.Sprintf("%v", record[nodeIdKey])
		record["nodeId"] = nodeId
		record["parentId"] = record[parentIdKey]
		rep.data[nodeId] = record
	}
	return rep
}

func (r *memoryRepository) GetNode(nodeId string, output interface{}) (err error) {
	return nil
}

func (r *memoryRepository) GetAllByPathPrefix(patthPrefix string, depth int, output interface{}) (err error) {
	return nil
}
func (r *memoryRepository) UpdatePath(nodeId string, path string, depth int) {
	r.data[nodeId]["path"] = path
	r.data[nodeId]["depth"] = depth
}

func (r *memoryRepository) AllNodeIdParentIdMap() (out map[string]string) {
	out = map[string]string{}
	for nodeId, record := range r.data {
		parentId := fmt.Sprintf("%v", record["parentId"])
		out[nodeId] = parentId
	}
	return out
}

// GetData 获取所有数据
func (r *memoryRepository) GetData() (out []map[string]interface{}) {
	out = make([]map[string]interface{}, 0)
	for _, record := range r.data {
		delete(record, "nodeId")
		delete(record, "parentId")
		out = append(out, record)
	}
	return out
}
