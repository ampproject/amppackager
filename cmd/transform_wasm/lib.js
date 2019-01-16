// Library wrapping the Go transform function in a simpler API.

const util = require('util');

const { join } = require('path');
const { spawnSync } = require('child_process');

/**
 * The main entry point.
 *
 * @param callback {function()} The async function to call once setup is done.
 *   All calls to transform should happen while this function is running.
 * @param opt_test {boolean=} whether to output heap stats for testing
 */
async function start(callback, opt_test) {
  GoBridge.begin = async function(done) {
    await callback();
    done();
  };
  GoBridge.test = !!opt_test;
  // The patched Go runtime doesn't support concurrent GC yet.
  process.env.GODEBUG = 'gcstoptheworld=1';
  // Telling the GC to run more often keeps memory usage low.
  process.env.GOGC = '20';
  const goroot = process.env.GOROOT || spawnSync('go', ['env', 'GOROOT']).stdout.toString().trim();
  require(join(goroot, 'misc/wasm/wasm_exec.js'));
}
exports.start = start;

/**
 * Transforms a doc. Can only be called within the callback to start().
 *
 * @param url {string}
 * @param html {string}
 * @return {string} The transformed HTML.
 */
async function transform(url, html) {
  await urlIn.set(url);
  await htmlIn.set(html);

  transformCB();
  const str = await htmlOut.get();
  // Minimum valid AMP is larger than 1K.
  if (str.length < 1000) console.log('URL', url, 'output is invalid: ', str);
  return str;
}
exports.transform = transform;

// Internal communication between lib.js and the Go runtime.
global.GoBridge = {};

// Internal class for use by the Go runtime. Wraps a TypedArray, taking care of:
// - Length-prefix and UTF-8 decoding/encoding in the get()/set() methods.
// - Checking that the given string will fit in the buffer, in set().
// - Getting a new TypedArray from Go, if the old one detaches due to WASM
//   memory growth.
class Bytes {
  constructor(getter, maxLen) {
    this._getter = getter;
    this._maxLen = maxLen;
    this._typedArray = null;
    this._releaser = null;
    this._decoder = new util.TextDecoder();
    this._encoder = new util.TextEncoder();
  }

  async /*int*/ length() {
    let ta = await this._buf();
    return new DataView(ta.slice(0, 4).buffer).getUint32(0);
  }

  async /*string*/ get() {
    let ta = await this._buf();
    let buf = ta.slice(0, 4).buffer;
    let len = new DataView(buf).getUint32(0);
    if (len > this._maxLen) throw new Error("str is corrupted; unexpected len " + len);
    return this._decoder.decode(ta.slice(4, 4 + len).buffer);
  }

  async set(str /*string*/) {
    let buf = this._encoder.encode(str);
    if (buf.length > this._maxLen) throw new Error("str too big: " + buf.length);
    let ta = await this._buf();
    let tmpBuf = new Uint8Array(4);
    new DataView(tmpBuf.buffer).setUint32(0, buf.length);
    ta.set(tmpBuf);
    ta.set(buf, 4);
  }

  // Returns a Uint8Array backed by a Go slice. For some reason you can access
  // the Uint8Array through most methods, but not its buffer property. As a
  // workaround, use slice to make a shallow copy, so you can get a Uint8Array
  // with a working buffer property.
  async _buf() {
    if (!this._typedArray /* first use */ || !this._typedArray.length /* detached */) {
      if (this._releaser) this._releaser();
      const {ta, release} = this._getter();
      this._typedArray = ta;
      this._releaser = release;
    }
    return this._typedArray;
  }
}
GoBridge.Bytes = Bytes;
