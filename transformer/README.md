# Local Transformer

The modifications in this package are described in more detail
[here](https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-modifications.md).

> NOTE: The transformed AMP HTML produced by the library is only valid inside of
> a signed exchange, and not to be served as normal HTML. Also, the library is
> still a work-in-progress and not all transformations described in the link
> above are implemented.

## How to use
The local transformer can be used separately from the packager/signer. Here's an
example use of the binary:

1. `go get -u github.com/ampproject/amppackager/cmd/transform`
1. `$GOPATH/bin/transform -url "documentURL" /path/to/html`

For more help, `$GOPATH/bin/transform -h`

See the [binary source code](../cmd/transform/transform.go) for an example use
of the library.
