package ddltogo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/smallnest/gen/dbmeta"
	"github.com/xwb1989/sqlparser"
	"golang.org/x/tools/imports"
)

// DdlToGo uses create sql to generate golang struct with gorm tags.
func DdlToGo(sql string) ([]byte, error) {
	stmt, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		return nil, fmt.Errorf("parse ddl error: %s", err.Error())
	}

	var ddl *sqlparser.DDL
	switch stmt := stmt.(type) {
	case *sqlparser.DDL:
		ddl = stmt
	default:
		return nil, fmt.Errorf("not support type")
	}

	tableName := ddl.NewName.Name.String()
	structName := dbmeta.FmtFieldName(tableName)

	columns := make([]map[string]string, 0, len(ddl.TableSpec.Columns))
	if ddl.TableSpec != nil {
		for i, col := range ddl.TableSpec.Columns {
			name := col.Name.String()
			nullable := strconv.FormatBool(bool(ddl.TableSpec.Columns[i].Type.NotNull))
			mysqlType := strings.ToLower(ddl.TableSpec.Columns[i].Type.Type)
			columns = append(columns, map[string]string{"name": name, "value": mysqlType, "nullable": nullable})
		}
	}

	orm, err := Generate(columns, tableName, structName, "model", true, true, true)
	if err != nil {
		return nil, fmt.Errorf("generate error: %s", err.Error())
	}

	orm, err = imports.Process("", orm, nil)
	if err != nil {
		return nil, fmt.Errorf("imports process error: %s", err.Error())
	}

	return orm, nil
}
