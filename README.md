# todo-backend-go-mux

- Go 1.9

## Packages

- gorilla/mux <https://github.com/gorilla/mux>

## Commands

```bash
make serve
```

## Test cases

```bash
# Create Todo
curl 'http://localhost:8080/todos' -d '{"title":"abc","order":1,"completed":false}'

# Get All
curl -s 'http://localhost:8080/todos' | jq
```

## References

- [Todo-Backend](http://www.todobackend.com/)
