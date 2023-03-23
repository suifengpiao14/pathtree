package area

import "gitea.programmerfamily.com/go/treeentity"

func GetAll(parentID string) (nodes AreaRecords, err error) {
	areaRecord := &AreaRecord{}
	areaTree := treeentity.NewTree(areaRecord, areaRecord.GetRepository())
	nodes = make(AreaRecords, 0)
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
	areaRecord, err := (&AreaRecord{}).GetRepository().GetByAreaID(areaId)
	if err != nil {
		return nil, err
	}
	areaTree := treeentity.NewTree(areaRecord, areaRecord.GetRepository())
	parents := make(AreaRecords, 0)
	err = areaTree.GetParents(treeentity.DEPTH_MIN, true, &parents)
	if err != nil {
		return nil, err
	}
	//todo 格式化输出
	return out, nil
}

func GetCityInfo(typ string) (out interface{}, err error) {

	(&AreaRecord{}).GetRepository().GetByLevel(LEVEL_CITY)
	return out, nil
}

func GetCityListByKeyword(keyword string) (out interface{}, err error) {

	(&AreaRecord{}).GetRepository().GetByLevel(LEVEL_CITY)
	return out, nil
}

func GetCountiesByCityId(cityId string) (nodes AreaRecords, err error) {
	nodes = make(AreaRecords, 0)
	areaRecord := &AreaRecord{}
	areaTree := treeentity.NewTree(areaRecord, areaRecord.GetRepository())
	err = areaTree.GetChildren(1, false, &nodes)
	return nodes, nil
}
