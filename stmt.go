package bigquery

import (
	"database/sql/driver"
	"github.com/sirupsen/logrus"
)

type stmt struct {
	query string
	c     *conn
}

func NewStmt(query string, c *conn) *stmt {
	return &stmt{query: query, c: c}
}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {

	return 0
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	logrus.Debugf("Got stmt.Exec: %s", s.query)
	return s.c.Exec(s.query, args)
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	logrus.Debugf("Got stmt.Query: %s", s.query)
	return s.c.Query(s.query, args)
}
