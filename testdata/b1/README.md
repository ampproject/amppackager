# Test certificates for b1

This is an example certificate built under the constraints set by `v=b1` (and
possibly also applicable to higher versions).

To generate:

```
$ openssl genrsa -out ca.privkey 2048
$ openssl req -x509 -new -nodes -key ca.privkey -sha256 -days 1825 -out ca.cert -subj '/C=US/ST=California/O=Google LLC/CN=Fake CA'
$ openssl ecparam -out server.privkey -name prime256v1 -genkey
$ openssl req -new -sha256 -key server.privkey -out server.csr -subj /CN=example.com
$ openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.privkey -CAcreateserial -out server.cert -days 3650
$ cat server.cert ca.cert >fullchain.cert
```

OCSP:

Make sure your openssl.cfg has an ocsp extension section, like:

[ ocsp ]
# Extension for OCSP signing certificates (`man ocsp`).
basicConstraints = CA:FALSE
subjectKeyIdentifier = hash
authorityKeyIdentifier = keyid,issuer
keyUsage = critical, digitalSignature
extendedKeyUsage = critical, OCSPSigning

Generate a cert with ocsp extensions:
$ openssl req -x509 -new -nodes -key ca.privkey -sha256 -days 1825 -out ca.ocsp.cert -subj '/C=US/ST=California/O=Google LLC/CN=Fake CA' -extensions ocsp

Generate an OCSP request and write it to a file:
$ openssl ocsp -issuer ca.cert -cert server.cert -reqout ocspreq.der

Generate an OCSP response. Note this step must be done manually because of the 7
day expiry. Do not commit the generated file.
$ openssl ocsp -index ./index.txt -port 8888 -rsigner ca.ocsp.cert -rkey ca.privkey -CA ca.cert -ndays 7 -reqin ocspreq.der -respout ocspresp.der

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
