package dao

import (
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jmoiron/sqlx"
)

var MDB *sqlx.DB

/**
 *@Method 数据库初始化
 *@Params
 *@Return
 *@Tips:
 */
func Init(engine string, dsn string) (err error) {
	if dsn == ""{
		log.Printf("db init: %s, dns:%s", engine, dsn)
		return
	}

	MDB, err = sqlx.Connect(engine, dsn)
	if err != nil {
		log.Fatal("connect MDB failed:", err.Error())
		return
	}

	log.Printf("db type: %s, dns:%s", engine, dsn)

	// 设置最大连接数
	MDB.SetMaxOpenConns(20)
	MDB.SetMaxIdleConns(10)
	MDB.SetConnMaxLifetime(time.Minute * 5)
	return
}

/**
 *@Method 关闭连接
 *@Params
 *@Return
 *@Tips:
 */
func CloseMysql(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}
