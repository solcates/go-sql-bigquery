package bigquery

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

const (
	ConnectionStringEnvKey = "BIGQUERY_CONNECTION_STRING"
)

type Driver struct {
	ctx    context.Context
	Config *Config
}

func (d *Driver) Open(connectionString string) (c driver.Conn, err error) {
	//ctx := context.TODO()
	if d.Config, err = ConfigFromConnString(connectionString); err != nil {
		return
	}

	return NewConnector(connectionString).Connect(context.TODO())
}

func init() {
	sql.Register("bigquery", &Driver{})
}
