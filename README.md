## Repo struct

- cmd/
    - main.go
- config/
    - config.go
- migrations/
    - sql/
        - mysql/
            - sql_file
- pkg/
    - domain/
        - user.go
    - repositories/
        - user_repository.go
    - usecases/
        - user_usecase.go
    - interfaces/
        - http/
            - handlers/
                - user_handler.go
            - router/
                - user/
                    - user.go
- vendor/
- go.mod
- go.sum

## Run server with golang

```bash
go run cmd/main.go
```


## Run with docker

updating ...


## Migrate database

I using [go-migrate]() tool to migrate database

* Create sql file

```bash
migrate create -ext sql -dir mirations/sql/mysql -seq <sql_file_name>
```
* After write sql to `sql_file_name.up.sql` flow [sql syntax](https://en.wikipedia.org/wiki/SQL_syntax)

* Apply migrate

```bash
migrate -path mirations/sql/mysql -database "mysql://username:password@tcp(ip:port)/database_name" up
```
