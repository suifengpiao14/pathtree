package entity

//NodeModel 树结构模型
type NodeModel struct {
	NodeID    string `json:"nodeId"`
	Title     string `json:"title"`
	ParentID  string `json:"parentId"`
	IsLeaf    int    `json:"isLeaf"`
	Depth     int    `json:"depth"`
	Order     int    `json:"order"`
	Path      string `json:"path"`
	Ext       string `json:"ext"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
}
