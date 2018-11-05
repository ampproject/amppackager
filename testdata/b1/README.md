# Test certificates for b1

This is an example certificate built under the constraints set by `v=b1` (and
possibly also applicable to higher versions).

To generate:

```
$ openssl genrsa -out ca.privkey 2048
$ openssl req -x509 -new -nodes -key ca.privkey -sha256 -days 1825 -out ca.cert -subj '/C=US/ST=California/O=Google LLC/CN=Fake CA'
$ openssl ecparam -out server.privkey -name prime256v1 -genkey
$ openssl req -new -sha256 -key server.privkey -out server.csr -subj /CN=example.com
$ openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.privkey -CAcreateserial -out server.cert -days 3650 \
 -extfile <(echo -e "1.3.6.1.4.1.11129.2.1.22 = ASN1:NULL")
$ cat server.cert ca.cert >fullchain.cert
```

## OCSP

To regenerate the OCSP signing certificate:

```
$ openssl req -new -sha256 -key ca.privkey -out ocsp.csr -subj '/C=US/ST=California/O=Google LLC/CN=ocsp.example.com'
$ openssl x509 -req -in ocsp.csr -signkey ca.privkey -out ca.ocsp.cert \
  -extfile <(echo -e "keyUsage = critical, digitalSignature\nextendedKeyUsage = critical, OCSPSigning\n")

```

Generate an OCSP response to seed the cache. Note this step must be done
manually on demand because of the 7 day expiry. DO NOT commit the generated
file.


```
$ ./seedcache.sh
```


### Appendix

<!--
TODO(twifkak): Update this to add CanSignHttpExchanges extension.
TODO(twifkak): Update this to add AIA for OCSP.
https://www.feistyduck.com/library/openssl-cookbook/online/ch-openssl.html
https://github.com/grimm-co/GOCSP-responder
https://github.com/OpenVPN/easy-rsa
https://gist.github.com/NoMan2000/06fffaca2ea710175cbcdd1a933c44af
-->

See some tutorials:
 - https://github.com/WICG/webpackage/tree/master/go/signedexchange
 - https://deliciousbrains.com/ssl-certificate-authority-for-local-https-development/
 - https://github.com/jmmcatee/cracklord/wiki/Creating-Certificate-Authentication-From-Scratch-OpenSSL
 - https://gist.github.com/Soarez/9688998
 - https://jamielinux.com/docs/openssl-certificate-authority/online-certificate-status-protocol.html
