# AMP Packager

> **WARNING**: This code is still brand new, and highly experimental. The
> specification is still changing, and this is an implementation of a snapshot
> of it, as a proof of concept. Feel free to play around with it, but be
> cautious. Examine the code.

AMP Packager creates "AMP Packages" (implemented as [Signed HTTP
Exchanges](https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00))
containing AMP documents. These packages are consumed by the [Google AMP
Cache](https://www.ampproject.org/docs/fundamentals/how_cached) and cached.
Eventually, a future variant of these packages will be linked to from Google
Search instead of normal AMP page. The packages are signed with a certificate,
which means that Chrome users will see that certificate's domain in the URL bar
instead of `google.com` or `ampproject.org`, and that the web page will run on
that origin.

These packages have a maximum lifetime of 7 days, which minimizes the risk of
another cache serving a stale copy of this signed content. In the future, the
[AMP update-cache API](https://developers.google.com/amp/cache/update-cache)
will support updating AMP Packages, though it doesn't yet.

The packager is an HTTP server that sits behind a frontend server; it fetches
and signs AMP documents as requested by the AMP Cache.

## Packager/Signer

### How to use

#### Configure your frontend server

The frontend server needs to forward three types of requests to the packager:
packages, certificates, and validity maps.

##### Packages

The frontend needs to forward requests for web packages to the packager. The
details of this are still being worked out, but for now, an easy way to do so is
by URL. Come up with a simple URL mapping between an AMP package URL and its
corresponding AMP HTML URL. For instance, you might append `.htxg` to the URL,
so for an AMP URL:

```
https://example.com/url/to/amp.html
```

the corresponding AMP package URL would be:

```
https://example.com/url/to/amp.html.htxg
```

The frontend server should then internally reverse-proxy such a request to
something like:

```
http://packager.internal/priv/doc?sign=https%3A%2F%2Fexample.com%2Furl%2Fto%2Famp.html
```

Let's break that down:

  `http://` The packager itself only knows how to serve HTTP. Connections to it
  should remain inside your network. If you want encryption-in-transit (for
  instance if the connection to the packager travels outside your network), you
  may place a TLS-terminating proxy in front of it.

  `packager.internal` This is the host (and optionally port) of the packager as
  known to the frontend server.

  `/priv/doc` This is a fixed string. The frontend server must rewrite
  the URL to start with this.

  `?sign=https%3A%2F%2Fexample.com%2Furl%2Fto%2Famp.html` The location that
  should appear in the browser's URL bar, URL-escaped for use in a query. This
  must be HTTPS, and must be on a domain that the packager's certificate can
  sign for. If the user hits Refresh on their browser, it will fetch from this
  URL. By default, the content for the package is fetched from this same URL
  anonymously (e.g. without a `Cookie` header). It may not contain a fragment.

##### Certificates

AMP Packages will contain a `cert-url` that indicates the certificate that can
be used to validate the package. The `cert-url` may be on any domain, and it may
be HTTP or HTTPS, but it will have a path of the form:

```
/amppkg/cert/blahblahblah
```

where `blahblahblah` is a base64 encoding of a hash of the public certificate.
You may optionally prefix such URLs' paths, via the config file. The frontend
server must internally reverse-proxy these requests to the packager (without the
custom prefix).

##### Validity maps

AMP Packages will also contain a `validity-url`. This must be on the same domain
as the signed URL. It will have a path of the form:

```
/amppkg/validity
```

#### Configure the packager

The packager needs to be set up to receive reverse-proxied requests from the
frontend as specified above. In addition, it:

  * Must not be accessible on the open internet (even by IP address). To do so
    removes defense-in-depth against allowing external parties to fetch
    arbitrary documents and sign them with different arbitrary URLs.
  * Must have a certificate/key pair for all the domains you wish to sign. If
    you want to sign for multiple domains with different certificates, then run
    different instances of the packager. We recommend using a different
    certificate/key pair from your normal web-serving traffic. See the example
    config file for details.
  * Must be able to make outgoing connections to `cdn.ampproject.org` on port
    443.
  * Should have its stdout redirected to a log file somewhere, probably rotated.
  * Should not run as superuser.

<!-- TODO(twifkak): Add instructions for getting an API key or service account,
     after the Transformer API is in place. Maybe make a script that automates
     it using gcloud. -->

Once you've chosen a setup that meets the above constraints, actual
configuration is fairly straightforward:

  1. `go get github.com/ampproject/amppackager/cmd/amppkg`
  2. Move the built `~/go/bin/amppkg` wherever you like.
  3. Create a packager config file; use `amppkg.example.toml` in this repo as a template.
  4. Launch with `/path/to/amppkg -config=/path/to/amppkg.toml >>/path/to/amppkg.log`.
  5. Set up log rotation for `amppkg.log`.

#### Test your config

  1. Run a [Chrome Canary](https://www.google.com/chrome/browser/canary.html),
     or if on Linux, a [nightly build of
     Chromium](https://www.chromium.org/getting-involved/download-chromium).
  2. Navigate to chrome://flags#enable-signed-http-exchange and enable that
     feature.
  3. Navigate to your AMP package URL (i.e. the `href` of your
     `<link rel="amppackage">` for a given page).
  4. Watch the URL transmogrify!

Optionally, you may pretend to be an AMP Cache:

  1. Use `wget` to download the package and save it as a `foo.htxg` file.
  2. `go run cmd/amppkg_test_cache/main.go -package=path/to/foo.htxg`
  3. Ensure the packager is still running; it's needed to serve the certificate.
  4. Visit `http://localhost:8000/` in the experimental Chromium.

#### How will these web packages be discovered?

The details of that are still being worked out. We have several alternatives in
mind and want to come up with an answer that's best for the web, including
crawlers, sites serving packages, sites not serving packages, and package
caches. If you have any constraints or suggestions, please comment on issue #5,
or feel free to reach out in private if needed.

### Limitations

Currently, the packager will refuse to sign any AMP documents larger than 4 MB.

The Go http.Client implementation doesn't support redirect chains that rely on
cookies being set and then read later in the chain; all requests are made
cookielessly.

To account for possible clock skew in user agents, the packager back-dates
packages by 24h, which means they effectively last only 6 days for most users.

This tool only packages AMP documents. To sign non-AMP documents, look at the
commandline tools on which this was based, at
https://github.com/WICG/webpackage/tree/master/go/signedexchange. This should be
seen as a reference implementation, and not a supported library. Proper usage
requires following the status of browser implementations and updating callers to
import the correct library snapshot.

## Local Transformer

The local transformer is a sub-library within the AMP Packager that transforms AMP HTML for security and performance improvements. These modifications are described in more detail [here](https://github.com/ampproject/amphtml/blob/master/spec/amp-cache-modifications.md).

> **WARNING**: This local transformer library is still a work-in-progress and not all transformations described in the link above are implemented.

### How to use
The local transformer can be used separately from the packager/signer.

1. `go get github.com/ampproject/amppackager/cmd/transform`
1. `$GOPATH/bin/transform -url "documentURL" /path/to/html`

For more help, `$GOPATH/bin/transform -h`
