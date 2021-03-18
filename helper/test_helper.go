package helper

import (
	"context"
	"fmt"
	. "github.com/keremavci/todo-api/log"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
)




func CreateMySqlContainerForTest() (mysqlConnStr string, err error) {
	ctx := context.Background()

	var env = map[string]string{
		"MYSQL_ROOT_PASSWORD": "todorootpwd",
		"MYSQL_DATABASE":     "todo",
		"MYSQL_USER":       "todouser",
		"MYSQL_PASSWORD": "todopassword",
	}

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mysql:5.7",
			ExposedPorts: []string{"3306/tcp"},
			Env:          env,
			//WaitingFor:   wait.ForListeningPort(nat.Port("3306/tcp")),
		},
		Started: true,
	}
	mysqlContainer, err := testcontainers.GenericContainer(context.Background(), req)
	if err != nil {
		return "", errors.Wrap(err,"Failed to start mysql container ")
	}

	natPort, err := mysqlContainer.MappedPort(ctx, "3306")
	if err != nil {
		Logger.Error(err)
	}
	mysqlConnStr = fmt.Sprintf("%s:%s@(localhost:%s)/%s", env["MYSQL_USER"], env["MYSQL_PASSWORD"], natPort.Port(), env["MYSQL_DATABASE"])


	return mysqlConnStr,nil
}