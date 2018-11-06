package modules

import (
	"github.com/illidan33/sql-builder"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	ID       uint   `json:"id" db:"id"`
	ParentID uint   `json:"parentID" db:"parent_id"`
	Name     string `json:"name" db:"name"`
	Desc     string `json:"desc" db:"desc"`
	Example  string `json:"example" db:"example"`
}

func GetWowApiByParentID(parentID int) []Api {
	conn := GetDbConn()

	builder := sql_builder.Select("api_wow")
	builder.WhereEq("parent_id", parentID)

	wowApis := make([]Api, 0)
	sqlx.Select(conn, &wowApis, builder.String(), builder.Args()...)

	return wowApis
}
