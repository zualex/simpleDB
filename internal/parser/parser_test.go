package parser

import (
	"testing"
)

func TestIsInsertSuccess(t *testing.T) {
	got := IsInsert("INSERT")
	if got != true {
		t.Errorf("IsInsert(\"INSERT\") = %t; want true", got)
	}

	got = IsInsert("insert")
	if got != true {
		t.Errorf("IsInsert(\"insert\") = %t; want true", got)
	}

	got = IsInsert("insert into table")
	if got != true {
		t.Errorf("IsInsert(\"insert into table\") = %t; want true", got)
	}
}

func TestIsInsertFail(t *testing.T) {
	got := IsInsert("INSER")
	if got != false {
		t.Errorf("IsInsert(\"INSER\") = %t; want false", got)
	}

	got = IsInsert("t INSERT")
	if got != false {
		t.Errorf("IsInsert(\"t INSERT\") = %t; want false", got)
	}
}
