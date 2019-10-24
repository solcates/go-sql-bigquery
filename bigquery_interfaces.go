package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"google.golang.org/api/iterator"
)

// This file is a HACK for the mocking that we have to do as the Google Cloud package doesn't include/use interfaces

type BigQueryClient interface {
	// Dataset creates a handle to a BigQuery dataset in the client's project.
	Dataset(id string) *bigquery.Dataset
	// DatasetInProject creates a handle to a BigQuery dataset in the specified project.
	DatasetInProject(projectID, datasetID string) *bigquery.Dataset
	// Datasets returns an iterator over the datasets in a project.
	// The Client's project is used by default, but that can be
	// changed by setting ProjectID on the returned iterator before calling Next.
	Datasets(ctx context.Context) *bigquery.DatasetIterator
	// DatasetsInProject returns an iterator over the datasets in the provided project.
	//
	// Deprecated: call Client.Datasets, then set ProjectID on the returned iterator.
	DatasetsInProject(ctx context.Context, projectID string) *bigquery.DatasetIterator
	// Close closes any resources held by the client.
	// Close should be called when the client is no longer needed.
	// It need not be called at program exit.
	Close() error
	// Query creates a query with string q.
	// The returned Query may optionally be further configured before its Run method is called.
	Query(q string) *bigquery.Query
	// JobFromID creates a Job which refers to an existing BigQuery job. The job
	// need not have been created by this package. For example, the job may have
	// been created in the BigQuery console.
	//
	// For jobs whose location is other than "US" or "EU", set Client.Location or use
	// JobFromIDLocation.
	JobFromID(ctx context.Context, id string) (*bigquery.Job, error)
	// JobFromIDLocation creates a Job which refers to an existing BigQuery job. The job
	// need not have been created by this package (for example, it may have
	// been created in the BigQuery console), but it must exist in the specified location.
	JobFromIDLocation(ctx context.Context, id, location string) (j *bigquery.Job, err error)
	// Jobs lists jobs within a project.
	Jobs(ctx context.Context) *bigquery.JobIterator
}

type BigQueryDataset interface {
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

// BigQueryTable ...
type BigQueryTable interface {
	// FullyQualifiedName returns the ID of the table in projectID:datasetID.tableID format.
	FullyQualifiedName() string
	// Create creates a table in the BigQuery service.
	// Pass in a TableMetadata value to configure the table.
	// If tm.View.Query is non-empty, the created table will be of type VIEW.
	// If no ExpirationTime is specified, the table will never expire.
	// After table creation, a view can be modified only if its table was initially created
	// with a view.
	Create(ctx context.Context, tm *bigquery.TableMetadata) (err error)
	// Metadata fetches the metadata for the table.
	Metadata(ctx context.Context) (md *bigquery.TableMetadata, err error)
	// Delete deletes the table.
	Delete(ctx context.Context) (err error)
	// Read fetches the contents of the table.
	Read(ctx context.Context) *bigquery.RowIterator
	// Update modifies specific Table metadata fields.
	Update(ctx context.Context, tm bigquery.TableMetadataToUpdate, etag string) (md *bigquery.TableMetadata, err error)
}

// BigQueryQuery ...
type BigQueryQuery interface {
	// Run initiates a query job.
	Run(ctx context.Context) (j *bigquery.Job, err error)
	// Read submits a query for execution and returns the results via a RowIterator.
	// It is a shorthand for Query.Run followed by Job.Read.
	Read(ctx context.Context) (*bigquery.RowIterator, error)
}

// BigQueryRowIterator ...
type BigQueryRowIterator interface {
	// Next loads the next row into dst. Its return value is iterator.Done if there
	// are no more results. Once Next returns iterator.Done, all subsequent calls
	// will return iterator.Done.
	//
	// dst may implement ValueLoader, or may be a *[]Value, *map[string]Value, or struct pointer.
	//
	// If dst is a *[]Value, it will be set to new []Value whose i'th element
	// will be populated with the i'th column of the row.
	//
	// If dst is a *map[string]Value, a new map will be created if dst is nil. Then
	// for each schema column name, the map key of that name will be set to the column's
	// value. STRUCT types (RECORD types or nested schemas) become nested maps.
	//
	// If dst is pointer to a struct, each column in the schema will be matched
	// with an exported field of the struct that has the same name, ignoring case.
	// Unmatched schema columns and struct fields will be ignored.
	//
	// Each BigQuery column type corresponds to one or more Go types; a matching struct
	// field must be of the correct type. The correspondences are:
	//
	//   STRING      string
	//   BOOL        bool
	//   INTEGER     int, int8, int16, int32, int64, uint8, uint16, uint32
	//   FLOAT       float32, float64
	//   BYTES       []byte
	//   TIMESTAMP   time.Time
	//   DATE        civil.Date
	//   TIME        civil.Time
	//   DATETIME    civil.DateTime
	//
	// A repeated field corresponds to a slice or array of the element type. A STRUCT
	// type (RECORD or nested schema) corresponds to a nested struct or struct pointer.
	// All calls to Next on the same iterator must use the same struct type.
	//
	// It is an error to attempt to read a BigQuery NULL value into a struct field,
	// unless the field is of type []byte or is one of the special Null types: NullInt64,
	// NullFloat64, NullBool, NullString, NullTimestamp, NullDate, NullTime or
	// NullDateTime. You can also use a *[]Value or *map[string]Value to read from a
	// table with NULLs.
	Next(dst interface{}) error
	// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
	PageInfo() *iterator.PageInfo
}
