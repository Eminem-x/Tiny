package schema

import (
	"go/ast"
	"reflect"
	"tinyGorm/dialect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema represents a table of database
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	filedMap   map[string]*Field // 记录 name 和 field 的映射关系，方便直接使用
}

func (schema *Schema) GetField(name string) *Field {
	return schema.filedMap[name]
}

func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// Indirect 获取指针指向的实例
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), // 获取结构体的名称作为表名
		filedMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) {
			filed := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("tGorm"); ok {
				filed.Tag = v
			}
			schema.Fields = append(schema.Fields, filed)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.filedMap[p.Name] = filed
		}
	}

	return schema
}
