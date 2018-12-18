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

// Bootstrapper for transform_wasm. Transforms all html in the test files
// specified on the commandline.
//
// To use:
//   GOOS=js GOARCH=wasm go build -o transform.wasm ./cmd/transform_wasm/ &&
//   node --max-old-space-size=7000 cmd/transform_wasm/main.js transform.wasm \
//     path/to/test/files*

// TODO(twifkak): Investigate slowdown over time.
// TODO(twifkak): Investigate memory usage growth.

const assert = require('assert');
const fs = require('fs');
const { join } = require('path');
const { spawnSync } = require('child_process');

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

global.begin = async function(transform, done) {
  let num = 0;
  let outs = (await readTestFiles()).map(([url, html]) =>
    new Promise((resolve) =>
      transform(url, html, (amphtml) => {
          if (++num % 100 == 0) console.log('num = ', num);
          // Minimum valid AMP is larger than 1K.
          if (amphtml.length < 1000) console.log('URL ', url, ' output is invalid: ', amphtml);
          resolve(amphtml);
      })));
  console.log('Pushed all %d thunks.', outs.length);
  const start = process.hrtime.bigint();
  await Promise.all(outs);
  const total = process.hrtime.bigint() - start;
  console.log(`Took ${total} nanoseconds, or ${Number(total) / outs.length / 1000000} millis per doc.`);
  done();
}

const goroot = process.env.GOROOT || spawnSync('go', ['env', 'GOROOT']).stdout.toString().trim();
require(join(goroot, 'misc/wasm/wasm_exec.js'));
