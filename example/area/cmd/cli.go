package cmd

import (
	"gitea.programmerfamily.com/go/pathtree"
	"gitea.programmerfamily.com/go/pathtree/example/area"
)

func ResetAllPath() (err error) {
	areaRecord := &area.CityInfoModel{}
	r := areaRecord.GetRepository()
	all := make(area.CityInfoModels, 0)
	err = r.GetAllByPathPrefix("", -1, &all)
	if err != nil {
		return err
	}
	treeNodes := pathtree.ConvertToTreeNodes(all)
	err = treeNodes.ResetAllPath()
	if err != nil {
		return err
	}
	err = r.UpdateBatchNode(treeNodes)
	if err != nil {
		return err
	}
	return nil
}
