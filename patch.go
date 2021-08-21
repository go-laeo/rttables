package rttables

import (
	"io"
	"os"
	"strconv"
)

// Patch convert rows to plaintext, and write to /etc/iproute2/rt_tables.
func Patch(rows []Table) error {
	return PatchCustom(_builtinPath, rows)
}

// PatchCustom convert rows to plaintext, then write to file.
func PatchCustom(file string, rows []Table) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	return PatchWriter(f, rows)
}

// PatchWriter convert rows to plaintext, then write to w.
func PatchWriter(w io.Writer, rows []Table) (err error) {
	for i, v := range rows {
		if i != 0 {
			_, err = io.WriteString(w, "\n")
			if err != nil {
				return
			}
		}

		if v.comment {
			_, err = io.WriteString(w, v.Name)
			if err != nil {
				return
			}
			continue
		}

		_, err = io.WriteString(w, strconv.Itoa(v.ID))
		if err != nil {
			return
		}
		_, err = io.WriteString(w, "\t")
		if err != nil {
			return
		}
		_, err = io.WriteString(w, v.Name)
		if err != nil {
			return
		}
	}
	return
}
