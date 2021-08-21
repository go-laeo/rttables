// Package rttables uses manipulate `/etc/iproute2/rt_tables`.
package rttables

var (
	standardRtTables = []Table{
		{ID: -1, Name: "#", reserved: false, comment: true},
		{ID: -1, Name: "# reserved values", reserved: false, comment: true},
		{ID: -1, Name: "#", reserved: false, comment: true},
		{ID: 255, Name: "local", reserved: true, comment: false},
		{ID: 254, Name: "main", reserved: true, comment: false},
		{ID: 253, Name: "default", reserved: true, comment: false},
		{ID: 0, Name: "unspec", reserved: true, comment: false},
		{ID: -1, Name: "#", reserved: false, comment: true},
		{ID: -1, Name: "# local", reserved: false, comment: true},
		{ID: -1, Name: "#", reserved: false, comment: true},
		{ID: -1, Name: "#1	inr.ruhep", reserved: false, comment: true},
	}
)
