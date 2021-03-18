# Go ToDO API

## Start Application
- Clone this repository.
- Run the following commands:
  - docker build -t todo-api .
  - docker-compose up -f

## API Routes


Path | Method  | Description
---|---|---
/api/todo | GET |   Get All Todo
/api/todo | POST | Save Todo
/api/todo | PUT | Update Todo
/api/todo/{todoId} | DELETE | Delete Todo

### Environment Variables
Name | Type | Default | Description / Mandatory
---|---|---|---
TODO_API_HOST | string |  | http address | no
TODO_API_PORT | string | 8080 | http port | no
TODO_API_API_BASE_URI | string | api | api base path | no
TODO_API_MYSQL_CONNECTION_STRING | string |  | mysql 'tcp' or 'unix' connection | yes
TODO_API_MYSQL_MAX_IDLE_CONN | string | 5 | database mysql max idle connection | no
TODO_API_MYSQL_MAX_OPEN_CONN | string | 5 |  database mysql max open connection | no
