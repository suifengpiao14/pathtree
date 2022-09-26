package repository

import (
	"encoding/json"
	"testing"
)

func TestAddNode(t *testing.T) {
	rep := NewDoaRepository()
	input := AddNodeInput{
		Title:    "子节点",
		NodeID:   "1002",
		ParentID: "1001",
		Label:    "test",
		Depth:    2,
		Path:     "/1001/1002",
	}
	b, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	err = rep.AddNode(b)
	if err != nil {
		panic(err)
	}
}
