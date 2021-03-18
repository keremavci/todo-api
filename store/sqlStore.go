package store

import (
	"context"
	"github.com/keremavci/todo-api/config"
	. "github.com/keremavci/todo-api/log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
	dbsql "database/sql"
	"github.com/go-gorp/gorp"

)

type SqlStore struct {
	dbMap *gorp.DbMap
	stores         SqlStoreStores

}

type SqlStoreStores struct {
	todo *TodoStore
}


func NewSqlStore(config *config.Config) *SqlStore {
	sqlStore := &SqlStore{}
	sqlStore.initConnection(config)
	sqlStore.stores.todo = newTodoStore(sqlStore);

	err := sqlStore.GetConnection().CreateTablesIfNotExists()
	if err != nil {
		if IsDuplicate(err) {
			Logger.Info("Test")
			Logger.Infof("Duplicate key error occurred; assuming table already created and proceeding.Error: %s", err.Error())
		} else {
			Logger.Fatalf("Error creating database tables. %s", err.Error())
			os.Exit(EXIT_CREATE_TABLE)
		}
	}
	return sqlStore
}


func (ss *SqlStore) TodoStore() *TodoStore {
	return ss.stores.todo
}

func (ss *SqlStore) GetConnection() (*gorp.DbMap) {
	if ss.dbMap == nil {
		Logger.Fatalf("Conenction doesn't exists.")
	}
	return ss.dbMap
}

func (ss *SqlStore) initConnection(config *config.Config) {
	ss.dbMap = setupConnection(config)//"username:password@tcp(127.0.0.1:3306)/test"
}


func setupConnection(config *config.Config) *gorp.DbMap {
	db, err := dbsql.Open("mysql", config.DBConfig.MySqlConnectionString)
	if err != nil {
		Logger.Errorf("Failed to open SQL connection to err.%s",err.Error())
		time.Sleep(time.Second)
		os.Exit(EXIT_DB_OPEN)
	}

	for i := 0; i < DB_PING_ATTEMPTS; i++ {
		Logger.Info("Pinging SQL")
		ctx, cancel := context.WithTimeout(context.Background(), DB_PING_TIMEOUT_SECS*time.Second)
		defer cancel()
		err = db.PingContext(ctx)
		if err == nil {
			break
		} else {
			if i == DB_PING_ATTEMPTS-1 {
				Logger.Fatalf("Failed to ping DB, server will exit.%s", err.Error())
				time.Sleep(time.Second)
				os.Exit(EXIT_PING)
			} else {
				Logger.Info("Failed to ping DB")
				time.Sleep(DB_PING_TIMEOUT_SECS * time.Second)
			}
		}
	}

	db.SetMaxIdleConns(config.DBConfig.MySqlMaxIdleConn)
	db.SetMaxOpenConns(config.DBConfig.MySqlMaxOpenConn)

	return  &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}

}