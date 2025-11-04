# hello-go

Minimal Go project that prints "Hello, World!" with a unit test.

## Run

```bash
go run .

# or build to bin/
mkdir -p bin && go build -o bin/hello-go .
./bin/hello-go
```

## Test

```bash
go test ./...
```

## Structure

- main.go: entry point and Hello() implementation
- main_test.go: minimal unit test
- .gitignore: ignore bin/ and build artifacts
- go.mod: module definition (hello-go)


