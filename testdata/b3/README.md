# Test certificates for b3

This is example certificates for tests built under the constraints set by `v=b3` except having AIA and SCT extension.

To generate:

CA private key and cert,
```
$ openssl genrsa -out ca.privkey 2048
$ openssl req -x509 -new -nodes -key ca.privkey -sha256 -days 1825 -out ca.cert -subj '/C=US/ST=California/O=Google LLC/CN=Fake CA'
```

server.privkey and server.cert of amppackageexample.com and www.amppackageexample.com,
```
$ openssl ecparam -out server.privkey -name prime256v1 -genkey
$ openssl req -new -sha256 -key server.privkey -out server.csr -subj /CN=amppackageexample.com
$ openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.privkey -CAcreateserial -out server.cert -days 90  -extfile <(echo -e "subjectAltName = DNS:amppackageexample.com,DNS:www.amppackageexample.com\n1.3.6.1.4.1.11129.2.1.22 = ASN1:NULL")
$ cat server.cert ca.cert > fullchain.cert

server2.privkey and server2.cert of amppackageexample2.com and www.amppackageexample2.com,
```
$ openssl ecparam -out server2.privkey -name prime256v1 -genkey
$ openssl req -new -sha256 -key server2.privkey -out server2.csr -subj /CN=amppackageexample2.com
$ openssl x509 -req -in server2.csr -CA ca.cert -CAkey ca.privkey -CAcreateserial -out server2.cert -days 90  -extfile <(echo -e "subjectAltName = DNS:amppackageexample2.com,DNS:www.amppackageexample2.com\n1.3.6.1.4.1.11129.2.1.22 = ASN1:NULL")
$ cat server2.cert ca.cert > fullchain2.cert
```

and others
```
$ openssl x509 -req -in server.csr -CA ca.cert -CAkey ca.privkey -CAcreateserial -out server_91days.cert -days 91  -extfile <(echo -e "subjectAltName = DNS:amppackageexample.com,DNS:www.amppackageexample.com\n1.3.6.1.4.1.11129.2.1.22 = ASN1:NULL")
$ cat server_91days.cert ca.cert > fullchain_91days.cert
$ openssl ecparam -out server_p521.privkey -name secp521r1 -genkey
```

### Appendix

<!--
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
