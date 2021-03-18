package config

import (
	"os"
	"testing"
)

func TestValidateConfig(t *testing.T) {
	config := DefaultConfig()

    if(config.DBConfig.MySqlConnectionString != os.Getenv("TODO_API_MYSQL_CONNECTION_STRING")){
		t.Fatalf("Didn't give valid mysql connection string")
	}


}