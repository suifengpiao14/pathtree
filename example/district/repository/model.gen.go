package repository

const (
	DISTRICT_LABEL_COUNTRY   = "country"
	DISTRICT_LABEL_PROVINCE  = "province"
	DISTRICT_LABEL_CITY      = "city"
	DISTRICT_LABEL_AREA      = "area"
	DISTRICT_LABEL_STREET    = "street"
	DISTRICT_LABEL_COUNTY    = "county"
	DISTRICT_LABEL_TOWN      = "town"
	DISTRICT_LABEL_VILLAGE   = "village"
	DISTRICT_IS_DEPRECATED_0 = "0"
	DISTRICT_IS_DEPRECATED_1 = "1"
)

type DistrictModel struct {
	// 自增ID
	ID int `json:"id"`
	// 城市编码
	Code string `json:"code"`
	// 城市名称
	Title string `json:"title"`
	// 分级标签(country-国家,province-省,city-市,area-区,street-街道,county-县,town-镇,village-村)
	Label string `json:"label"`
	// 上级城市编码
	ParentCode string `json:"parentCode"`
	// 城市层级路径
	Path string `json:"path"`
	// 城市层级
	Depth int `json:"depth"`
	// 名称首字母
	FirstLetter string `json:"firstLetter"`
	// 是否废弃(0-弃用,1-废弃)
	IsDeprecated string `json:"isDeprecated"`
	// 创建时间
	CreatedAt string `json:"createdAt"`
	// 更新时间
	UpdatedAt string `json:"updatedAt"`
	// 删除时间
	DeletedAt string `json:"deletedAt"`
}

func (t *DistrictModel) TableName() string {
	return "district"
}
func (t *DistrictModel) PrimaryKey() string {
	return "id"
}
func (t *DistrictModel) PrimaryKeyCamel() string {
	return "ID"
}
func (t *DistrictModel) IsDeprecatedTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[DISTRICT_IS_DEPRECATED_0] = "弃用"
	enumMap[DISTRICT_IS_DEPRECATED_1] = "废弃"
	return enumMap
}
func (t *DistrictModel) IsDeprecatedTitle() string {
	enumMap := t.IsDeprecatedTitleMap()
	title, ok := enumMap[t.IsDeprecated]
	if !ok {
		msg := "func IsDeprecatedTitle not found title by key " + t.IsDeprecated
		panic(msg)
	}
	return title
}
func (t *DistrictModel) LabelTitleMap() map[string]string {
	enumMap := make(map[string]string)
	enumMap[DISTRICT_LABEL_COUNTRY] = "国家"
	enumMap[DISTRICT_LABEL_PROVINCE] = "省"
	enumMap[DISTRICT_LABEL_CITY] = "市"
	enumMap[DISTRICT_LABEL_AREA] = "区"
	enumMap[DISTRICT_LABEL_STREET] = "街道"
	enumMap[DISTRICT_LABEL_COUNTY] = "县"
	enumMap[DISTRICT_LABEL_TOWN] = "镇"
	enumMap[DISTRICT_LABEL_VILLAGE] = "村"
	return enumMap
}
func (t *DistrictModel) LabelTitle() string {
	enumMap := t.LabelTitleMap()
	title, ok := enumMap[t.Label]
	if !ok {
		msg := "func LabelTitle not found title by key " + t.Label
		panic(msg)
	}
	return title
}
