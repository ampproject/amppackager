// Library wrapping the Go transform function in a simpler API.

const util = require('util');

/**
 * @param url {string}
 * @param html {string}
 * @return {string} The transformed HTML.
 */
async function transform(url, html) {
  if (!(urlIn instanceof GoBytes)) {
    urlIn = new GoBytes(urlIn);
    htmlIn = new GoBytes(htmlIn);
    htmlOut = new GoBytes(htmlOut);
  }

  await urlIn.set(url);
  await htmlIn.set(html);

  return await new Promise((resolve) =>
    transformCB(() => {
      // Minimum valid AMP is larger than 1K.
      htmlOut.get().then((str) => {
        if (str.length < 1000) console.log('URL', url, 'output is invalid: ', str);
        resolve(str);
      });
    }));
}
exports.transform = transform;

// Wraps a TypedArray as received by Go, taking care of:
// - Length-prefix and UTF-8 decoding/encoding in the get()/set() methods.
// - Checking that the given string will fit in the buffer, in set().
// - Getting a new TypedArray from Go, if the old one detaches due to WASM
//   memory growth.
class GoBytes {
  constructor(wrapper) {
    this._wrapper = wrapper;
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
    if (len > this._wrapper.maxLen) throw new Error("str is corrupted; unexpected len " + len);
    return this._decoder.decode(new DataView(ta.slice(4, 4 + len).buffer));
  }

  async set(str /*string*/) {
    let buf = this._encoder.encode(str);
    if (buf.length > this._wrapper.maxLen) throw new Error("str too big: " + buf.length);
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
      if (this._releaser) await new Promise((resolve) => this._releaser(() => resolve()));
      await new Promise((resolve) =>
          this._wrapper.getter((ta, rel) => {
            this._typedArray = ta;
            this._releaser = rel;
            resolve();
          }));
    }
    return this._typedArray;
  }
}

