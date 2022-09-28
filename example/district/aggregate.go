package district

type DistrictKV struct {
	Code       string `json:"code"`
	Title      string `json:"title"`
	ParentCode string `json:"parentcode"`
	Label      string `json:"label"`
	Depth      int    `json:"depth,string"`
	Path       string `json:"path"`
}
