package area

import "github.com/suifengpiao14/gqt/v2"

type AreaSQLGetAllByAreaIDListEntity struct {
	AreaIDList []int
	gqt.TplEmptyEntity
}

func (t *AreaSQLGetAllByAreaIDListEntity) TplName() string {
	return "area.sql.GetAllByAreaIDList"
}
func (t *AreaSQLGetAllByAreaIDListEntity) TplType() string {
	return "sql_select"
}

type AreaSQLGetByAreaIDEntity struct {
	AreaID int
	gqt.TplEmptyEntity
}

func (t *AreaSQLGetByAreaIDEntity) TplName() string {
	return "area.sql.GetByAreaID"
}
func (t *AreaSQLGetByAreaIDEntity) TplType() string {
	return "sql_select"
}

type AreaSQLGetByCityLevelEntity struct {
	CityLevel int
	gqt.TplEmptyEntity
}

func (t *AreaSQLGetByCityLevelEntity) TplName() string {
	return "area.sql.GetByCityLevel"
}
func (t *AreaSQLGetByCityLevelEntity) TplType() string {
	return "sql_select"
}

type AreaSQLGetByCityPathPrefixEntity struct {
	CityLevel  int
	PathPrefix interface{}
	gqt.TplEmptyEntity
}

func (t *AreaSQLGetByCityPathPrefixEntity) TplName() string {
	return "area.sql.GetByCityPathPrefix"
}
func (t *AreaSQLGetByCityPathPrefixEntity) TplType() string {
	return "sql_select"
}

type AreaSQLListByKeywordEntity struct {
	CityLevel int
	gqt.TplEmptyEntity
}

func (t *AreaSQLListByKeywordEntity) TplName() string {
	return "area.sql.ListByKeyword"
}
func (t *AreaSQLListByKeywordEntity) TplType() string {
	return "sql_select"
}
