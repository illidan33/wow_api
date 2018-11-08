package modules

import (
	"github.com/illidan33/sql-builder"
	"github.com/jmoiron/sqlx"
)

func GetApiByParentID(table string, parentID int) []ApiForGet {
	conn := GetDbConn()

	builder := sql_builder.Select(table)
	builder.WhereEq("parent_id", parentID)
	builder.WhereEq("enabled", 1)
	builder.SetSearchFields([]string{"id", "name", "name_cn", "desc"})

	wowApis := make([]ApiForGet, 0)
	sqlx.Select(conn, &wowApis, builder.String(), builder.Args()...)

	return wowApis
}
