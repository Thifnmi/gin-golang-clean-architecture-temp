```repo struct```

- cmd/
    - main.go
- config/
    - config.go
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

```run server with go```

`go run cmd/main.go`


```run with docker```

updating ...