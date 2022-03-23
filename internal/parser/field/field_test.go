package field

import "testing"

func TestGetFieldsFromStringSuccess(t *testing.T) {
	fields := GetFieldsFromString("id,field")
	if len(fields) != 2 {
		t.Errorf("GetFieldsFromString(\"id,field\") length: %d", len(fields))
	}
	if fields[0] != "id" {
		t.Errorf("GetFieldsFromString(\"id,field\") fields[0]: %s, expected: id", fields[0])
	}
	if fields[1] != "field" {
		t.Errorf("GetFieldsFromString(\"id,field\") fields[1]: %s, expected: field", fields[1])
	}

	fields = GetFieldsFromString("id,field,")
	if len(fields) != 2 {
		t.Errorf("GetFieldsFromString(\"id,field,\") length: %d", len(fields))
	}

	fields = GetFieldsFromString("id,field,,")
	if len(fields) != 2 {
		t.Errorf("GetFieldsFromString(\"id,field,,\") length: %d", len(fields))
	}

	fields = GetFieldsFromString("1,'value', \"test\"")
	if len(fields) != 3 {
		t.Errorf("GetFieldsFromString(\"1,'value', \"test\"\") length: %d", len(fields))
	}
	if fields[0] != "1" {
		t.Errorf("GetFieldsFromString(\"1,'value', \"test\"\") field[0]: %s, expected %s", fields[0], "1")
	}
	if fields[1] != "value" {
		t.Errorf("GetFieldsFromString(\"1,'value', \"test\"\") field[0]: %s, expected %s", fields[1], "value")
	}
	if fields[2] != "test" {
		t.Errorf("GetFieldsFromString(\"1,'value', \"test\"\") field[0]: %s, expected %s", fields[2], "test")
	}
}
