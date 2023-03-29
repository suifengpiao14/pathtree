package treeentity_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitea.programmerfamily.com/go/treeentity"
)

func TestGetNode(t *testing.T) {
	var repository treeentity.RepositoryInterface
	instance := treeentity.NewNodeEntity(repository)
	nodeId := "first"
	label := "test"
	parentId := ""
	out, err := instance.AddNode(nodeId, parentId, label)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func TestBuildTree(t *testing.T) {
	// jsonStr := `[
	// {"nodeId":"B","parentId":"A","name":"first"},
	// {"nodeId":"C","parentId":"B","name":"first-chirden"},
	// {"nodeId":"A","parentId":"","name":"root"},
	// {"nodeId":"D","parentId":"A","name":"second"}
	// ]`
	jsonStr2 := `[
	{"nodeId":4,"parentId":1,"name":"first"},
	{"nodeId":5,"parentId":4,"name":"first-chirden"},
	{"nodeId":1,"parentId":0,"name":"root"},
	{"nodeId":3,"parentId":1,"name":"second"}
	]`
	recordList := make([]map[string]interface{}, 0)
	json.Unmarshal([]byte(jsonStr2), &recordList)
	out, err := treeentity.BuildTree(recordList, "nodeId", "parentId")
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	str := string(b)
	fmt.Println(str)
}

func TestBatchAddPathAndDepth(t *testing.T) {
	data := `[
		{"code":1,"parentCode":0},
		{"code":2,"parentCode":1},
		{"code":3,"parentCode":2},
		{"code":4,"parentCode":1}
	]`
	record := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(data), &record)
	if err != nil {
		panic(err)
	}
	out, err := treeentity.ResetAllPath1(record, "code", "parentCode")
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}
	str := string(b)
	fmt.Println(str)
}

func TestChildrenCount(t *testing.T) {
	data := `[
		{"code":1,"parentCode":0,"path":"/1"},
		{"code":2,"parentCode":1,"path":"/1/2"},
		{"code":3,"parentCode":2,"path":"/1/2/3"},
		{"code":4,"parentCode":1,"path":"/1/4"}
	]`
	records := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(data), &records)
	if err != nil {
		panic(err)
	}
	out := treeentity.ChildrenCount(records, "code", "path")
	fmt.Println(out)
}
