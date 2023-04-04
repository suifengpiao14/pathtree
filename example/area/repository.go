package area

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"strconv"
	"text/template"

	"gitea.programmerfamily.com/go/pathtree"
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
		DSN:         "hjx:123456@tcp(recycle.m.mysql.hsb.com:3306)/recycle?charset=utf8&timeout=1s&readTimeout=5s&writeTimeout=5s&parseTime=False&loc=Local&multiStatements=true",
		LogLevel:    "debug",
		Timeout:     5,
		MaxOpen:     1,
		MaxIdle:     1,
		MaxIdleTime: 60,
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
	pathtree.EmptyTreeRpository
}

func (r *areaRecordRepository) GetAllByPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error) {
	if pathPrefix != "" {
		pathPrefix = fmt.Sprintf("%%%s%%", pathPrefix)
	}
	entity := AreaSQLGetAllByPathPrefixEntity{
		PathPrefix: pathPrefix,
	}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, nodes)
	if err != nil {
		return err
	}
	return nil
}
func (r *areaRecordRepository) GetByCityPathPrefix(pathPrefix string, depth int, nodes interface{}) (err error) {
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
	if errors.Is(err, templatedb.ERROR_DB_RECORD_NOT_FOUND) {
		err = nil
	}
	if err != nil {
		return nil, err
	}
	return areaRecord, nil
}
func (r *areaRecordRepository) GetByLevel(depth int) (areaRecord CityInfoModels, err error) {
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
func (r *areaRecordRepository) GetByKeyWord(keyword string, depth string) (areaRecord CityInfoModels, err error) {
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
func (r *areaRecordRepository) UpdatePathAndDepth(cityInfo *CityInfoModel) (err error) {
	entity := AreaSQLUpdatePathAndDepthEntity{
		CityLevel: strconv.Itoa(cityInfo.CityLevel),
		CityPath:  cityInfo.CityPath,
		AreaID:    strconv.Itoa(cityInfo.AreaID),
	}
	err = gotemplatefunc.ExecSQLTpl(context.Background(), SQL_TPL_IDENTITY, entity.TplName(), &entity, nil)
	if err != nil {
		return err
	}
	return nil
}
