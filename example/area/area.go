package area

import (
	"context"
	"fmt"
	"strconv"

	"gitea.programmerfamily.com/go/pathtree"
	"github.com/suifengpiao14/gotemplatefunc/templatedb"
	"github.com/suifengpiao14/logchan/v2"
)

func init() {
	logchan.SetLoggerWriter(context.Background(), loggerFn)
}

func loggerFn(logInfo logchan.LogInforInterface, typeName string, err error) {

	if err != nil {
		fmt.Println(logInfo)
		return
	}
	switch typeName {
	case templatedb.LOG_INFO_EXEC_SQL:
		sqlLog := logInfo.(*templatedb.LogInfoEXECSQL)
		fmt.Println(sqlLog.SQL)
	}

}

type AreaRecordRepository interface {
	pathtree.TreeRepositoryI
	GetByAreaID(areaID string) (areaRecord *CityInfoModel, err error)
	GetByLevel(depth int) (areaRecord CityInfoModels, err error)
	GetByKeyWord(keyword string, depth string) (areaRecord CityInfoModels, err error)
	UpdatePathAndDepth(areaRecord *CityInfoModel) (err error)
}

const (
	LEVEL_CITY int = 3
)

type CityInfoModels []*CityInfoModel

func (node *CityInfoModel) GetNodeID() (nodeID string) {
	nodeID = strconv.Itoa(node.AreaID)
	return nodeID
}
func (node *CityInfoModel) SetPath(path string) {
	node.CityPath = path

}
func (node *CityInfoModel) GetPath() (path string) {
	return node.CityPath
}
func (node *CityInfoModel) SetDepth(depth int) {
	node.CityLevel = depth

}
func (node *CityInfoModel) GetDepth() (depth int) {
	return node.CityLevel
}

func (node *CityInfoModel) SetParentID(parentId string) {
	node.ParentID, _ = strconv.Atoi(parentId)

}
func (node *CityInfoModel) GetParent() (parent pathtree.TreeNodeI, err error) {
	parentArea, err := node.GetRepository().GetByAreaID(strconv.Itoa(node.ParentID))
	if err != nil {
		return nil, err
	}
	return parentArea, nil
}

func (node *CityInfoModel) GetRepository() (r AreaRecordRepository) {
	r = &areaRecordRepository{}
	return r
}
func (node *CityInfoModel) IsRoot() (ok bool) {
	return node.AreaID == 1
}
