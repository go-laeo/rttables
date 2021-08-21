package rttables

// Contains check if t exists in /etc/iproute2/rt_tables.
func Contains(t Table) (bool, error) {
	return ContainsCustom(_builtinPath, t)
}

// ContainsCustom check t exists in file.
func ContainsCustom(file string, t Table) (bool, error) {
	rows, err := ParseCustom(file)
	if err != nil {
		return false, err
	}

	return ContainsIn(rows, t), nil
}

// ContainsIn check t exists in rows.
func ContainsIn(rows []Table, t Table) bool {
	for _, v := range rows {
		if v.ID == t.ID {
			return true
		}
	}

	return false
}
