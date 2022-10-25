package common

import "errors"

var (
	RecordNotFound = errors.New("mongo: no documents in result")
)
