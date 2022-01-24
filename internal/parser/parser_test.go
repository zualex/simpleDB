package parser

import (
	"testing"
)

func TestIsCreateTableSuccess(t *testing.T) {
	got := IsCreateTable("CREATE TABLE")
	if got != true {
		t.Errorf("IsCreateTable(\"CREATE TABLE\") = %t; want true", got)
	}

	got = IsCreateTable("create TABLE")
	if got != true {
		t.Errorf("IsCreateTable(\"create TABLE\") = %t; want true", got)
	}

	got = IsCreateTable("create table test")
	if got != true {
		t.Errorf("IsCreateTable(\"create table test\") = %t; want true", got)
	}
}

func TestIsCreateTableFail(t *testing.T) {
	got := IsCreateTable("CREATE")
	if got != false {
		t.Errorf("IsCreateTable(\"CREATE\") = %t; want false", got)
	}

	got = IsCreateTable("t CREATE TABLE")
	if got != false {
		t.Errorf("IsCreateTable(\"t CREATE TABLE\") = %t; want false", got)
	}
}

func TestIsDropTableSuccess(t *testing.T) {
	got := IsDropTable("DROP TABLE")
	if got != true {
		t.Errorf("IsDropTable(\"DROP TABLE\") = %t; want true", got)
	}

	got = IsDropTable("drop TABLE")
	if got != true {
		t.Errorf("IsDropTable(\"drop TABLE\") = %t; want true", got)
	}

	got = IsDropTable("drop table test")
	if got != true {
		t.Errorf("IsDropTable(\"drop table test\") = %t; want true", got)
	}
}

func TestIsDropTableFail(t *testing.T) {
	got := IsDropTable("DROP")
	if got != false {
		t.Errorf("IsDropTable(\"DROP\") = %t; want false", got)
	}

	got = IsDropTable("t DROP TABLE")
	if got != false {
		t.Errorf("IsDropTable(\"t DROP TABLE\") = %t; want false", got)
	}
}

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

func TestIsSelectSuccess(t *testing.T) {
	got := IsSelect("SELECT")
	if got != true {
		t.Errorf("IsSelect(\"SELECT\") = %t; want true", got)
	}

	got = IsSelect("select")
	if got != true {
		t.Errorf("IsSelect(\"select\") = %t; want true", got)
	}

	got = IsSelect("select into table")
	if got != true {
		t.Errorf("IsSelect(\"select into table\") = %t; want true", got)
	}
}

func TestIsSelectFail(t *testing.T) {
	got := IsSelect("SELEC")
	if got != false {
		t.Errorf("IsSelect(\"SELEC\") = %t; want false", got)
	}

	got = IsSelect("t SELECT")
	if got != false {
		t.Errorf("IsSelect(\"t SELECT\") = %t; want false", got)
	}
}

func TestIsUpdateSuccess(t *testing.T) {
	got := IsUpdate("UPDATE")
	if got != true {
		t.Errorf("IsUpdate(\"UPDATE\") = %t; want true", got)
	}

	got = IsUpdate("update")
	if got != true {
		t.Errorf("IsUpdate(\"update\") = %t; want true", got)
	}

	got = IsUpdate("update into table")
	if got != true {
		t.Errorf("IsUpdate(\"update into table\") = %t; want true", got)
	}
}

func TestIsUpdateFail(t *testing.T) {
	got := IsUpdate("UPDA")
	if got != false {
		t.Errorf("IsUpdate(\"UPDA\") = %t; want false", got)
	}

	got = IsUpdate("t UPDATE")
	if got != false {
		t.Errorf("IsUpdate(\"t UPDATE\") = %t; want false", got)
	}
}
