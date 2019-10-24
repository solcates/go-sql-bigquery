package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"database/sql/driver"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

const (
	mockProjectID = "mock-project"
)

var testDriver *Driver
var testConn *conn
var testClient *bigquery.Client
var testDataSet *bigquery.Dataset

func setupConnTests(t testing.TB) func(t testing.TB) {
	cfg, err := cfgFromConnString(testConnectionString)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.TODO()
	testConn, err = newConn(ctx, cfg)
	if err != nil {
		t.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
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
			gotCfg, err := cfgFromConnString(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("cfgFromConnString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("cfgFromConnString() gotCfg = %v, want %v", gotCfg, tt.wantCfg)
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
