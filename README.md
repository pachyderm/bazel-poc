# Sample monorepo

This repository attempts to be a testbed for exploring Bazel.

# How to

## Run the tests

```shell
bazel test //...
```

## Run the Go server

```shell
bazel run //go/cmd/go-echo:go-echo
```

## Generate BUILD.bazel files for new Go code

```shell
bazel run //:gazelle
```

## Use a go module

```shell
go get github.com/somebody/some-library@latest
bazel run //:gazelle-update-repos
go mod tidy
```
