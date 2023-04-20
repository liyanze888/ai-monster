package pkg

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lqs/sqlingo"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

const (
	mysql = "mysql"
)

var (
	database         sqlingo.Database
	initDatabaseOnce sync.Once
)

// GetDatabase TODO
func GetDatabase() sqlingo.Database {
	initDatabaseOnce.Do(initMysql)
	return database
}

// initMysql init mysql database connection
// database.host: database host address
// database.port: database host port
// database.user: database access username
// database.pass: database access password
// database.name: database schema used
func initMysql() {
	dbHost := GetString(`database.host`)
	dbPort := GetString(`database.port`)
	dbUser := GetString(`database.user`)
	dbPass := GetString(`database.pass`)
	dbName := GetString(`database.name`)
	maxIdleConn := 8
	if value := viper.GetInt(`database.max_idle_conn`); value != 0 {
		maxIdleConn = value
	}
	maxOpenConn := 8
	if value := viper.GetInt(`database.max_open_conn`); value != 0 {
		maxOpenConn = value
	}

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	//log.Logger.Info("database connection url.", zap.String("connection", connection))
	var err error
	database, err = sqlingo.Open(mysql, connection)

	if err != nil {
		//log.Logger.Fatal("database connection error.", zap.Error(err))
	}
	database.EnableCallerInfo(true)
	database.SetInterceptor(func(ctx context.Context, sql string, invoker sqlingo.InvokerFunc) error {
		//非开发环境打开sql日志
		if !IsDev() {
			start := time.Now()
			defer func() {
				log.Printf("%v %s", time.Since(start), sql)
			}()
		}
		return invoker(ctx, sql)
	})
	database.GetDB().SetMaxIdleConns(maxIdleConn)
	database.GetDB().SetMaxOpenConns(maxOpenConn)
	//log.Logger.Info("Init mysql database connection")
	return
}
