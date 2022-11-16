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
# Get All
curl -s 'http://localhost:8080/todos' | jq

# Create Todo
curl 'http://localhost:8080/todos' -d '{"title":"abc1","order":1,"completed":false}'
curl 'http://localhost:8080/todos' -d '{"title":"abc2","order":2,"completed":false}'

# Update Todo
curl -X 'PATCH' 'http://localhost:8080/todos/1' -d '{"title":"def","order":1,"completed":true}'

# Delete Todo
curl -X 'DELETE' 'http://localhost:8080/todos/1'

# Delete all Todos
curl -X 'DELETE' 'http://localhost:8080/todos'
```

## References

- [Todo-Backend](http://www.todobackend.com/)
