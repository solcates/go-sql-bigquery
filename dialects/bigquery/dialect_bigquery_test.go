package bigquery

import (
	"github.com/jinzhu/gorm"
	"testing"
)

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
	type fields struct {
		db                     gorm.SQLCommon
		DefaultForeignKeyNamer gorm.DefaultForeignKeyNamer
	}
	type args struct {
		tableName string
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
