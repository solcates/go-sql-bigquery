package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strings"
	"testing"
)

const (
	testTableName = "table1"
)

var testConn *conn
var testConfig *Config

func init() {
	var err error
	ctx := context.TODO()
	if os.Getenv(ConnectionStringEnvKey) != "" {
		testConnectionString = os.Getenv(ConnectionStringEnvKey)
	} else {
		testConnectionString = mockConnectString
	}
	if err = setupConnections(); err != nil {
		panic("Can not setup connections to Google Cloud... Check credentials, and connection string")
	}

	ds := testConn.client.Dataset(testConfig.DataSet)
	_, err = ds.Metadata(ctx)
	if err != nil {
		panic("Can not get dataset, check your connection string, permissions, and that it exists in your project")
	}
	// Check if the teable is there... if not let's create it
	t := ds.Table(testTableName)
	if _, err := t.Metadata(ctx); err != nil {
		// Table error
		if strings.HasSuffix(err.Error(), ", notFound") {
			// Need to create the table...
			err = t.Create(ctx, &bigquery.TableMetadata{
				Name:        testTableName,
				Description: "",
				Schema: bigquery.Schema{
					{
						Name: "name",
						Type: "STRING",
					}, {
						Name: "number",
						Type: "INT64",
					},
				},
			})
			if err != nil {
				panic(err)
			}

			// Add a single record for the test later
			q := testConn.client.Query("INSERT INTO dataset1.table1 (name, number) VALUES('hello',1);")
			_, err = q.Read(ctx)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}

	}
}

func setupConnections() (err error) {
	testConfig, err = ConfigFromConnString(testConnectionString)
	if err != nil {
		return
	}
	ctx := context.TODO()
	testConn, err = newConn(ctx, testConfig)
	if err != nil {
		return
	}
	return
}

func setupConnTests(t testing.TB) func(t testing.TB) {
	if err := setupConnections(); err != nil {
		t.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)
	// Check if the dataset and test table are live...
	return func(t testing.TB) {

	}
}

func Test_cfgFromConnString(t *testing.T) {
	teardown := setupConnTests(t)
	defer teardown(t)
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				in: "bigquery://projectid/us/dataset1",
			},

			wantCfg: &Config{
				Location:  "us",
				DataSet:   "dataset1",
				ProjectID: "projectid",
			},
			wantErr: false,
		},
		{
			name: "Bad Prefix",
			args: args{
				in: "bigquey://projectid/us/dataset1",
			},

			wantCfg: nil,
			wantErr: true,
		},
		{
			name: "Bad Connection String",
			args: args{
				in: "bigquery://projectid/us/dataset1/table",
			},

			wantCfg: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := ConfigFromConnString(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConfigFromConnString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("ConfigFromConnString() gotCfg = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func Test_conn_Ping(t *testing.T) {
	teardown := setupConnTests(t)
	defer teardown(t)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		dataset string
		args    args
		wantErr bool
	}{
		{
			name:    "OK",
			dataset: "dataset1",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testConn
			if err := c.Ping(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_conn_Query(t *testing.T) {
	teardown := setupConnTests(t)
	defer teardown(t)
	type fields struct {
		cfg       *Config
		client    *bigquery.Client
		ds        *bigquery.Dataset
		projectID string
		bad       bool
		closed    bool
		ctx       context.Context
	}
	type args struct {
		query string
		args  []driver.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantRows *bqRows
		wantErr  error
	}{
		{
			name: "SELECT *",
			args: args{
				query: "SELECT * FROM dataset1.table1;",
				args:  nil,
			},
			wantRows: &bqRows{
				columns: []string{"name", "number"},
				rs: resultSet{
					data: [][]bigquery.Value{
						{"hello", int64(1)},
					},
					num: 0,
				},
				c: testConn,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := testConn

			gotRows, err := c.Query(tt.args.query, tt.args.args)
			if err != nil {
				if tt.wantErr != nil {
					if tt.wantErr.Error() != err.Error() {
						t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					return
				} else {
					t.Errorf("Query() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(gotRows, tt.wantRows) {
				t.Errorf("Query() gotRows = %v, want %v", gotRows, tt.wantRows)
			}
		})
	}
}
