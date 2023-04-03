package area

type AreaSQLGetAllByAreaIDListEntity struct {
	AreaIDList []string
}

func (t *AreaSQLGetAllByAreaIDListEntity) TplName() string {
	return "GetAllByAreaIDList"
}
func (t *AreaSQLGetAllByAreaIDListEntity) TplType() string {
	return "sql_select"
}

type AreaSQLGetByAreaIDEntity struct {
	AreaID string
}

func (t *AreaSQLGetByAreaIDEntity) TplName() string {
	return "GetByAreaID"
}
func (t *AreaSQLGetByAreaIDEntity) TplType() string {
	return "sql_select"
}

type AreaSQLGetByCityLevelEntity struct {
	CityLevel int
}

func (t *AreaSQLGetByCityLevelEntity) TplName() string {
	return "GetByCityLevel"
}
func (t *AreaSQLGetByCityLevelEntity) TplType() string {
	return "sql_select"
}

type AreaSQLGetByCityPathPrefixEntity struct {
	CityLevel  int
	PathPrefix string
}

func (t *AreaSQLGetByCityPathPrefixEntity) TplName() string {
	return "GetByCityPathPrefix"
}
func (t *AreaSQLGetByCityPathPrefixEntity) TplType() string {
	return "sql_select"
}

type AreaSQLListByKeywordEntity struct {
	AreaName string
}

func (t *AreaSQLListByKeywordEntity) TplName() string {
	return "ListByKeyword"
}
func (t *AreaSQLListByKeywordEntity) TplType() string {
	return "sql_select"
}
