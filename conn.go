package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"reflect"
	"strings"
	"time"
)

type Dataset interface {
	// Create creates a dataset in the BigQuery service. An error will be returned if the
	// dataset already exists. Pass in a DatasetMetadata value to configure the dataset.
	Create(ctx context.Context, md *bigquery.DatasetMetadata) (err error)
	// Delete deletes the dataset.  Delete will fail if the dataset is not empty.
	Delete(ctx context.Context) (err error)
	// DeleteWithContents deletes the dataset, as well as contained resources.
	DeleteWithContents(ctx context.Context) (err error)
	// Metadata fetches the metadata for the dataset.
	Metadata(ctx context.Context) (md *bigquery.DatasetMetadata, err error)
	// Update modifies specific Dataset metadata fields.
	// To perform a read-modify-write that protects against intervening reads,
	// set the etag argument to the DatasetMetadata.ETag field from the read.
	// Pass the empty string for etag for a "blind write" that will always succeed.
	Update(ctx context.Context, dm bigquery.DatasetMetadataToUpdate, etag string) (md *bigquery.DatasetMetadata, err error)
	// Table creates a handle to a BigQuery table in the dataset.
	// To determine if a table exists, call Table.Metadata.
	// If the table does not already exist, use Table.Create to create it.
	Table(tableID string) *bigquery.Table
	// Tables returns an iterator over the tables in the Dataset.
	Tables(ctx context.Context) *bigquery.TableIterator
	// Model creates a handle to a BigQuery model in the dataset.
	// To determine if a model exists, call Model.Metadata.
	// If the model does not already exist, you can create it via execution
	// of a CREATE MODEL query.
	Model(modelID string) *bigquery.Model
	// Models returns an iterator over the models in the Dataset.
	Models(ctx context.Context) *bigquery.ModelIterator
	// Routine creates a handle to a BigQuery routine in the dataset.
	// To determine if a routine exists, call Routine.Metadata.
	Routine(routineID string) *bigquery.Routine
	// Routines returns an iterator over the routines in the Dataset.
	Routines(ctx context.Context) *bigquery.RoutineIterator
}

type Config struct {
	ProjectID string
	Location  string
	DataSet   string
}

type Conn struct {
	cfg       *Config
	client    *bigquery.Client
	ds        Dataset
	projectID string
	bad       bool
	closed    bool
}

func (c *Conn) prepareQuery(query string, args []driver.Value) (out string, err error) {
	if len(args) > 0 {
		logrus.Debugf("Preparing Query: %s ", query)

		for _, arg := range args {
			switch arg.(type) {
			case string:
				query = strings.Replace(query, "?", fmt.Sprintf("'%s'", arg), 1)
			case int, int64, int8, int32, int16:
				query = strings.Replace(query, "?", fmt.Sprintf("%d", arg), 1)
			case time.Time:
				t := arg.(time.Time)
				query = strings.Replace(query, "?", fmt.Sprintf("'%s'", t.Format("2006-01-02 15:04:05")), 1)
			case []byte:
				data, ok := arg.([]byte)
				if ok {
					if len(data) == 0 {
						query = strings.Replace(query, "?", "NULL", 1)

					} else {
						newdata := base64.StdEncoding.EncodeToString(data)
						query = strings.Replace(query, "?", fmt.Sprintf("FROM_BASE64('%s')", newdata), 1)
					}
				}

			default:
				logrus.Debugf("unknown type: %s", reflect.TypeOf(arg).String())
				query = strings.Replace(query, "?", fmt.Sprintf("'%s'", arg), 1)
			}

		}
		logrus.Debugf("Prepared Query: %s ", query)
		out = query

	} else {
		out = query
	}
	return
}

func (c *Conn) Exec(query string, args []driver.Value) (res driver.Result, err error) {
	logrus.Debugf("Got Conn.Exec: %s", query)
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
	logrus.Debugf("Results for Conn.Exec: %s", data)

	return
}

// NewConn returns a connection for this Config
func NewConn(ctx context.Context, cfg *Config) (c *Conn, err error) {
	c = &Conn{
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
	return NewConn(ctx, cfg)
}

func (c *Connector) Driver() driver.Driver {
	return &Driver{}
}

//Ping the BigQuery service and make sure it's reachable
func (c *Conn) Ping(ctx context.Context) (err error) {
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

func (c *Conn) Query(query string, args []driver.Value) (rows driver.Rows, err error) {
	// This is a HACK for the mocking that we have to do as the google cloud package doesn't include/use interfaces
	// TODO: Come back if we ever can avoid the Interface hack...
	logrus.Debugf("Got Conn.Query: %s", query)
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
func (c *Conn) Prepare(query string) (stmt driver.Stmt, err error) {
	stmt = NewStmt(query, c)
	return
}

//Begin  is stubbed out and not used
func (c *Conn) Begin() (driver.Tx, error) {
	return newTx(c)
}

//Close closes the connection
func (c *Conn) Close() (err error) {
	if c.closed {
		return nil
	}
	if c.bad {
		return driver.ErrBadConn
	}
	c.closed = true
	return c.client.Close()
}
