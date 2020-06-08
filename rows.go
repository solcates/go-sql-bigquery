package bigquery

import (
	"cloud.google.com/go/bigquery"
	"database/sql/driver"
	"io"
)

type resultSet struct {
	data [][]bigquery.Value
	num  int
}

type bqRows struct {
	columns []string
	types   []string
	rs      resultSet
	c       *Conn
}

func (b *bqRows) Columns() []string {
	return b.columns
}

func (b *bqRows) Close() error {
	return b.c.Close()
}

func (b *bqRows) Next(dest []driver.Value) error {
	if b.rs.num == len(b.rs.data) {
		return io.EOF
	}
	for i, bgValue := range b.rs.data[b.rs.num] {
		dest[i] = bgValue
	}
	b.rs.num++
	return nil
}

func (b *bqRows) ColumnTypeDatabaseTypeName(index int) string {
	return b.types[index]
}