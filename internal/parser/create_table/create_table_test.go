package create_table

import "testing"

type createTableToken struct {
	isValid bool
	table   string
	fields  []string
}

func TestGetTokenSuccess(t *testing.T) {
	m := map[string]createTableToken{
		"create table test (id, value);": {true, "test", []string{"id", "value"}},
	}

	// TODO добавить сценариев

	for sql, expected := range m {
		isValid, table, fields := getToken(sql)
		if isValid != expected.isValid {
			t.Errorf("getToken(\"%s\"); got isValid: %t, expected %t", sql, isValid, expected.isValid)
		}
		if table != expected.table {
			t.Errorf("getToken(\"%s\"); got table: %s, expected %s", sql, table, expected.table)
		}
		if len(fields) != len(expected.fields) {
			t.Errorf("getToken(\"%s\"); got len(fields): %d, expected %d", sql, len(fields), len(expected.fields))
		}

		for i := range expected.fields {
			if fields[i] != expected.fields[i] {
				t.Errorf("getToken(\"%s\"); got fields[%d]: %s, expected %s", sql, i, fields[i], expected.fields[i])
			}
		}
	}
}
