package bigquery

import "testing"

func Test_result_LastInsertId(t *testing.T) {
	type fields struct {
		rowsAffected int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				rowsAffected: tt.fields.rowsAffected,
			}
			got, err := r.LastInsertId()
			if (err != nil) != tt.wantErr {
				t.Errorf("LastInsertId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LastInsertId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_result_RowsAffected(t *testing.T) {
	type fields struct {
		rowsAffected int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &result{
				rowsAffected: tt.fields.rowsAffected,
			}
			got, err := r.RowsAffected()
			if (err != nil) != tt.wantErr {
				t.Errorf("RowsAffected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RowsAffected() got = %v, want %v", got, tt.want)
			}
		})
	}
}