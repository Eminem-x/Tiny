package tinyGorm

import (
	"database/sql"
	"tinyGorm/log"
	"tinyGorm/session"
)

// Engine 是用户与 ORM 交互的入口，通过其实例创建具体的 Session
// 负责交互前的准备（比如连接/测试数据库），交互后的收尾（关闭连接）

// Engine is the main struct of tinyGorm, manages all db sessions and transactions.
type Engine struct {
	db *sql.DB
}

// NewEngine create an instance of Engine
// connect database and ping it to test whether it's alive
func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	// send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

// Close database connection
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

// NewSession creates a new session for next operations
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db)
}
