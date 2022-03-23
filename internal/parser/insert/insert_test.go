package insert

import (
	"testing"
)

type createTableToken struct {
	table    string
	valueMap map[string]string
	isErr    bool
}

func TestGetTokenSuccess(t *testing.T) {
	m := map[string]createTableToken{
		"insert into test (id, name, value) values (1, 'test', 'value');": {"test", map[string]string{"id": "1", "name": "test", "value": "value"}, false},
		"insert into test (id, name, value) values (1, 'test');":          {"", nil, true},
	}

	for sql, expected := range m {
		table, valueMap, err := getToken(sql)

		if table != expected.table {
			t.Errorf("getToken(\"%s\"); got table: %s, expected %s", sql, table, expected.table)
		}

		if len(valueMap) != len(expected.valueMap) {
			t.Errorf("getToken(\"%s\"); got len(valueMap): %d, expected %d", sql, len(valueMap), len(expected.valueMap))
		}

		for key, value := range valueMap {
			if value != expected.valueMap[key] {
				t.Errorf("getToken(\"%s\"); got fields[%s]: %s, expected %s", sql, key, value, expected.valueMap[key])
			}
		}

		if (err == nil && expected.isErr == true) || (err != nil && expected.isErr == false) {
			t.Errorf("getToken(\"%s\"); got err: %s, expected %t", sql, err, expected.isErr)
		}
	}

}
