package store

import (
	"github.com/keremavci/todo-api/config"
	"github.com/keremavci/todo-api/helper"
	"testing"
)




func TestNewSqlStore(t *testing.T){
	mysqlConnString, err:=helper.CreateMySqlContainerForTest()
	if err != nil {
		t.Fatalf("Cannot open mysql container.%v", err.Error())

	}
	config := &config.Config{}
	config.DBConfig.MySqlConnectionString=mysqlConnString
	mysqlConn := NewSqlStore(config)

	if err = mysqlConn.GetConnection().Db.Ping(); err != nil{
		t.Fatalf("Cannot ping mysql.%v", err.Error())
	}

}