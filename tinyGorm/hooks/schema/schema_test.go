package schema

import (
	"testing"
	"tinyGorm/dialect"
)

var TestDial, _ = dialect.GetDialect("sqlite3")

type UserTest struct {
	Name string `tinyGorm:"PRIMARY KEY"`
	Age  int    `tinyGorm:"int"`
}

func (u *UserTest) TableName() string {
	return "ns_user_test"
}

func TestSchema_TableName(t *testing.T) {
	schema := Parse(&UserTest{}, TestDial)
	if schema.Name != "ns_user_test" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}
	for _, field := range schema.Fields {
		if field.Name == "Name" && field.Tag != "PRIMARY KEY" {
			t.Fatalf("tag error: %#v", field.Tag)
		}
		if field.Name == "Age" && field.Tag != "int" {
			t.Fatalf("tag error: %#v", field.Tag)
		}
	}
}

func TestSchema_RecordValues(t *testing.T) {
	schema := Parse(&UserTest{}, TestDial)
	values := schema.RecordValues(&UserTest{"Tom", 18})

	name := values[0].(string)
	age := values[1].(int)

	if name != "Tom" || age != 18 {
		t.Fatal("failed to get values")
	}
}
