package cmd

import (
	"gitea.programmerfamily.com/go/pathtree"
	"gitea.programmerfamily.com/go/pathtree/example/area"
)

type CityInfoModel struct {
	area.CityInfoModel
	_set *CityInfoModels
}

func (node *CityInfoModel) GetParent() (parent pathtree.TreeNodeI, err error) {

	for _, record := range *node._set {
		if record.AreaID == node.ParentID {
			return record, nil
		}
	}
	return nil, nil
}

type CityInfoModels []*CityInfoModel

func (cis *CityInfoModels) Init() {
	for _, cityInfo := range *cis {
		cityInfo._set = cis
	}
}

func ResetAllPath() (err error) {
	areaRecord := &CityInfoModel{}

	all := make(CityInfoModels, 0)
	err = areaRecord.GetAllByPathPrefix("", -1, &all)
	if err != nil {
		return err
	}
	all.Init()
	treeNodes := pathtree.ConvertToTreeNodes(all)
	err = treeNodes.ResetAllPath()
	if err != nil {
		return err
	}
	r := areaRecord.GetRepository()
	for _, record := range all {
		err = r.UpdatePathAndDepth(&record.CityInfoModel)
		if err != nil {
			return err
		}
	}

	return nil
}
