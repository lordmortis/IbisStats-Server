package datasource

import (
	"errors"
)

var ErrUUIDParse = errors.New("unable to parse UUID")
var ErrNoResults = errors.New("no results from query")
var ErrNoDb = errors.New("no database connection")