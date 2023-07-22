# Cake CRUD

## Dependencies
- Git
- Go
- Docker
- Mysql

## Getting Started
1. Configure db connection in `.env` file 
2. Run the docker file, create and start the container 

## Using Migration
The migration here is using https://github.com/golang-migrate/migrate, run the below command to migrate the database
`migrate -path Migration/ -database "{your_mysql_string_connection}" -verbose up`

## Documentation
- https://github.com/golang-migrate/migrate
- https://go.dev/doc/
- https://github.com/go-sql-driver/mysql
- https://github.com/google/wire
- https://docs.docker.com/
