# Task API

Install dependencies
```
go mod tidy
```

Migrate database
```
goose postgres "postgres://postgres:password@localhost:5432/task" up
```

Run API
```
go run cmd/task-api/main.go
```

Try to insert data using Postman