# Sample monorepo

This repository attempts to be a testbed for exploring Bazel. The idea is to build a mini version of
everything we need to do for Pachyderm production code here first, and see if it's okay before we
commit to tackling that complexity.

Today we have:

-   Go (server, tests, force-disabling cgo like pachyderm/pachyderm does, nogo lint)

To add:

-   Go linting (checks from golangci-lint)
-   Helm (releases)
-   Jsonnet (lint, tests)
-   Python (PyPI releases, lint)
-   Typescript + React (npm releases, lint, tests)
-   Docker builds (multiarch)
-   Kubernetes helpers (build containers, setup minikube, deploy helm chart, upgrade helm chart
    after rebuild, etc.)
-   Darwin/Arm64 cc toolcahin
-   Properly vendor tooling (prettier, kind/minikube, etc.)
-   Buildifier lint / test; verifying that gazelle was run.

# How to

## Install Bazel

https://bazel.build/install

You can setup tab completion in your shell after this, see https://bazel.build/install/completion

## Setup your editor for Bazel

https://bazel.build/install/ide

## Setup your editor for Go

https://github.com/bazelbuild/rules_go/wiki/Editor-setup

Basically, configure LSP integration to set `GOPACKAGESDRIVER` to `tools/gopackagesdriver.sh` when
such a thing exists in the project root (as it does in this project). This allows you to treat
generated files (like protobufs) as though they were actually checked in, even though they're not.

## Run the tests

```shell
bazel test ...
```

Output goes to a file by default. That's useful when testing everything, but not so helpful when
you're debugging a single test. Run this instead:

```shell
bazel test --test_output=all ...
```

You can add a `test --test_output=errors` or `test --test_output=all` line to a `.bazelrc.local`
file in this directory if you find yourself always running the tests like this.

## Run the Go server

```shell
bazel run //go/cmd/go-echo
```

## Generate BUILD.bazel files for new Go / proto / Python code

```shell
bazel run gazelle
```

If you're creating a new package or file, you probably want to just make a minimal go file:

```go
package whatever
```

and then run Gazelle. You won't have editor hints until there is a bazel rule that builds the file
you're editing.

## Use a go module

```shell
go get github.com/somebody/some-library@latest
bazel run gazelle-update-repos
```

Do note that things like `go mod tidy` don't work anymore; `go mod` doesn't know about
`GOPACKAGESDRIVER` and thus thinks generated packages are third-party libraries that we depend on,
but they're not.

It seems likely that go.mod and the actual dependencies we depend on will get out of sync over time.
At the end of the day, `deps.bzl` is authoritative.

## go mod why

You'll need to use [bazel queries](https://docs.bazel.build/versions/main/query-how-to.html) instead
of Go tooling here.

## Cross-compiling

We support Bazel platforms, so you can build a binary for a foreign architecture like:

```shell
bazel build //go/cmd/go-echo  --platforms=//:linux_arm64
```

Platforms are declared in `platforms.bzl`. List them with:

```shell
bazel query 'kind("platform", ...)'
```

Adding the `--toolchain_resolution_debug=.` to the build command will print some messages about
which toolchain is being used. Cross-compilation involves installing a mini Linux distribution in
the output directory that contains tools that will exist on the target system. Right now we use
ChromeOS's sysroots; a list of which is available here:
https://chromium.googlesource.com/chromium/src/build/+/refs/heads/main/linux/sysroot_scripts/sysroots.json

In the future we will probably use a Docker rule to pull Debian stable or something.

It would also be nice to somehow run cross-compiled code in an emulator for a basic sanity check.
Probably quite possible. I don't know how to do it.

## Deploy the current state of the workspace to a live k8s environment

Create a KinD cluster:

```shell
reg_name='kind-registry'
reg_port='5001'
if [ "$(docker inspect -f '{{.State.Running}}' "${reg_name}" 2>/dev/null || true)" != 'true' ]; then
  docker run \
    -d --restart=always -p "127.0.0.1:${reg_port}:5000" --name "${reg_name}" \
    registry:2
fi

cat <<EOF | kind create cluster --name=kind --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:5000"]
EOF
```

Then deploy: `bazel run local.create`.

After you make a change: `bazel run local.replace`. You can see what `replace` is going to do with
`bazel run local.diff`.

## Run Python

```shell
bazel run //py/hello
```

Note that you need Python 3 installed, but only to do some housekeeping; the actual script runs with
the Python version declared in WORKSPACE: https://github.com/bazelbuild/rules_python/issues/691
