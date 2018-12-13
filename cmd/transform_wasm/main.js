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

// Bootstrapper for transform_wasm. Transforms all html files in
// ${TESTDIR:-/tmp/amps}.
//
// To use:
//   GOOS=js GOARCH=wasm go build -o transform.wasm ./cmd/transform_wasm/ &&
//   node --max-old-space-size=4096 cmd/transform_wasm/main.js transform.wasm

// TODO(twifkak): Investigate slowdown over time.
// TODO(twifkak): Investigate memory usage growth.

const assert = require('assert');
const { join } = require('path');
const { spawnSync } = require('child_process');

function getRecursive(dir = process.env.TESTDIR || '/tmp/amps') {
  return [].concat(...fs.readdirSync(dir, {withFileTypes: true}).map((dirent) =>
      dirent.isDirectory() ? getRecursive(join(dir, dirent.name)) : join(dir, dirent.name)));
}

global.begin = async function(transform, done) {
  const htmlPaths = getRecursive().filter((file) => file.endsWith('.html'));
  let num = 0;
  let outs = [];
  for (path of htmlPaths) {
    const html = fs.readFileSync(path);
    outs.push(new Promise((resolve) => {
      transform('https://example.com/', html, (amphtml) => {
          if (++num % 100 == 0) console.log('num = ', num);
          assert(amphtml.length > 1000);  // "Minimum valid AMP" is larger than 1K.
          resolve(amphtml);
      });
    }));
  }
  console.log('Pushed all %d thunks.', htmlPaths.length);
  const start = process.hrtime.bigint();
  await Promise.all(outs);
  const total = process.hrtime.bigint() - start;
  console.log(`Took ${total} nanoseconds, or ${Number(total) / htmlPaths.length / 1000000} millis per doc.`);
  done();
}

const goroot = process.env.GOROOT || spawnSync('go', ['env', 'GOROOT']).stdout.toString().trim();
require(join(goroot, 'misc/wasm/wasm_exec.js'));
