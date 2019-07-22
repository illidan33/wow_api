package modules

import (
	"github.com/illidan33/sql-builder"
	"github.com/illidan33/wow_hong_golang/modules"
	"github.com/jmoiron/sqlx"
)

func GetApiByParentID(table string, parentID int) []ApiForGet {
	conn := modules.DbConn

	builder := sql_builder.Select(table)
	builder.WhereEq("parent_id", parentID)
	builder.WhereEq("enabled", 1)
	builder.SetSearchFields([]string{"id", "name", "name_cn", "desc"})

	wowApis := make([]ApiForGet, 0)
	sqlx.Select(conn, &wowApis, builder.String(), builder.Args()...)

	return wowApis
}

func GetApiByID(table string, ID int) Api {
	conn := modules.DbConn

	builder := sql_builder.Select(table)
	builder.WhereEq("id", ID)

	wowApi := Api{}
	sqlx.Get(conn, &wowApi, builder.String(), builder.Args()...)

	return wowApi
}

func SaveApiUnverify(api UnVerifyApi) error {
	conn := modules.DbConn

	builder := sql_builder.Insert("api_unverify")
	builder.InsertByStruct(api)

	_, err := conn.Exec(builder.String(), builder.Args()...)
	if err != nil {
		return err
	}

	return nil
}
