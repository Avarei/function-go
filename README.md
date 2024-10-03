# function-go-plugin
[![CI](https://github.com/crossplane/function-go/actions/workflows/ci.yml/badge.svg)](https://github.com/crossplane/function-go/actions/workflows/ci.yml)

A **UNFINISHED CONCEPT** for a [composition function][functions] in [Go][go], which dynamically loads plugins from OCI Artifacts.
The Plugins just need to Implement the [FunctionRunnerServiceServer](https://pkg.go.dev/github.com/crossplane/function-sdk-go@v0.3.0/proto/v1#FunctionRunnerServiceServer) interface. (See example in [plugin/plugin.go](plugin/plugin.go)).

Advantages:
* No Boilerplate code required as it is abstracted in the Main function. Only focus is on the implementation.
* Instead of deploying a function AND a composition, which references the function only the composition needs to reference the oci artifact.

Disadvantages:
* Strict Dependency bindings between main App and 
* Relatively low reduction of complexity compared to function-template-go
* Requires additional Tooling to push the oci artifact

## Development

```shell
# Run code generation - see input/generate.go
$ go generate ./...

# Run tests - see fn_test.go
$ go test ./...

# Build the function's runtime image - see Dockerfile
$ docker build . --tag=runtime

# Build a function package - see package/crossplane.yaml
$ crossplane xpkg build -f package --embed-runtime-image=runtime
```

[functions]: https://docs.crossplane.io/latest/concepts/composition-functions
[go]: https://go.dev
[function guide]: https://docs.crossplane.io/knowledge-base/guides/write-a-composition-function-in-go
[package docs]: https://pkg.go.dev/github.com/crossplane/function-sdk-go
[docker]: https://www.docker.com
[cli]: https://docs.crossplane.io/latest/cli