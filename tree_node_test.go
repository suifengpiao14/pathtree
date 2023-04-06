package pathtree

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func getNodes() (treeNodeIs treeNodeIs, err error) {
	data := `[
		{"id":"1","parentId":"0"},
		{"id":"2","parentId":"1"},
		{"id":"3","parentId":"2"},
		{"id":"4","parentId":"1"}
	]`
	simpleTrees := SimpleTrees{}
	err = json.Unmarshal([]byte(data), &simpleTrees)
	if err != nil {
		return nil, err
	}
	simpleTrees.Init()
	treeNodes := ConvertToTreeNodes(simpleTrees)
	return treeNodes, nil
}

func TestResetAllPath(t *testing.T) {
	nodes, err := getNodes()
	require.NoError(t, err)
	err = nodes.ResetAllPath()
	require.NoError(t, err)
	simpleTrees := &SimpleTrees{}
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

	simpleTrees := &SimpleTrees{}
	tree.Convert(simpleTrees)
	b, err := json.Marshal(simpleTrees)
	require.NoError(t, err)
	s := string(b)
	fmt.Println(s)
}
