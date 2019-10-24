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
	rs      resultSet
	c       *conn
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
	dest[0] = b.rs.data[b.rs.num]
	b.rs.num++
	return nil
}
