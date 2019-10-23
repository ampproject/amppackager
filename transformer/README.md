# Local Transformer

The modifications in this package are described in more detail
[here](https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-modifications.md).

The transformed AMP HTML produced by the library is meant to be used inside of
a signed exchange, but may be valid in other contexts, as well.

## How to use
The local transformer can be used separately from the packager/signer. Here's an
example use of the binary:

1. `go get -u github.com/ampproject/amppackager/cmd/transform`
1. `$GOPATH/bin/transform -url "documentURL" /path/to/html`

For more help, `$GOPATH/bin/transform -h`

See the [binary source code](../cmd/transform/transform.go) for an example use
of the library.
