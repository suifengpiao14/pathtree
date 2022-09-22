package district

type DistrictKV struct {
	Code     string `json:"code"`
	Title    string `json:"title"`
	NodeID   string `json:"nodeId"`
	ParentID string `json:"parentId"`
	Label    string `json:"label"`
	Depth    int    `json:"depth,string"`
	Path     string `json:"path"`
}
