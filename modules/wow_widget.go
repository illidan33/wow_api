package modules

import (
	"github.com/illidan33/sql-builder"
	"github.com/jmoiron/sqlx"
)

type Widget struct {
	ID       uint   `json:"id" db:"id"`
	ParentID uint   `json:"parentID" db:"parent_id"`
	Name     string `json:"name" db:"name"`
	Desc     string `json:"desc" db:"desc"`
	Example  string `json:"example" db:"example"`
}

func GetWowWidgetByParentID(parentID int) []Widget {
	conn := GetDbConn()

	builder := sql_builder.Select("api_widget")
	builder.WhereEq("parent_id", parentID)

	wowApis := make([]Widget, 0)
	sqlx.Select(conn, &wowApis, builder.String(), builder.Args()...)

	return wowApis
}
