// Package store provides a simple and convenient data store interface for plugins to persist
// data along with a default filed-based leveldb implementation.
package store

import (
	"io"
)

// StringStorer is implemented by any value that has the Get/Put/Scan and Closer methods
// on string keys/values.
type StringStorer interface {
	io.Closer
	Scanner

	GetString(key string) (value string, err error)
	PutString(key string, value string) (err error)
}

// BytesStorer is implemented by any value that has the Get/Put/Scan and Closer methods
// on byte arrays for keys/values.
type BytesStorer interface {
	io.Closer
	Scanner

	Get(key []byte) (value []byte, err error)
	Put(key []byte, value []byte) (err error)
}

// Scanner is implemented by any value that has the Scan method returning
// all key/values as strings.
//
// Since []byte aren't allowed as map keys, we use don't have an equivalent
// Scanner interface returning the data as bytes.
// For implementers, it should be easy to convert the []byte values
// to string with a simple string(value) conversion as strings
// are immutable arrays of bytes
type Scanner interface {
	Scan() (entries map[string]string, err error)
}
