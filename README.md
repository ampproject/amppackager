# AMP Packager

> **WARNING**: This code is still brand new, and highly experimental. Feel free
> to play around with it, but be cautious. Examine the code.

The AMP Packager creates "AMP Packages" (implemented as [Signed HTTP
Exchanges](https://wicg.github.io/webpackage/draft-yasskin-httpbis-origin-signed-exchanges-impl.html))
containing AMP documents. These packages are consumed by the [Google AMP
Cache](https://www.ampproject.org/docs/fundamentals/how_cached), cached, and
when available, linked to from Google Search instead of normal AMP page. The
packages are signed with a certificate, which means that Chrome users will see
that certificate's domain in the URL bar instead of `google.com` or
`ampproject.org`.

These packages have a maximum lifetime of 7 days, which minimizes the risk of
another cache serving a stale copy of this signed content. In the future, the
[AMP update-cache API](https://developers.google.com/amp/cache/update-cache)
will support updating AMP Packages, though it doesn't yet.

The packager is an HTTP server that sits behind your frontend server; it
fetches and signs AMP documents as requested by the AMP Cache.

## How to use

### Add <link> tags to your pages

In order for Google Search to discover these web packages, you must link to them
from the page that googlebot crawls (whether that page is AMP or non-AMP). Do so
like:

```html
<link rel="amppackage" href="https://example.com/url/to/amp/package.html">
```

The URL must be HTTPS, and it should be absolute.

Your frontend server must know how to convert these package URLs back into their
corresponding AMP URLs, so make it a static transformation of such—for instance,
prepending `/htxg/` or appending `.htxg` to the path:

```html
<link rel="amphtml"    href="https://example.com/url/to/amp.html">
<link rel="amppackage" href="https://example.com/htxg/url/to/amp.html">
```

### Configure your frontend server

The frontend server needs to forward two types of requests to the packager:
packages and certificates.

#### Packages

For URLs that look like `https://example.com/htxg/url/to/amp.html`, the frontend
server must internally reverse-proxy these requests to something like:

```
http://packager.internal/priv/doc?fetch=https%3A%2F%2example.com%2Furl%2Fto%2Famp.html&sign=https%3A%2F%2Fexample.com%2Furl%2Fto%2Famp.html
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

  `?fetch=https%3A%2F%2Fexample.com%2Furl%2Fto%2Famp.html` The location of the
  AMP document to package, URL-escaped for use in a query. The same URL
  transformation that you applied to the `<link>` tag should be reversed by the
  web server. The packager will instruct the AMP CDN to fetch this URL
  anonymously (e.g. without a `Cookie` header). This URL can be HTTP or HTTPS,
  though the latter is strongly encouraged. The URL must be visible on the open
  internet.

  `&sign=https%3A%2F%2Fexample.com%2Furl%2Fto%2Famp.html` The location that
  should appear in the browser's URL bar, URL-escaped for use in a query. This
  must be HTTPS, and must be on a domain that the packager's certificate can
  sign for. If the user hits Refresh on their browser, it will fetch from this
  URL, so it must contain the same content as the fetch URL. Like the fetch URL,
  the frontend server will need to statically derive this URL from the
  amppackage URL.

#### Certificates

AMP Packages will contain a `certUrl` that indicates the certificate that can be
used to validate the package. The `certUrl` may be on any domain, and it may be
HTTP or HTTPS, but it will have a path of the form:

```
/amppkg/cert/blahblahblah
```

where `blahblahblah` is a base64 encoding of a hash of the public certificate.
You may optionally prefix such URLs' paths, via the config file. The frontend
server must internally reverse-proxy these requests to the packager (without the
custom prefix).

### Configure the packager

The packager needs to be set up to receive reverse-proxied requests from the
frontend as specified above. In addition, it:

  * Must not be accessible on the open internet (even by IP address). To do so
    would allow external parties to fetch arbitrary documents and sign them with
    different arbitrary URLs. (We provide some mitigation of this in the config
    file, via `URLSet`s.)
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

  1. `git clone http://github.com/ampproject/amp-packager && cd amp-packager && go build cmd/amppkg/main.go`
  2. Move the built `./amppkg` wherever you like.
  3. Create a packager config file; use `amppkg.example.toml` in this repo as a template.
  4. Launch with `/path/to/amppkg -config=/path/to/amppkg.toml >>/path/to/amppkg.log`.
  5. Set up log rotation for `amppkg.log`.

### Test your config

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

## Limitations

Currently, the packager will refuse to sign any AMP documents larger than 4 MB.

The Go http.Client implementation doesn't support redirect chains that rely on
cookies being set and then read later in the chain; all requests are made
cookielessly.

To account for possible clock skew in user agents, the packager back-dates
packages by 24h, which means they effectively last only 6 days for most users.
