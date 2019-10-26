package bigquery

import (
	"database/sql"
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	testDialect *Dialect
	testDB      *sql.DB
	testDBMock  sqlmock.Sqlmock
	testRows    *sqlmock.Rows
)

type mockRows struct {
	mock.Mock
}

func (m mockRows) Columns() []string {
	panic("implement me")
}

func (m mockRows) Close() error {
	panic("implement me")
}

func (m mockRows) Next(dest []driver.Value) error {
	args := m.Called(dest)
	return args.Error(0)
}

type mockDB struct {
	mock.Mock
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	panic("implement me")
}

func (m *mockDB) Prepare(query string) (*sql.Stmt, error) {
	panic("implement me")
}

func (m *mockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {

	a := m.Called(query, args)
	return a.Get(0).(*sql.Rows), a.Error(1)
}

func (m *mockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	panic("implement me")
}

func setupDialectTests(t testing.TB) func(t testing.TB) {
	// Debug stuff
	logrus.SetLevel(logrus.DebugLevel)
	testDialect = new(Dialect)
	var err error
	testDB, testDBMock, err = sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	testDialect.db = testDB
	testDialect.dataset = "dataset1"
	return func(t testing.TB) {

	}
}

func TestDialect_BindVar(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "OK",
			fields: fields{},
			args:   args{},
			want:   "$$$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.BindVar(tt.args.i); got != tt.want {
				t.Errorf("BindVar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_CurrentDatabase(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.CurrentDatabase(); got != tt.want {
				t.Errorf("CurrentDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_DataTypeOf(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		field *gorm.StructField
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.DataTypeOf(tt.args.field); got != tt.want {
				t.Errorf("DataTypeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_DefaultValueStr(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.DefaultValueStr(); got != tt.want {
				t.Errorf("DefaultValueStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_GetName(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_HasColumn(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName  string
		columnName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.HasColumn(tt.args.tableName, tt.args.columnName); got != tt.want {
				t.Errorf("HasColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_HasForeignKey(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName      string
		foreignKeyName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.HasForeignKey(tt.args.tableName, tt.args.foreignKeyName); got != tt.want {
				t.Errorf("HasForeignKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_HasIndex(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName string
		indexName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.HasIndex(tt.args.tableName, tt.args.indexName); got != tt.want {
				t.Errorf("HasIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_HasTable(t *testing.T) {
	teardown := setupDialectTests(t)
	defer teardown(t)
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
		data                   []byte
		dataset                string
	}
	type args struct {
		tableName string
		query     string
		args      []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "OK, Simple",
			fields: fields{
				db:                     testDialect.db,
				DefaultForeignKeyNamer: testDialect.DefaultForeignKeyNamer,
				data:                   []byte("table1"),
				dataset:                "dataset1",
			},
			args: args{
				tableName: "table1",
				query:     "SELECT table_name FROM dataset1.INFORMATION_SCHEMA.TABLES where table_name = \"table1\"",
				args:      nil,
			},

			want: true,
		}, {
			name: "OK, Underscore, found",
			fields: fields{
				db:                     testDialect.db,
				DefaultForeignKeyNamer: testDialect.DefaultForeignKeyNamer,
				data:                   []byte("data_stores"),
				dataset:                "app_bigquery",
			},
			args: args{
				tableName: "data_stores",
				query:     "SELECT table_name FROM app_bigquery.INFORMATION_SCHEMA.TABLES where table_name = \"data_stores\"",
				args:      nil,
			},
			want: true,
		}, {
			name: "Error",
			fields: fields{
				db:                     testDialect.db,
				DefaultForeignKeyNamer: testDialect.DefaultForeignKeyNamer,
				data:                   nil,
				dataset:                "dataset1",

			},
			args: args{
				tableName: "table2",
				query:     "SELECT table_name FROM dataset1.INFORMATION_SCHEMA.TABLES where table_name = \"table2\"",
				args:      nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				dataset:                tt.fields.dataset,
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}

			rows := sqlmock.NewRows([]string{"table_name"}).
				AddRow(tt.fields.data)

			testDBMock.ExpectQuery(tt.args.query).WillReturnRows(rows)
			if got := b.HasTable(tt.args.tableName); got != tt.want {
				t.Errorf("HasTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_LastInsertIDReturningSuffix(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName  string
		columnName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.LastInsertIDReturningSuffix(tt.args.tableName, tt.args.columnName); got != tt.want {
				t.Errorf("LastInsertIDReturningSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_LimitAndOffsetSQL(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		limit  interface{}
		offset interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.LimitAndOffsetSQL(tt.args.limit, tt.args.offset); got != tt.want {
				t.Errorf("LimitAndOffsetSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_ModifyColumn(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName  string
		columnName string
		typ        string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if err := b.ModifyColumn(tt.args.tableName, tt.args.columnName, tt.args.typ); (err != nil) != tt.wantErr {
				t.Errorf("ModifyColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDialect_NormalizeIndexAndColumn(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		indexName  string
		columnName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			got, got1 := b.NormalizeIndexAndColumn(tt.args.indexName, tt.args.columnName)
			if got != tt.want {
				t.Errorf("NormalizeIndexAndColumn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NormalizeIndexAndColumn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDialect_Quote(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.Quote(tt.args.key); got != tt.want {
				t.Errorf("Quote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_RemoveIndex(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName string
		indexName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if err := b.RemoveIndex(tt.args.tableName, tt.args.indexName); (err != nil) != tt.wantErr {
				t.Errorf("RemoveIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDialect_SelectFromDummyTable(t *testing.T) {
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Dialect{
				db:                     tt.fields.db,
				DefaultForeignKeyNamer: tt.fields.DefaultForeignKeyNamer,
			}
			if got := b.SelectFromDummyTable(); got != tt.want {
				t.Errorf("SelectFromDummyTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
