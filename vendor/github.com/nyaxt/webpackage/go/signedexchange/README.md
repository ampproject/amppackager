# gen-signedexchange

This directory holds a fork of the [gen-signedexchange tool](https://github.com/WICG/webpackage/tree/master/go/signedexchange) that best matches the current Chrome implementation.

## Install
Simply go get:
```
go get github.com/nyaxt/webpackage/go/signedexchange/cmd/gen-{certurl,signedexchange}
```

## Basic Usage
Suppose you want to create a signed exchange envelope `foo.sxg` which encapsulates https://example.com/article.html.

First, prepare a private key `key.pem` and its X509 certificate `cert.pem` for the domain example.com encoded in PEM format. For now, the private key must be a RSA 2048-bit private key, and the certificate must be using the SHA-256 signing algorithm.

First, convert the certificate to certUrl format using gen-certurl:
```
gen-certurl cert.pem > cert.pem.msg
```

Then, host the certUrl at a public URL. In this example, let's suppose we hosted the `cert.pem.msg` at https://cert.example.org/cert.pem.msg.

Finally, using the key pair and `cert.pem.msg`, generate the signed exchange envelope using gen-signedexchange:
```
gen-signedexchange \
  -uri https://example.com/article.html \
  -content ./article.html \
  -certificate ./cert.pem \
  -certUrl https://cert.example.org/cert.pem.msg \
  -validityUrl https://cert.example.org/resource.validity.msg \
  -privateKey ./key.pem \
  -o ./foo.sxg
```

The `validityUrl` is not currently being fetched from Chrome. For now any URL should work, as long as it is a valid URL.
