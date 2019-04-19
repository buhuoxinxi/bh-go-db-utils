package configs

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

// database env
const (
	// connection
	envKeyDatabaseDriver     = "BhDatabaseDriver"     // driver
	envKeyDatabaseUsername   = "BhDatabaseUsername"   // username
	envKeyDatabasePassword   = "BhDatabasePassword"   // password
	envKeyDatabaseHost       = "BhDatabaseHost"       // host
	envKeyDatabasePort       = "BhDatabasePort"       // port
	envKeyDatabaseDBName     = "BhDatabaseDBName"     // db name
	envKeyDatabaseParameters = "BhDatabaseParameters" // parameters

	// options
	envKeyDatabaseOrmDebug        = "BhDatabaseOrmDebug"        // debug
	envKeyDatabaseTablePrefix     = "BhDatabaseTablePrefix"     // prefix
	envKeyDatabaseMaxOpenConn     = "BhDatabaseMaxOpenConn"     // max open
	envKeyDatabaseMaxIdleConn     = "BhDatabaseMaxIdleConn"     // idle
	envKeyDatabaseConnMaxLifetime = "BhDatabaseConnMaxLifetime" // lifetime
)

// AuthConfig auth cfg
type AuthConfig struct {
	Driver     string // os.Setenv("DatabaseDriver", "mysql")
	Username   string // os.Setenv("DatabaseUsername", "username")
	Password   string // os.Setenv("DatabasePassword", "password")
	Host       string // os.Setenv("DatabaseHost", "127.0.0.1")
	Port       string // os.Setenv("DatabasePort", "3306")
	DBName     string // os.Setenv("DatabaseDBName", "test")
	Parameters string // os.Setenv("DatabaseParameters", "charset=utf8&timeout=60s&loc=Local&autocommit=true")
}

// GetAuthConfigHandler get database config
var GetAuthConfigHandler = func() *AuthConfig {
	var cfg AuthConfig

	cfg.Driver = strings.TrimSpace(os.Getenv(envKeyDatabaseDriver))
	cfg.Username = strings.TrimSpace(os.Getenv(envKeyDatabaseUsername))
	cfg.Password = strings.TrimSpace(os.Getenv(envKeyDatabasePassword))
	cfg.Host = strings.TrimSpace(os.Getenv(envKeyDatabaseHost))
	cfg.Port = strings.TrimSpace(os.Getenv(envKeyDatabasePort))
	cfg.DBName = strings.TrimSpace(os.Getenv(envKeyDatabaseDBName))
	cfg.Parameters = strings.TrimSpace(os.Getenv(envKeyDatabaseParameters))

	return &cfg
}

// tablePrefix : database table prefix
var tablePrefix = strings.TrimSpace(os.Getenv(envKeyDatabaseTablePrefix))

// GetTablePrefixHandler get table prefix
var GetTablePrefixHandler = func() string {
	return tablePrefix
}

// OptionConfig option cfg
type OptionConfig struct {
	// sets orm LogMode
	Debug bool // os.Setenv("DatabaseOrmDebug", "true")

	// sets table prefix
	Prefix string // os.Setenv("DatabaseTablePrefix", "ikg_")

	// sets the maximum number of open connections to the database.
	MaxOpen int // os.Setenv("DatabaseMaxOpenConn", "10")

	// sets the maximum number of connections in the idle
	MaxIdle int // os.Setenv("DatabaseMaxIdleConn", "10")

	// sets the maximum amount of time a connection may be reused.
	MaxLifetime time.Duration // os.Setenv("DatabaseConnMaxLifetime", "30s")
}

// GetOptionConfigHandler get database config
var GetOptionConfigHandler = func() *OptionConfig {
	var cfg OptionConfig
	var err error

	//cfg.Debug = strings.ToLower(debugString) == "true"
	cfg.Debug, _ = strconv.ParseBool(strings.TrimSpace(os.Getenv(envKeyDatabaseOrmDebug)))

	// prefix
	cfg.Prefix = GetTablePrefixHandler()

	// max open
	cfg.MaxOpen, err = strconv.Atoi(strings.TrimSpace(os.Getenv(envKeyDatabaseMaxOpenConn)))
	if err != nil {
		logrus.Println("strconv.Atoi(envKeyDatabaseMaxOpenConn) error : ", err)
	}

	// max idle
	cfg.MaxIdle, err = strconv.Atoi(strings.TrimSpace(os.Getenv(envKeyDatabaseMaxIdleConn)))
	if err != nil {
		logrus.Println("strconv.Atoi(envKeyDatabaseMaxIdleConn) error : ", err)
	}

	// max lifetime
	cfg.MaxLifetime, err = time.ParseDuration(strings.TrimSpace(os.Getenv(envKeyDatabaseConnMaxLifetime)))
	if err != nil {
		logrus.Println("time.ParseDuration(envKeyDatabaseConnMaxLifetime) error : ", err)
	}
	return &cfg
}
