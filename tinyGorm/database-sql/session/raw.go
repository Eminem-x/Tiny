package session

import (
	"database/sql"
	"strings"
	"tinyGorm/log"
)

// 封装 Session 是为了多次复用，开启一次会话，可以执行多次 SQL

// Session keep a pointer to sql.DB and provides all execution of
// all kind of database operations
type Session struct {
	db      *sql.DB
	sql     strings.Builder // 拼接 sql 语句
	sqlVars []interface{}   // sql 语句中占位符的对应值
}

// New creates an instance of Session
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear initialize the state of a session
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB returns *sql.DB
func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	result, err = s.DB().Exec(s.sql.String(), s.sqlVars...) // ignore_security_alert
	if err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...) // ignore_security_alert
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	rows, err = s.DB().Query(s.sql.String(), s.sqlVars...) // ignore_security_alert
	if err != nil {
		log.Error(err)
	}
	return
}

// Raw appends sql and sqlVars
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}
