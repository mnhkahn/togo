package dmltogo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Shelnutt2/db2struct"
	"github.com/jinzhu/inflection"
	"github.com/xwb1989/sqlparser"
	"golang.org/x/tools/imports"
)

// DmlToGo uses create sql to generate golang struct with gorm tags.
func DmlToGo(dml string) ([]byte, error) {
	stmt, err := sqlparser.Parse(dml)
	if err != nil {
		return nil, err
	}

	var ddl *sqlparser.DDL
	switch stmt := stmt.(type) {
	case *sqlparser.DDL:
		ddl = stmt
	default:
		return nil, fmt.Errorf("not support type")
	}

	tableName := ddl.NewName.Name.String()
	structName := inflection.Singular(tableName)

	columns := make(map[string]map[string]string)
	if ddl.TableSpec != nil {
		for i, col := range ddl.TableSpec.Columns {
			name := col.Name.String()
			nullable := strconv.FormatBool(bool(ddl.TableSpec.Columns[i].Type.NotNull))
			mysqlType := strings.ToLower(ddl.TableSpec.Columns[i].Type.Type)
			columns[name] = map[string]string{"value": mysqlType, "nullable": nullable}
		}
	}

	orm, err := db2struct.Generate(columns, tableName, structName, "model", true, true, true)
	if err != nil {
		return nil, err
	}

	orm, err = imports.Process("", orm, nil)
	if err != nil {
		return nil, err
	}

	return orm, nil
}
