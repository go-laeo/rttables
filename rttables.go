// Package rttables uses manipulate `/etc/iproute2/rt_tables`.
package rttables

import (
	"errors"
)

const (
	_builtinPath = "/etc/iproute2/rt_tables"
)

var (
	ErrMalformedFormat = errors.New("malformed format")
)

// Table is an item in rt_tables file.
type Table struct {
	ID       int
	Name     string
	reserved bool
	comment  bool
}
