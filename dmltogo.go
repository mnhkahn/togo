package dmltogo

import (
	"bytes"
	"database/sql"
	"fmt"
	"html/template"
	"reflect"
	"strings"

	"bou.ke/monkey"
	"github.com/jimsmart/schema"
	"github.com/jinzhu/inflection"
	"github.com/serenize/snaker"
	"github.com/smallnest/gen/dbmeta"
	gtmpl "github.com/smallnest/gen/template"
	"github.com/xwb1989/sqlparser"
)

// DmlToGo uses create sql to generate golang struct with gorm tags.
func DmlToGo(dml string) (string, error) {
	stmt, err := sqlparser.Parse(dml)
	if err != nil {
		return "", err
	}

	var ddl *sqlparser.DDL
	switch stmt := stmt.(type) {
	case *sqlparser.DDL:
		ddl = stmt
	default:
		return "", fmt.Errorf("not support type")
	}

	tableName := dbmeta.FmtFieldName(ddl.NewName.Name.String())
	structName := inflection.Singular(tableName)

	// Use monkey to hack the schema.Table function that skip the *sql.DB object.
	monkey.Patch(schema.Table, func(db *sql.DB, name string) ([]*sql.ColumnType, error) {
		cols := make([]*sql.ColumnType, len(ddl.TableSpec.Columns))
		for i, _ := range ddl.TableSpec.Columns {
			cols[i] = &sql.ColumnType{}
		}
		return cols, nil
	})
	columnPatch := new(sql.ColumnType)
	i := 0
	monkey.PatchInstanceMethod(reflect.TypeOf(columnPatch), "Name", func(columnPatch *sql.ColumnType) string {
		return ddl.TableSpec.Columns[i].Name.String()
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(columnPatch), "Nullable", func(columnPatch *sql.ColumnType) (nullable, ok bool) {
		//nullable = bool(ddl.TableSpec.Columns[i].Type.NotNull)
		return false, true
	})
	monkey.PatchInstanceMethod(reflect.TypeOf(columnPatch), "DatabaseTypeName", func(columnPatch *sql.ColumnType) string {
		mysqlType := strings.ToUpper(ddl.TableSpec.Columns[i].Type.Type)
		// DatabaseTypeName called at last, so i++ is here
		i++
		return mysqlType
	})
	defer monkey.UnpatchAll()

	modelInfo := dbmeta.GenerateStruct(nil, tableName, structName, "model", true, true, true)

	var buf bytes.Buffer
	t, err := getTemplate(gtmpl.ModelTmpl)
	if err != nil {
		return "", err
	}
	err = t.Execute(&buf, modelInfo)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// port from github.com/smallnest/gen
func getTemplate(t string) (*template.Template, error) {
	var funcMap = template.FuncMap{
		"pluralize":        inflection.Plural,
		"title":            strings.Title,
		"toLower":          strings.ToLower,
		"toLowerCamelCase": camelToLowerCamel,
		"toSnakeCase":      snaker.CamelToSnake,
	}

	tmpl, err := template.New("model").Funcs(funcMap).Parse(t)

	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

// port from github.com/smallnest/gen
func camelToLowerCamel(s string) string {
	ss := strings.Split(s, "")
	ss[0] = strings.ToLower(ss[0])

	return strings.Join(ss, "")
}
