package bigquery

import (
	"fmt"
	"strings"
)

// ConfigFromConnString will return the Config structures
func ConfigFromConnString(in string) (cfg *Config, err error) {
	// Expects format to be bigquery://projectid/location/dataset   that's IT!
	// anything else will fail
	cfg = &Config{}
	if strings.HasPrefix(in, "bigquery://") {
		in = strings.ToLower(in)
		path := strings.TrimPrefix(in, "bigquery://")
		fields := strings.Split(path, "/")
		if len(fields) != 3 {
			err = fmt.Errorf("invalid connection string : %s", in)
			return
		}
		cfg.ProjectID = fields[0]
		cfg.Location = fields[1]
		cfg.DataSet = fields[2]
		return
	} else {
		// Nope, bad prefix
		err = fmt.Errorf("invalid prefix, expected bigquery:// got: %s", in)
		return
	}

}
