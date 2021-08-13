package dao

import (
	"testing"
)

/**
 *@Method 数据库初始化
 *@Params
 *@Return
 *@Tips:
 */
func TestInit(t *testing.T){
	Init("sqlite3", ":memory:")
}

