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
}
