package bigquery

import (
	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"database/sql/driver"
	"fmt"
	"io"
	"math/big"
	"strings"
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
		switch bgValue.(type) {
		// FIXME: For types that cannot be converted by sql.convertAssignRows function successfully,
		// we transform these values into string temporary.
		case *big.Rat, []bigquery.Value, civil.Time, civil.Date, civil.DateTime:
			dest[i] = bqValueToString(bgValue)
			break
		default:
			dest[i] = bgValue
			break
		}
	}
	b.rs.num++
	return nil
}

func bqValueToString(bgValue bigquery.Value) string {
	switch bgValue.(type) {
	// For type BIGNUMERIC
	case *big.Rat:
		return bigquery.BigNumericString(bgValue.(*big.Rat))
	// For type ARRAY and STRUCT
	case []bigquery.Value:
		strSlice := make([]string, 0, len(bgValue.([]bigquery.Value)))
		values := bgValue.([]bigquery.Value)
		for _, value := range values {
			strSlice = append(strSlice, bqValueToString(value))
		}
		return strings.Join(strSlice, "")
	// For Time
	case civil.Time:
		return bgValue.(civil.Time).String()
	// For Date
	case civil.Date:
		return bgValue.(civil.Date).String()
	// For DateTime
	case civil.DateTime:
		return bgValue.(civil.DateTime).String()
	}
	return fmt.Sprintf("%+v", bgValue)
}

func (b *bqRows) ColumnTypeDatabaseTypeName(index int) string {
	return b.types[index]
}
