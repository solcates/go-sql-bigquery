package bigquery

type result struct {
	rowsAffected int64
}

func (r *result) LastInsertId() (int64, error) {
	panic("implement me")
}

func (r *result) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}
