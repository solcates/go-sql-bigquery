package bigquery

import (
	"fmt"
	"strings"
)

// ConfigFromConnString will return the Config structures
func ConfigFromConnString(in string) (*Config, error) {
	// Expects format to be bigquery://projectid/location/dataset   that's IT!
	// anything else will fail
	cfg := &Config{}
	if strings.HasPrefix(in, "bigquery://") {
		in = strings.ToLower(in)
		path := strings.TrimPrefix(in, "bigquery://")
		fields := strings.Split(path, "/")
		if len(fields) != 3 {
			return nil, fmt.Errorf("invalid connection string : %s", in)
		}
		cfg.ProjectID = fields[0]
		cfg.Location = fields[1]
		cfg.DataSet = fields[2]
		return cfg, nil
	} else {
		// Nope, bad prefix
		return nil, fmt.Errorf("invalid prefix, expected bigquery:// got: %s", in)
	}

}
