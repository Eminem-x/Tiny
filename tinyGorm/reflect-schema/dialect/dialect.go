package dialect

import (
	"reflect"
)

// 不同数据库支持的数据类型是有差异的，即使功能相同，在 SQL 语句表达上也可能有差异
// ORM 框架往往需要兼容多种数据库，所以需要将差异部分抽取出来，实现解耦和复用

var dialectMap = map[string]Dialect{}

// Dialect is an interface contains methods that a dialect has no implement
type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    // 将 Go 语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{}) // 返回某个表是否存在
}

// RegisterDialect register a dialect to the global variable
func RegisterDialect(name string, dialect Dialect) {
	dialectMap[name] = dialect
}

// GetDialect get the dialect from global variable if it exists
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
