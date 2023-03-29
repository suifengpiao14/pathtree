package area

type CityInfoModel struct {
	// 城市ID
	AreaID int `json:"areaID,string" gorm:"column:Farea_id"`
	// 上级城市ID
	ParentID int `json:"parentID,string" gorm:"column:Fparent_id"`
	// 城市名称
	AreaName string `json:"areaName" gorm:"column:Farea_name"`
	// 城市全称
	AllName string `json:"allName" gorm:"column:Fall_name"`
	// 名称首字母
	FirstLetter string `json:"firstLetter" gorm:"column:Ffirst_letter"`
	// 城市简拼
	CityJianpin string `json:"cityJianpin" gorm:"column:Fcity_jianpin"`
	// 城市全拼
	CityQp string `json:"cityQp" gorm:"column:Fcity_qp"`
	// 城市级别
	CityLevel int `json:"cityLevel,string" gorm:"column:Fcity_level"`
	// 城市路径
	CityPath string `json:"cityPath" gorm:"column:Fcity_path"`
	// 城市信息(备注)
	CityMsg string `json:"cityMsg" gorm:"column:Fcity_msg"`
	// 创建时间
	CreateTime string `json:"createTime" gorm:"column:Fcreate_time"`
	// 更新时间
	UpdateTime string `json:"updateTime" gorm:"column:Fupdate_time"`
	// 城市状态信息, 1:启用,2:停用,99:删除
	CityStatus int `json:"cityStatus,string" gorm:"column:Fcity_status"`
	// 省会id
	ProvinceID string `json:"provinceID" gorm:"column:Fprovince_id"`
	// 城市id
	CityID string `json:"cityID" gorm:"column:Fcity_id"`
	// 县id
	CountyID string `json:"countyID" gorm:"column:Fcounty_id"`
	// 变更时间
	AutoUpdateTime string `json:"autoUpdateTime" gorm:"column:Fauto_update_time"`
}

func (t *CityInfoModel) TableName() string {
	return "t_city_info"
}
func (t *CityInfoModel) PrimaryKey() string {
	return "Farea_id"
}
func (t *CityInfoModel) PrimaryKeyCamel() string {
	return "FareaID"
}
