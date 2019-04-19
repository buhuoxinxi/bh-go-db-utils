package bhdbutils

import (
	"github.com/jinzhu/gorm"
	"os"
	"testing"
)

func TestNewDBConn(t *testing.T) {
	// option
	os.Setenv("BhDatabaseOrmDebug", "true")
	os.Setenv("BhDatabaseTablePrefix", "ikg_")
	os.Setenv("BhDatabaseMaxOpenConn", "10")
	os.Setenv("BhDatabaseMaxIdleConn", "10")
	os.Setenv("BhDatabaseConnMaxLifetime", "30s")

	var db *gorm.DB

	// mysql
	t.Logf("test mysql ... \n")
	// auth
	// root:Mysql.123456@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local
	os.Setenv("BhDatabaseDriver", "mysql")
	os.Setenv("BhDatabaseUsername", "root")
	os.Setenv("BhDatabasePassword", "Mysql.123456")
	os.Setenv("BhDatabaseHost", "127.0.0.1")
	os.Setenv("BhDatabasePort", "3306")
	os.Setenv("BhDatabaseDBName", "test")
	os.Setenv("BhDatabaseParameters", "charset=utf8&timeout=60s&loc=Local&autocommit=true")
	db = NewDBConn()
	t.Logf("%v \n", db)

	// postgres
	t.Logf("test postgres ... \n")
	// auth
	// host=myhost port=myport user=gorm dbname=gorm password=mypassword
	// postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
	os.Setenv("BhDatabaseDriver", "postgres")
	os.Setenv("BhDatabaseUsername", "postgres")
	os.Setenv("BhDatabasePassword", "Postgres.123456")
	os.Setenv("BhDatabaseHost", "127.0.0.1")
	os.Setenv("BhDatabasePort", "5432")
	os.Setenv("BhDatabaseDBName", "postgres")
	//os.Setenv("BhDatabaseParameters", "sslmode=disable connect_timeout=20")
	os.Setenv("BhDatabaseParameters", "connect_timeout=20&sslmode=disable")
	db = NewDBConn()
	t.Logf("%v \n", db)
}
