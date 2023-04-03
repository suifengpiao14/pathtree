package pathtree_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"gitea.programmerfamily.com/go/pathtree"
	"github.com/stretchr/testify/require"
)

func getNodes() (treeNodeIs pathtree.TreeNodeIs, err error) {
	data := `[
		{"id":"1","parentId":"0"},
		{"id":"2","parentId":"1"},
		{"id":"3","parentId":"2"},
		{"id":"4","parentId":"1"}
	]`
	simpleTrees := pathtree.SimpleTrees{}
	err = json.Unmarshal([]byte(data), &simpleTrees)
	if err != nil {
		return nil, err
	}
	simpleTrees.Init()
	treeNodes := simpleTrees.ConvertToTreeNodes()
	return treeNodes, nil
}

func TestResetAllPath(t *testing.T) {
	nodes, err := getNodes()
	require.NoError(t, err)
	err = nodes.ResetAllPath()
	require.NoError(t, err)
	simpleTrees := &pathtree.SimpleTrees{}
	nodes.Convert(simpleTrees)
	b, err := json.Marshal(simpleTrees)
	require.NoError(t, err)
	s := string(b)
	fmt.Println(s)
}
func TestFormatToTree(t *testing.T) {
	nodes, err := getNodes()
	require.NoError(t, err)

	err = nodes.ResetAllPath()
	require.NoError(t, err)

	nodes.CountChildren()
	tree := nodes.FormatToTree()

	simpleTrees := &pathtree.SimpleTrees{}
	tree.Convert(simpleTrees)
	b, err := json.Marshal(simpleTrees)
	require.NoError(t, err)
	s := string(b)
	fmt.Println(s)
}
