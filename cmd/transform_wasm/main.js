// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// JS runtime for transform_wasm. Transforms all html in the test files
// specified on the commandline. This employs several tricks to get reasonable
// performance out of Go-WASM:
//
// 1. To eliminate Go bootstrapping costs: Keep a long-running Go process open
//    and communicate with it via continuation-passing style.
// 2. To eliminate memory leaks of data passed between Go/JS due to lack of
//    cross-heap GC: Communicate large data by mutating persistent TypedArrays
//    created from the Go side.
// 3. To eliminate the 1G hard-coded minimum Go heap size: use the third-party
//    wams tool to hack the wasm binary. This requires more cruft because
//    `memory.grow` requests from the Go runtime cause TypedArrays to
//    invalidate as the memory underneath them moves.
//
// This runtime leaks ~603 bytes per call to transform, associated with the
// flyweight arrow functions that are passed to getter() and transform(). I
// could engineer away these leaks by adding more global state, but ultimately
// they will go away with Go 1.12, which replaces the asynchronous js.Callback
// with a synchronous js.Func: https://tip.golang.org/pkg/syscall/js/#Func.
//
// Build dependency: go get github.com/termonio/wams
//
// To use:
//   GOOS=js GOARCH=wasm go build -o transform.wasm ./cmd/transform_wasm/ &&
//   wams -pages 2048 -write transform.wasm &&
//   node --max-old-space-size=4000 cmd/transform_wasm/main.js transform.wasm \
//     cmd/transform_wasm/testfile

const lib = require('./lib');

const fs = require('fs');
const util = require('util');

const { join } = require('path');

// Polyfill to flatten an array by one level.
const flat = [].flat ? (arr) => arr.flat() : (arr) => [].concat(...arr);

function listRecursive(dir) {
  return flat(fs.readdirSync(dir, {withFileTypes: true}).map((dirent) =>
      dirent.isDirectory() ? listRecursive(join(dir, dirent.name)) : join(dir, dirent.name)));
}

// Take everything after "transform.wasm" and remove it from argv so that
// wasm_exec.js doesn't pass it to the Go binary.
const testFiles = process.argv.splice(3);

const markerText = '>>>>>>>>>> Test Case <<<<<<<<<<\n';

async function readTestFiles() {
  // Read all the HTML test cases into memory.
  let htmls = [];
  for (const testFile of testFiles) {
    console.log(`Opening ${testFile}...`);
    let pending = '';
    for await (const chunk of fs.createReadStream(testFile, {encoding: 'utf8'})) {
      pending += chunk;
      let pastLastMarker = 0; // Position just past the previously found marker.
      let marker; // Position of the current marker.
      while (marker = pending.indexOf(markerText, pastLastMarker), marker !== -1) {
        if (marker > pastLastMarker)
          htmls.push(pending.substring(pastLastMarker, marker));
        pastLastMarker = marker + markerText.length;
      }
      pending = pending.substring(pastLastMarker);
    }
    htmls.push(pending);
  }

  // Parse the URL from each test case.
  htmls.forEach((html, i) => {
    let newline = html.indexOf('\n');
    htmls[i] = [html.substring(0, newline), html.substring(newline + 1)];
  });

  console.log('Pushed all %d tests.', htmls.length);

  return htmls;
}

const heapdump = (() => { try { return require('heapdump') } catch { } })(); // npm install heapdump
function dumpHeap(name, full) {
  console.log('%s: %s', name, util.inspect(process.memoryUsage(), {colors: true, breakLength: Infinity}))
  if (full && heapdump) heapdump.writeSnapshot('wasm.js.' + name + '.heapsnapshot');
}

(async function() {
  global.tests = await readTestFiles();

  dumpHeap('compile.before');
  lib.start(async function() {
    dumpHeap('compile.after', true);

    const start = process.hrtime.bigint();
    let num = 0;
    for (const [url, html] of tests) {
      if (++num % 100 === 0) console.log('num =', num);
      if (num % 2000 === 0) dumpHeap('transform.' + num);

      try {
        let out = await lib.transform(url, html);
      } catch(err) {
        console.error("error for", url, err);
        continue;
      }
    }
    const total = process.hrtime.bigint() - start;

    dumpHeap('transform.after', true);
    console.log(`Took ${total} nanoseconds, or ${Number(total) / tests.length / 1000000} millis per doc.`);
  }, true);
})();
