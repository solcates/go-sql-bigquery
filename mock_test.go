package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/stretchr/testify/mock"
	"google.golang.org/api/iterator"
)

type MockClient struct {
	mock.Mock
	*bigquery.Client
}

func (m *MockClient) Dataset(id string) *bigquery.Dataset {
	args := m.Called(id)
	return args.Get(0).(*bigquery.Dataset)
}

func (m *MockClient) Query(q string) *bigquery.Query {
	args := m.Called(q)
	return args.Get(0).(*bigquery.Query)
}

type MockDataset struct {
	mock.Mock
	*bigquery.Dataset
}

func (m MockDataset) Create(ctx context.Context, md *bigquery.DatasetMetadata) (err error) {
	panic("implement me")
}

func (m MockDataset) Delete(ctx context.Context) (err error) {
	panic("implement me")
}

func (m MockDataset) DeleteWithContents(ctx context.Context) (err error) {
	panic("implement me")
}

func (m MockDataset) Metadata(ctx context.Context) (md *bigquery.DatasetMetadata, err error) {
	args := m.Called(ctx)
	return args.Get(0).(*bigquery.DatasetMetadata), args.Error(1)
}

func (m MockDataset) Update(ctx context.Context, dm bigquery.DatasetMetadataToUpdate, etag string) (md *bigquery.DatasetMetadata, err error) {
	panic("implement me")
}

func (m MockDataset) Table(tableID string) *bigquery.Table {
	panic("implement me")
}

func (m MockDataset) Tables(ctx context.Context) *bigquery.TableIterator {
	panic("implement me")
}

func (m MockDataset) Model(modelID string) *bigquery.Model {
	panic("implement me")
}

func (m MockDataset) Models(ctx context.Context) *bigquery.ModelIterator {
	panic("implement me")
}

func (m MockDataset) Routine(routineID string) *bigquery.Routine {
	panic("implement me")
}

func (m MockDataset) Routines(ctx context.Context) *bigquery.RoutineIterator {
	panic("implement me")
}

type MockQuery struct {
	mock.Mock
	*bigquery.Query
}

func (m MockQuery) Run(ctx context.Context) (j *bigquery.Job, err error) {
	panic("implement me")
}

func (m MockQuery) Read(ctx context.Context) (*bigquery.RowIterator, error) {
	args := m.Called(ctx)
	return args.Get(0).(*bigquery.RowIterator), args.Error(1)
}

type MockRowIterator struct {
	mock.Mock
	*bigquery.RowIterator
}

func (m MockRowIterator) Next(dst interface{}) error {
	panic("implement me")
}

func (m MockRowIterator) PageInfo() *iterator.PageInfo {
	panic("implement me")
}
