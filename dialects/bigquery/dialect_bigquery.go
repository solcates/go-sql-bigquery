package bigquery

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	_ "github.com/solcates/go-sql-bigquery"
	bigquery "github.com/solcates/go-sql-bigquery"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Dialect struct {
	db gorm.SQLCommon
	gorm.DefaultForeignKeyNamer
}

func init() {
	gorm.RegisterDialect("bigquery", &Dialect{})

}

func (b *Dialect) GetName() string {
	return "bigquery"
}

func (b *Dialect) SetDB(db gorm.SQLCommon) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		uri := os.Getenv(bigquery.ConnectionStringEnvKey)
		if uri == "" {
			logrus.Panicf("no connection string found in environment... required currently, set %s", bigquery.ConnectionStringEnvKey)
		}
		var cfg *bigquery.Config
		cfg, err := bigquery.ConfigFromConnString(uri)
		if err != nil {
			logrus.Panic("invalid bigquery connection string should be like bigquery://projectid/us/somedataset")
		}
		defaultTableName = fmt.Sprintf("%s.%s", cfg.DataSet, defaultTableName)
		return defaultTableName
	}
	b.db = db
}

func (b *Dialect) BindVar(i int) string {
	return "$$$" // ?
}

func (b *Dialect) Quote(key string) string {
	return fmt.Sprintf("`%s`", key)
}

func (b *Dialect) DataTypeOf(field *gorm.StructField) string {
	var dataValue, sqlType, _, additionalType = gorm.ParseFieldStructForDialect(field, b)
	if sqlType == "" {
		switch dataValue.Kind() {
		case reflect.Bool:
			sqlType = "BOOL"
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
			if b.fieldCanAutoIncrement(field) {
				sqlType = "INT64 AUTO_INCREMENT"
			} else {
				sqlType = "INT64"
			}
		case reflect.Int64, reflect.Uint64:
			if b.fieldCanAutoIncrement(field) {
				sqlType = "INT64 AUTO_INCREMENT"
			} else {
				sqlType = "INT64"
			}
		case reflect.Float32, reflect.Float64:
			sqlType = "FLOAT64"
		case reflect.String:
			sqlType = "STRING"

		case reflect.Struct:
			if _, ok := dataValue.Interface().(time.Time); ok {
				sqlType = "TIMESTAMP"
			}
		default:
			if _, ok := dataValue.Interface().([]byte); ok {

				sqlType = "BYTES"

			}
		}
	}

	if sqlType == "" {
		panic(fmt.Sprintf("invalid sql type %s (%s) for commonDialect", dataValue.Type().Name(), dataValue.Kind().String()))
	}

	if strings.TrimSpace(additionalType) == "" {
		return sqlType
	}
	return fmt.Sprintf("%v %v", sqlType, additionalType)
}

func (b *Dialect) HasIndex(tableName string, indexName string) bool {
	panic("implement me")
}

func (b Dialect) HasForeignKey(tableName string, foreignKeyName string) bool {
	panic("implement me")
}

func (b Dialect) RemoveIndex(tableName string, indexName string) error {
	panic("implement me")
}

func (b *Dialect) HasTable(tableName string) bool {
	logrus.Debugf("Asking for Table: %s", tableName)
	ds := strings.Split(tableName, ".")
	if len(ds) != 2 {
		panic("got a bad tablename... shouldn't happen")
	}
	query := fmt.Sprintf("SELECT table_name FROM %s.INFORMATION_SCHEMA.TABLES where table_name = \"%s\"", ds[0], ds[1])
	rows, err := b.db.Query(query)
	if err != nil {
		panic(err)
	}
	type Tables struct {
		TableName string
	}
	if !rows.Next() {
		logrus.Debug("Did Not Find Table")
		return false
	} else {
		logrus.Debug("Found Table")
		return true
	}
	return false
}

func (b *Dialect) HasColumn(tableName string, columnName string) bool {
	query := fmt.Sprintf("SELECT * FROM %s LIMIT 0", tableName)
	rows, err := b.db.Query(query)
	if err != nil {
		return false
	}
	columns, err := rows.Columns()
	if err != nil {
		return false
	}
	for _, column := range columns {
		if column == columnName {
			return true
		}
	}
	return false
}

func (b Dialect) ModifyColumn(tableName string, columnName string, typ string) error {
	panic("implement me")
}

func (b Dialect) LimitAndOffsetSQL(limit, offset interface{}) (sql string) {
	if limit != nil {
		if parsedLimit, err := strconv.ParseInt(fmt.Sprint(limit), 0, 0); err == nil && parsedLimit >= 0 {
			sql += fmt.Sprintf(" LIMIT %d", parsedLimit)
		}
	}
	if offset != nil {
		if parsedOffset, err := strconv.ParseInt(fmt.Sprint(offset), 0, 0); err == nil && parsedOffset >= 0 {
			sql += fmt.Sprintf(" OFFSET %d", parsedOffset)
		}
	}
	return
}

func (b Dialect) SelectFromDummyTable() string {
	return ""
}

func (b Dialect) LastInsertIDReturningSuffix(tableName, columnName string) string {
	return ""
}

func (b Dialect) DefaultValueStr() string {
	return "DEFAULT VALUES"
}

func (b Dialect) NormalizeIndexAndColumn(indexName, columnName string) (string, string) {
	return indexName, columnName
}

func (b Dialect) CurrentDatabase() string {
	panic("implement me")
}

func (b *Dialect) fieldCanAutoIncrement(field *gorm.StructField) bool {
	if value, ok := field.TagSettingsGet("AUTO_INCREMENT"); ok {
		return strings.ToLower(value) != "false"
	}
	return field.IsPrimaryKey
}
