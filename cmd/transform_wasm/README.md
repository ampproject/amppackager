This is a proof of concept for running the [local
transformer](../../transformer) in a JS environment, using WebAssembly.

To build the Go WASM binary:
```bash
GOOS=js GOARCH=wasm go get github.com/ampproject/amppackager/cmd/transform_wasm
```

Optionally, you can use [wams](https://github.com/termonio/wams) to shrink the
initial Go heap size, though it's not clear this has an effect on the process's
RSS:

```bash
go get github.com/termonio/wams
wams -pages 2048 -write transform.wasm
```

To run the test:
```bash
git clone https://github.com/ampproject/amppackager.git
cd amppackager
node --max-old-space-size=4000 cmd/transform_wasm/main.js \
  ${GOROOT:-~/.go}/bin/js_wasm/transform_wasm cmd/transform_wasm/testfile
```

lib.js can be reused for other purposes. See main.js for an example use.
