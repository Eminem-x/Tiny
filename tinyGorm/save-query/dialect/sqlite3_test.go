package dialect

import (
	"reflect"
	"testing"
)

func TestDataTypeOf(t *testing.T) {
	dial := &sqlite3{}
	cases := []struct {
		Value interface{}
		Type  string
	}{
		{"Ycx", "text"},
		{123, "integer"},
		{1.2, "real"},
		{[]int{1, 2, 3}, "blob"},
	}
	for i, c := range cases {
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Type {
			t.Fatalf("%d except %s, but got %s", i, c.Type, typ)
		}
	}
}
