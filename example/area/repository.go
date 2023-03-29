package area

import (
	"context"
	"embed"
	"text/template"

	"gitea.programmerfamily.com/go/treeentity"
	"github.com/pkg/errors"
	"github.com/suifengpiao14/gotemplatefunc"
	"github.com/suifengpiao14/gotemplatefunc/templatedb"
	"github.com/suifengpiao14/gotemplatefunc/templatefunc"
	"github.com/suifengpiao14/gotemplatefunc/templateload"
	"github.com/suifengpiao14/gotemplatefunc/templatesql"
)

const (
	SQL_TPL_IDENTITY = "t_city_info"
)

//go:embed  template
var RepositoryFS embed.FS

func init() {
	cfg := templatedb.DBConfig{
		DSN:      "hjx:123456@tcp(recycle.m.mysql.hsb.com:3306)/recycle?charset=utf8&timeout=1s&readTimeout=5s&writeTimeout=5s&parseTime=False&loc=Local&multiStatements=true",
		LogLevel: "debug",
		Timeout:  5,
	}
	dbExecutorGetter := templatedb.NewExecutorGormGetter(cfg)
	t := template.New("").Funcs(templatefunc.TemplatefuncMapSQL)
	templateload.AddFromFS(t, RepositoryFS, "template/area.sql.tpl")
	err := templatesql.RegisterSQLTpl(SQL_TPL_IDENTITY, t, dbExecutorGetter)
	if err != nil {
		panic(err)
	}
}

type areaRecordRepository struct {
}

func (r *areaRecordRepository) AddNode(node treeentity.TreeNodeI) (err error) {
	err = errors.Errorf("Not implemented")
	return err
}
func (r *areaRecordRepository) UpdateNode(node treeentity.TreeNodeI) (err error) {
	err = errors.Errorf("Not implemented")
	return err
}
func (r *areaRecordRepository) UpdateBatchNode(nodes []treeentity.TreeNodeI) (err error) {
	err = errors.Errorf("Not implemented")
	return err
}
func (r *areaRecordRepository) GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error) {
	entity := AreaSQLGetByCityPathPrefixEntity{
		CityLevel:  LEVEL_CITY,
		PathPrefix: pathPrefix,
	}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, nodes)
	if err != nil {
		return err
	}
	return nil
}
func (r *areaRecordRepository) GetAllByNodeIds(nodeIds []string, nodes interface{}) (err error) {

	entity := AreaSQLGetAllByAreaIDListEntity{
		AreaIDList: nodeIds,
	}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, nodes)
	if err != nil {
		return err
	}
	return nil
}
func (r *areaRecordRepository) GetByAreaID(areaID string) (areaRecord *CityInfoModel, err error) {
	entity := AreaSQLGetByAreaIDEntity{
		AreaID: areaID,
	}
	areaRecord = &CityInfoModel{}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, areaRecord)
	if err != nil {
		return nil, err
	}
	return areaRecord, nil
}
func (r *areaRecordRepository) GetByLevel(depth string) (areaRecord CityInfoModels, err error) {
	entity := AreaSQLGetByCityLevelEntity{
		CityLevel: depth,
	}
	areaRecord = CityInfoModels{}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, &areaRecord)
	if err != nil {
		return nil, err
	}
	return areaRecord, nil

}
func (r *areaRecordRepository) GetByKeyWord(keyword string) (areaRecord CityInfoModels, err error) {
	entity := AreaSQLListByKeywordEntity{
		AreaName: keyword,
	}
	areaRecord = CityInfoModels{}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, &areaRecord)
	if err != nil {
		return nil, err
	}
	return areaRecord, nil
}
