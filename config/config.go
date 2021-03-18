package config

import "os"
import "strconv"

type Config struct {
	DBConfig DBConfig
	ServerConfig ServerConfig
	APIBaseUri string
}

type ServerConfig struct {
	Port string
	Host string
}

type DBConfig struct {
	MySqlConnectionString string
	MySqlMaxIdleConn int
	MySqlMaxOpenConn int

}
func DefaultConfig() *Config {
	mysqlMaxIdleConn, _ := strconv.Atoi(getEnv("TODO_API_MYSQL_MAX_IDLE_CONN","5"))
	mysqlMaxOpenConn, _ := strconv.Atoi(getEnv("TODO_API_MYSQL_MAX_OPEN_CONN","5"))

	return &Config{
		ServerConfig: ServerConfig{
			Port: getEnv("TODO_API_PORT","8080"),
			Host: getEnv("TODO_API_HOST", ""),
		},
		DBConfig: DBConfig{
			MySqlConnectionString: getEnv("TODO_API_MYSQL_CONNECTION_STRING",""),
			MySqlMaxIdleConn: mysqlMaxIdleConn,
			MySqlMaxOpenConn: mysqlMaxOpenConn,
		},
		APIBaseUri: getEnv("TODO_API_API_BASE_URI","/api"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}