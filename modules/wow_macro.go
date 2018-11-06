package modules

import (
	"github.com/illidan33/sql-builder"
	"github.com/jmoiron/sqlx"
)

type Macro struct {
	ID       uint   `json:"id" db:"id"`
	ParentID uint   `json:"parentID" db:"parent_id"`
	Name     string `json:"name" db:"name"`
	Desc     string `json:"desc" db:"desc"`
	Example  string `json:"example" db:"example"`
}

func GetWowMacroByParentID(parentID int) []Macro {
	conn := GetDbConn()

	builder := sql_builder.Select("api_macro")
	builder.WhereEq("parent_id", parentID)

	wowApis := make([]Macro, 0)
	sqlx.Select(conn, &wowApis, builder.String(), builder.Args()...)

	return wowApis
}
