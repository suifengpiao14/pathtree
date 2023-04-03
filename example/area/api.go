package area

import treeentity "gitea.programmerfamily.com/go/pathtree"

func GetAll(parentID string) (nodes CityInfoModels, err error) {
	areaRecord := &CityInfoModel{}
	areaTree := treeentity.NewTree(areaRecord, areaRecord.GetRepository())
	nodes = make(CityInfoModels, 0)
	depth := -1
	//传值: 0 表示省份和直瞎市, 传父ID区取当前子级; 不传则获所有,同时返回children
	switch parentID {
	case "0":
		depth = 2
	case "":
		depth = -1
	default:
		depth = 1
	}
	err = areaTree.GetChildren(depth, true, &nodes)
	if err != nil {
		return nil, err
	}
	//todo 转成树状结构
	return nodes, nil
}

type GetParentAreaOut struct {
	Level        string `json:"level"`
	ProvinceName string `json:"provinceName"`
	ProvinceCode string `json:"provinceCode"`
	CityName     string `json:"cityName"`
	CityCode     string `json:"cityCode"`
	AreaName     string `json:"areaName"`
	AreaCode     string `json:"areaCode"`
}

func GetParentArea(areaId string) (out *GetParentAreaOut, err error) {
	areaRecord, err := (&CityInfoModel{}).GetRepository().GetByAreaID(areaId)
	if err != nil {
		return nil, err
	}
	areaTree := treeentity.NewTree(areaRecord, areaRecord.GetRepository())
	parents := make(CityInfoModels, 0)
	err = areaTree.GetParents(treeentity.DEPTH_MIN, true, &parents)
	if err != nil {
		return nil, err
	}
	//todo 格式化输出
	return out, nil
}

func GetCityInfo(typ string) (out interface{}, err error) {

	(&CityInfoModel{}).GetRepository().GetByLevel(LEVEL_CITY)
	return out, nil
}

func GetCityListByKeyword(keyword string) (out interface{}, err error) {

	(&CityInfoModel{}).GetRepository().GetByLevel(LEVEL_CITY)
	return out, nil
}

func GetCountiesByCityId(cityId string) (nodes CityInfoModels, err error) {
	nodes = make(CityInfoModels, 0)
	areaRecord := &CityInfoModel{}
	areaTree := treeentity.NewTree(areaRecord, areaRecord.GetRepository())
	err = areaTree.GetChildren(1, false, &nodes)
	return nodes, nil
}
