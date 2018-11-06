package modules

import (
	"github.com/illidan33/sql-builder"
	"github.com/jmoiron/sqlx"
)

type Event struct {
	ID       uint   `json:"id" db:"id"`
	ParentID uint   `json:"parentID" db:"parent_id"`
	Name     string `json:"name" db:"name"`
	Desc     string `json:"desc" db:"desc"`
	Example  string `json:"example" db:"example"`
}

func GetWowEventByParentID(parentID int) []Event {
	conn := GetDbConn()

	builder := sql_builder.Select("api_event")
	builder.WhereEq("parent_id", parentID)

	wowApis := make([]Event, 0)
	sqlx.Select(conn, &wowApis, builder.String(), builder.Args()...)

	return wowApis
}
