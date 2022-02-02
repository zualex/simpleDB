package ddl

import (
	"testing"
)

func TestIsByDdlCommandSuccess(t *testing.T) {
	m := map[string]string{
		"CREATE TABLE test": CREATE_TABLE,
		"create table":      CREATE_TABLE,

		"DROP TABLE test": DROP_TABLE,
		"DROP TABLE":      DROP_TABLE,

		"insert": INSERT,
		"UpDate": UPDATE,
		"SELECT": SELECT,
		"delete": DELETE,
	}

	for sql, token := range m {
		got := isByDdlCommand(sql, token)
		if got != true {
			t.Errorf("isByDdlCommand(\"%s\") = %t; want true", sql, got)
		}
	}
}

func TestIsByDdlCommandFail(t *testing.T) {
	m := map[string]string{
		"CREATE":          CREATE_TABLE,
		"tt create table": CREATE_TABLE,
		"create tables":   CREATE_TABLE,

		"DROP":    DROP_TABLE,
		"inser":   INSERT,
		"UpDat":   UPDATE,
		"SELEC":   SELECT,
		"deleted": DELETE,
	}

	for sql, token := range m {
		got := isByDdlCommand(sql, token)
		if got != false {
			t.Errorf("isByDdlCommand(\"%s\") = %t; want true", sql, got)
		}
	}
}
