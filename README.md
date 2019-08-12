# go-playground

## Work with Go installed on current machine

Follow [Getting Started](https://golang.org/doc/install) to install Go

Run:

```
go run src/main.go
```

## Work with Docker

### Build Image

```
docker build -t go-playground .
```

### Run go-playground

```
docker run --rm -it -v /$(pwd)/src:/app/src go-playground
```
