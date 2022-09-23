package repository

import "github.com/suifengpiao14/gqt/v2/gqttpl"

type GenDistrictSQLInsertEntity struct {
	Code         string
	Depth        int
	FirstLetter  string
	IsDeprecated string
	Label        string
	ParentCode   string
	Path         string
	Title        string
	gqttpl.TplEmptyEntity
}

func (t *GenDistrictSQLInsertEntity) TplName() string {
	return "gen.district.sql.Insert"
}
func (t *GenDistrictSQLInsertEntity) TplType() string {
	return "sql_insert"
}

type GenDistrictSQLPaginateEntity struct {
	Limit  int
	Offset int
	GenDistrictSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *GenDistrictSQLPaginateEntity) TplName() string {
	return "gen.district.sql.Paginate"
}
func (t *GenDistrictSQLPaginateEntity) TplType() string {
	return "sql_select"
}

type GenDistrictSQLPaginateTotalEntity struct {
	GenDistrictSQLPaginateWhereEntity
	gqttpl.TplEmptyEntity
}

func (t *GenDistrictSQLPaginateTotalEntity) TplName() string {
	return "gen.district.sql.PaginateTotal"
}
func (t *GenDistrictSQLPaginateTotalEntity) TplType() string {
	return "sql_select"
}

type GenDistrictSQLPaginateWhereEntity struct {
	gqttpl.TplEmptyEntity
}

func (t *GenDistrictSQLPaginateWhereEntity) TplName() string {
	return "gen.district.sql.PaginateWhere"
}
func (t *GenDistrictSQLPaginateWhereEntity) TplType() string {
	return "text"
}

type GenDistrictSQLUpdateEntity struct {
	Code         string
	Depth        int
	FirstLetter  string
	ID           int
	IsDeprecated string
	Label        string
	ParentCode   string
	Path         string
	Title        string
	gqttpl.TplEmptyEntity
}

func (t *GenDistrictSQLUpdateEntity) TplName() string {
	return "gen.district.sql.Update"
}
func (t *GenDistrictSQLUpdateEntity) TplType() string {
	return "sql_update"
}
