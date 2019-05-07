# AMP Cache SXG Requirements

## Audience

The audience for this document is people intending on implementing their own AMP
signed exchange generator, independent of amppackager, and those implementing
their own AMP SXG cache for the purposes of privacy-preserving prefetch. Users
of `amppkg` need not read this, as the tool should automatically guarantee the
following requirements are met.

## Google AMP cache

The Google AMP cache sets some requirements in addition to the ones set by the
[SXG spec](https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html).
These include:

 * The signed `fallback URL` must equal the URL at which the SXG was delivered.
 * The payload must be:
   * non-empty.
   * valid transformed AMP. The canonical definition of transformed AMP is the
     return value of [`transform.Process()`](https://github.com/ampproject/amppackager/blob/e4bf0430ba152cfe82ccf063df92021dfc0f26a5/transformer/transformer.go#L219).
     If given a [valid AMP](https://github.com/ampproject/amphtml/tree/master/validator)
     doc as input, it should produce a valid transformed AMP doc. There may be
     other ways of achieving this, but they are unsupported (i.e. may
     arbitrarily break in the future).
   * unchanged after calling [`transform -config NONE`](https://github.com/ampproject/amppackager/tree/releases/transformer#how-to-use).
   * matching one of the versions requested by the `AMP-Cache-Transform` header.
     Note that this version range will increase over time, at a cadence TBD
     (likely 6-8 weeks with 2 or 3 supported latest versions).
 * If the signed `cache-control` header has a `no-cache` directive, it cannot
   have a value (i.e. `no-cache=some-header` is disallowed).
 * The signed `content-security-policy` header must be present and comply with
   these rules:
   * `default-src`, `script-src`, `object-src`, `style-src`, and `report-uri`
     must equal those from the [AMP cache CSP](https://github.com/ampproject/amppackager/blob/releases/packager/signer/signer.go#L272)
   * `base-uri`, `block-all-mixed-content`, `font-src`, `form-action`,
     `manifest-src`, `referrer`, and `upgrade-insecure-requests` may be omitted
     or have any value
   * all other directives are disallowed
 * The signed `link` header, if present, must look like [this](https://github.com/ampproject/amppackager/blob/e4bf0430ba152cfe82ccf063df92021dfc0f26a5/packager/signer/signer.go#L426)
   (the validation logic is currently very picky about its serialization); and
   have limits like [this](https://github.com/ampproject/amppackager/blob/e4bf0430ba152cfe82ccf063df92021dfc0f26a5/transformer/transformer.go#L177)
   (e.g. max 20 urls, rel=preload only, as=script|style only). URLs must be
   limited to `cdn.ampproject.org` and the allowlisted [font provider URLs](https://github.com/ampproject/amphtml/blob/b0ff92429923c86f3973009a84ff02f4f1868b4d/validator/validator-main.protoascii#L310).
 * The signature's duration (expiry minus date) must be >= 4 days.

The above is an attempt at a complete list of SXG-related requirements, but it
is not guaranteed to be complete.

Some of the above limitations are overly strict for an AMP SXG cache's needs,
and were implemented as such for the sake of expediency. They may be loosened
over time, especially in response to publisher feedback.

## Other AMP caches

As other AMP caches adopt support for signed exchanges, they could define their
own set of requirements. It would be most useful for publishers and users,
however, if the requirements were the same across all caches. If you see a need
for a different requirement on your cache, please contact the AMP Caching
working group, for example via [Slack](https://amphtml.slack.com/) on the
`#signed-exchanges` channel, via one of [these methods](https://github.com/ampproject/wg-caching#communication-channels).

## Testing

There is no known publicly available tool for validating an SXG against the
above requirements, though one is certainly welcome. In the interim, one may
issue a request against the Google AMP Cache and see if the response is a valid
SXG.

Meets requirements:
```
$ curl -s -i -H 'Accept: application/signed-exchange;v=b3' https://amppackageexample-com.cdn.ampproject.org/wp/s/amppackageexample.com/ | grep -a -i content-type:
content-type: application/signed-exchange;v=b3
```

Does not meet requirements:
```
$ curl -s -i -H 'Accept: application/signed-exchange;v=b3' https://amppackageexample-com.cdn.ampproject.org/wp/s/amppackageexample.com/gen/invalid.sxg | grep META
<META HTTP-EQUIV="refresh" content="1; url=https://amppackageexample.com/gen/invalid.sxg">
```
