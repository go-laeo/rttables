package rttables

import (
	"bytes"
	"os"
	"strconv"
)

// Parse convert from /etc/iproute2/rt_tables to []Table.
func Parse() ([]Table, error) {
	return ParseCustom(_builtinPath)
}

// ParseCustom convert file to []Table.
func ParseCustom(file string) ([]Table, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	v := make([]Table, 0)
	series := bytes.Split(b, []byte("\n"))
	for _, seg := range series {
		if len(seg) == 0 || seg[0] == '#' {
			v = append(v, Table{
				ID:       -1,
				Name:     string(seg),
				reserved: false,
				comment:  true,
			})
			continue
		}

		m := bytes.SplitN(seg, []byte("\t"), 2)
		if len(m) != 2 {
			return nil, ErrMalformedFormat
		}

		id, err := strconv.Atoi(string(bytes.TrimSpace(m[0])))
		if err != nil {
			return nil, err
		}

		name := string(bytes.TrimSpace(m[1]))

		v = append(v, Table{
			ID:       id,
			Name:     name,
			reserved: id == 0 || id == 255 || id == 254 || id == 253,
			comment:  false,
		})
	}

	return v, nil
}
