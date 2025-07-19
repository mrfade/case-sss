package errors

import errs "errors"

var (
	// Database errors
	ErrUnableToConnectDB = errs.New("unable to connect to the database")
	ErrNotFound          = errs.New("record not found")
	ErrConflictingData   = errs.New("conflicting data")

	// Redis errors
	ErrUnableToConnectRedis = errs.New("unable to connect to Redis")

	// JSON Provider errors
	ErrJSONProviderRequestFailed = errs.New("JSON provider failed to fetch data")
	ErrJSONProviderDecodeFailed  = errs.New("JSON provider failed to decode response")

	// XML Provider errors
	ErrXMLProviderRequestFailed = errs.New("XML provider failed to fetch data")
	ErrXMLProviderDecodeFailed  = errs.New("XML provider failed to decode response")
)
