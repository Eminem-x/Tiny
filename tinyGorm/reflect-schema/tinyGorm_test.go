package tinyGorm

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

// 1. 为适配不同的数据库，映射数据类型和特定的 SQL 语句，创建 Dialect 层屏蔽数据库差异
// 2. 设计 Schema，利用反射(reflect)完成结构体和数据库表结构的映射，包括表名、字段名、字段类型、字段 tag 等
// 3. 构造创建(create)、删除(drop)、存在性(table exists) 的 SQL 语句完成数据库表的基本操作

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite3", "tinyGorm.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}
