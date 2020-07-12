build binary
```shell script
go build -o main main.go
./main
```

or build and run

```shell script
go run main.go
```

List all macros
```shell script
clang -dM -E -x c /dev/null
```