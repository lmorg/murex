# Murex Documentation

## Running Tests

### Go tests

```
go test ./...
```

```
go run test/count/server &
export MUREX_TEST_COUNT=http
go test ./...
curl localhost:38000/total
```
