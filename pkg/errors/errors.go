package errors

import errs "errors"

var (
	// Database errors
	ErrUnableToConnectDB = errs.New("unable to connect to the database")

	// Redis errors
	ErrUnableToConnectRedis = errs.New("unable to connect to Redis")
)
