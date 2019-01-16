# WebAssembly local transformer

This is a proof of concept for running the [local
transformer](../../transformer) in a JS environment, using WebAssembly. It
depends on Go 1.12 (currently available [in
beta](https://golang.org/dl/#unstable)).

## Running

To run the test:
```bash
git clone https://github.com/ampproject/amppackager.git
cd amppackager/cmd/transform_wasm
GOROOT=~/sdk/go1.12beta2 node main.js main.wasm testfile
```

For a bigger testfile, pass `node` something like `--max-old-space-size=4000`.

lib.js can be reused for other purposes. See main.js for an example use.

## Building

Go's current wasm backend allocates a lot of memory and never frees it. To
conserve memory usage, build a patched Go compiler:

```bash
git clone https://github.com/twifkak/go.git
cd go
git checkout small
src/make.bash
```

And then use that to build the wasm binary:

```bash
cd amppackager/cmd/transform_wasm
GOOS=js GOARCH=wasm path/to/go/bin/go build -o main.wasm .
```
