package pathtree

type SimpleTree struct {
	ID       string `json:"id"`
	ParentID string `json:"parentId"`
	Path     string `json:"path"`
	Depth    int    `json:"depth,string"`
	_set     []*SimpleTree
	EmptyTreeNode
	Childern         []*SimpleTree `json:"childern"`
	DirectChildCount int           `json:"directChildCount,string"`
	AllChildCount    int           `json:"allChildCount,string"`
}

func (st *SimpleTree) GetNodeID() (nodeID string) {
	return st.ID
}
func (st *SimpleTree) GetParentID() (parentID string) {
	return st.ParentID
}
func (st *SimpleTree) SetPath(path string) {
	st.Path = path
}
func (st *SimpleTree) SetDepth(depth int) {
	st.Depth = depth

}
func (st *SimpleTree) GetParent() (parent TreeNodeI, err error) {

	for _, node := range st._set {
		if node.ID == st.ParentID {
			return node, nil
		}
	}
	err = ERROR_NODE_NOT_FOUND
	return nil, err
}
func (st *SimpleTree) GetPath() (path string) {
	path = st.Path
	return path
}
func (st *SimpleTree) GetDepth() (depth int) {
	depth = st.Depth
	return depth
}
func (st *SimpleTree) SetParentID(parentId string) {
	st.ParentID = parentId
}

func (st *SimpleTree) IsRoot() (ok bool) {
	return st.ParentID == "" || st.ParentID == "0"
}

func (st *SimpleTree) AddChildren(node TreeNodeI) {
	if st.Childern == nil {
		st.Childern = make([]*SimpleTree, 0)
	}
	simpleTree := node.(*SimpleTree)
	st.Childern = append(st.Childern, simpleTree)
}
func (st *SimpleTree) IncrChildrenCount(causeNode TreeNodeI) {
	causeSimpleTree := causeNode.(*SimpleTree)
	if st.ID == causeSimpleTree.ParentID {
		st.DirectChildCount++
	}
	st.AllChildCount++
	parent, _ := st.GetParent()
	if parent != nil { // 实现统计所有子节点
		parent.IncrChildrenCount(causeNode)
	}
}

type SimpleTrees []*SimpleTree

func (sts *SimpleTrees) Init() {
	for _, simpleNode := range *sts {
		simpleNode._set = *sts
		simpleNode.Childern = make([]*SimpleTree, 0)
	}
}
