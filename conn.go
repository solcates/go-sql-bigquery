package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"strings"
	"time"
)

type Config struct {
	ProjectID string
	Location  string
	DataSet   string
}

type conn struct {
	cfg       *Config
	client    *bigquery.Client
	ds        *bigquery.Dataset
	projectID string
	bad       bool
	closed    bool
}

func (c *conn) prepareQuery(query string, args []driver.Value) (out string, err error) {
	if len(args) > 0 {
		//logrus.Debugf("Preparing Query: %s ", query)

		for _, arg := range args {
			switch arg.(type) {
			case string:
				query = strings.Replace(query, "?", fmt.Sprintf("'%s'", arg), 1)
			case int, int64, int8, int32, int16:
				query = strings.Replace(query, "?", fmt.Sprintf("%d", arg), 1)
			case time.Time:
				t := arg.(time.Time)

				query = strings.Replace(query, "?", fmt.Sprintf("'%s'", t.Format("2006-01-02 15:04:05")), 1)

			default:
				query = strings.Replace(query, "?", fmt.Sprintf("'%s'", arg), 1)
			}

		}
		//logrus.Debugf("Prepared Query: %s ", query)
		out = query

	} else {
		out = query
	}
	return
}

func (c *conn) Exec(query string, args []driver.Value) (res driver.Result, err error) {
	logrus.Debugf("Got conn.Exec: %s", query)
	if query, err = c.prepareQuery(query, args); err != nil {
		return
	}
	ctx := context.TODO()
	q := c.client.Query(query)
	it, err := q.Read(ctx)
	if err != nil {
		return
	}
	var data [][]bigquery.Value
	for {
		var row []bigquery.Value
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		data = append(data, row)
	}
	res = &result{
		rowsAffected: int64(it.TotalRows),
	}
	logrus.Debugf("Results for conn.Exec: %s", data)

	return
}

// newConn returns a connection for this Config
func newConn(ctx context.Context, cfg *Config) (c *conn, err error) {
	c = &conn{
		cfg: cfg,
	}
	c.client, err = bigquery.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		return nil, err
	}
	c.ds = c.client.Dataset(c.cfg.DataSet)

	return
}

type Connector struct {
	Info             map[string]string
	Client           *bigquery.Client
	connectionString string
}

func NewConnector(connectionString string) *Connector {
	return &Connector{connectionString: connectionString}
}

func (c *Connector) Connect(ctx context.Context) (driver.Conn, error) {
	cfg, err := ConfigFromConnString(c.connectionString)
	if err != nil {
		return nil, err
	}
	return newConn(ctx, cfg)
}

func (c *Connector) Driver() driver.Driver {
	return &Driver{}
}

//Ping the BigQuery service and make sure it's reachable
func (c *conn) Ping(ctx context.Context) (err error) {
	if c.ds == nil {
		c.ds = c.client.Dataset(c.cfg.DataSet)
	}
	var md *bigquery.DatasetMetadata
	md, err = c.ds.Metadata(ctx)
	if err != nil {
		logrus.Debugf("Failed Ping Dataset: %s", c.cfg.DataSet)
		return
	}
	logrus.Debugf("Successful Ping: %s", md.FullID)
	return
}

func (c *conn) Query(query string, args []driver.Value) (rows driver.Rows, err error) {
	// This is a HACK for the mocking that we have to do as the google cloud package doesn't include/use interfaces
	// TODO: Come back if we ever can avoid the Interface hack...
	logrus.Debugf("Got conn.Query: %s", query)
	q := c.client.Query(query)
	ctx := context.TODO()
	var rowsIterator *bigquery.RowIterator
	rowsIterator, err = q.Read(ctx)
	if err != nil {
		return
	}

	bqrows := &bqRows{
		columns: nil,
		rs:      resultSet{},
		c:       c,
	}

	for {
		var row []bigquery.Value
		err := rowsIterator.Next(&row)
		if bqrows.columns == nil {

			for _, column := range rowsIterator.Schema {
				bqrows.columns = append(bqrows.columns, column.Name)
			}
		}
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		bqrows.rs.data = append(bqrows.rs.data, row)
	}
	rows = bqrows
	return
}

// Prepare is stubbed out and not used
func (c *conn) Prepare(query string) (stmt driver.Stmt, err error) {
	stmt = NewStmt(query, c)
	return
}

//Begin  is stubbed out and not used
func (c *conn) Begin() (driver.Tx, error) {
	return newTx(c)
}

//Close closes the connection
func (c *conn) Close() (err error) {
	if c.closed {
		return nil
	}
	if c.bad {
		return driver.ErrBadConn
	}
	c.closed = true
	return c.client.Close()
}
