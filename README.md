# AMP Packager

> **WARNING**: This code is still brand new, and highly experimental. The
> specification is still changing, and this is an implementation of a snapshot
> of it, as a proof of concept. Feel free to play around with it, but be
> cautious. Examine the code.

AMP Packager is a tool to [improve AMP
URLs](https://amphtml.wordpress.com/2018/01/09/improving-urls-for-amp-pages/).
By running it in a proper configuration, web publishers may (eventually) have
origin URLs appear in AMP search results.

The AMP Packager works by creating [Signed HTTP
Exchanges (SXGs)](https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html)
containing AMP documents, signed with a certificate associated with the origin,
with a maximum lifetime of 7 days. In the future, the [Google AMP
Cache](https://www.ampproject.org/docs/fundamentals/how_cached) will fetch,
cache, and serve them, similar to what it does for normal AMP HTML documents.
When a user loads such an SXG, Chrome validates the signature and then displays
the certificate's domain in the URL bar instead of `google.com`, and treats the
web page as though it were on that domain.

The packager is an HTTP server that sits behind a frontend server; it fetches
and signs AMP documents as requested by the AMP Cache.

## Packager/Signer

### How to use

In all the instructions below, replace `example.com` with a domain you own and
can obtain certificates for.

#### Development server

  1. Install Go version 1.10 or higher. Optionally, set
     [$GOPATH](https://github.com/golang/go/wiki/GOPATH) to something (default
     is `~/go`) and/or add `$GOPATH/bin` to `$PATH`.
  2. `go get -u github.com/ampproject/amppackager/cmd/amppkg`

     Optionally, move the built `~/go/bin/amppkg` wherever you like.
  3. Create a file `amppkg.toml`. A minimal config looks like this:
     ```
     LocalOnly = true
     PackagerBase = 'https://localhost:8080/'
     CertFile = 'path/to/fullchain.pem'
     KeyFile = 'path/to/privkey.pem'
     OCSPCache = '/tmp/amppkg/ocsp'

     [[URLSet]]
       [URLSet.Sign]
         Domain = "example.com"
     ```
     More details can be found in [amppkg.example.toml](amppkg.example.toml).
  4. `mkdir /tmp/amppkg`
  5. `amppkg -development`

     If `amppkg.toml` is not in the current working directory, pass
     `-config=/path/to/amppkg.toml`.

#### Test your config

  1. Run Chrome M70 or later (as of 2018-09-18, this is
     [Beta](https://www.google.com/chrome/beta/) or
     [Dev](https://www.google.com/chrome/dev/)). On the
     command-line, pass the following flags:
     ```
     --user-data-dir=/tmp/udd
     --ignore-certificate-errors-spki-list=$(openssl x509 -pubkey -noout -in path/to/fullchain.pem | openssl pkey -pubin -outform der | openssl dgst -sha256 -binary | base64)
     --enable-features=SignedHTTPExchange
     'data:text/html,<a href="https://localhost:8080/priv/doc/https://example.com/">click me'
     ```
  2. Open DevTools. Check 'Preserve log'.
  3. Click the `click me` link.
  4. Watch the URL transmogrify! Verify it came from an SXG by switching
     DevTools to the Network tab and looking in the Size column for "(from
     signed-exchange)".

#### Demonstrate privacy-preserving prefetch

This step is optional; just to show how [privacy-preserving
prefetch](https://wicg.github.io/webpackage/draft-yasskin-webpackage-use-cases.html#private-prefetch)
works with SXGs.

  1. `go get -u github.com/ampproject/amppackager/cmd/amppkg_dl_sxg`.
  2. `amppkg_dl_sxg https://localhost:8080/priv/doc/https://example.com/`
  3. Stop `amppkg` with Ctrl-C.
  4. `go get -u github.com/ampproject/amppackager/cmd/amppkg_test_cache`.
  5. `amppkg_test_cache`
  6. Open Chrome and DevTools.
  7. Visit `https://localhost:8000/`. Observe the prefetch of `/test.sxg`.
  8. Click the link. Observe that the cached SXG is used.

#### Productionizing

For now, productionizing is a bit manual. The minimum steps are:

  1. Don't pass `-development` flag to `amppkg`. This causes it to serve HTTP
     rather than HTTPS, among other changes.
  2. Don't expose `amppkg` to the outside world; keep it on your internal
     network.
  3. Configure your TLS-serving frontend server to conditionally proxy to
     `amppkg`, if any of:
     1. The URL starts with `/amppkg/`.
     2. The URL is AMP and the `AMP-Cache-Transform` request header is present.
  4. Configure your frontend server to add `Vary: AMP-Cache-Transform` if the
     URL is AMP. This should occur for both HTML and SXG responses.
  5. Get an SXG cert from your CA. It must use an EC key with the prime256v1
     algorithm, and it must have a [CanSignHttpExchanges
     extension](https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html#cross-origin-cert-req).
     You MUST use this in `amppkg.toml`, and MUST NOT use it in your frontend.

You may also want to:

  1. Launch `amppkg` as a restricted user.
  2. Save its stdout to a rotated log somewhere.

Once you've done the above, you should be able to test by launching Chrome
without any comamndline flags; just make sure
chrome://flags/#enable-signed-http-exchange is enabled. To test by visiting the
packager URL directly, first add a Chrome extension to send an
`AMP-Cache-Transform: any` request header. Otherwise, follow the above
"Demonstrate privacy-preserving prefetch" instructions.

#### Redundancy

If you need to load balance across multiple instances of `amppkg`, you'll want
your `OCSPCache` to be backed by a shared storage device (e.g. NFS). It doesn't
need to be shared among all instances globally, but perhaps among all instances
per datacenter. The reason for this is to reduce the number of OCSP requests
`amppkg` needs to make, per [OCSP stapling
recommendations](https://gist.github.com/sleevi/5efe9ef98961ecfb4da8).

#### How will these web packages be discovered by Google?

For now, the presence of the `Vary: AMP-Cache-Transform` response header on an
AMP HTML page will allow the Google AMP Cache to make a second request with
`AMP-Cache-Transform: google` for the SXG.

In the future, Googlebot may make all requests with `AMP-Cache-Transform: google`,
eliminating the double fetch.

### Limitations

Currently, the packager will refuse to sign any AMP documents larger than 4 MB.
Patches that allow for streamed signing are welcome.

The packager refuses to sign any URL that results in a redirect. This is by
design, as neither the original URL nor the final URL makes sense as the signed
URL.

To account for possible clock skew in user agents, the packager back-dates
packages by 24h, which means they effectively last only 6 days for most users.

This tool only packages AMP documents. To sign non-AMP documents, look at the
commandline tools on which this was based, at
https://github.com/WICG/webpackage/tree/master/go/signedexchange.

## Local Transformer

The local transformer is a library within the AMP Packager that transforms AMP
HTML for security and performance improvements. Ports of or alternatives to the
AMP Packager will need to include these transforms.

More info [here](transformer/README.md).
