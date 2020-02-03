#!/bin/bash

# Seed the OCSP cache using the fake test certs. Do NOT run this if you
# are using real certificates with the AMP Packager.

openssl ocsp -index ./index.txt -rsigner ca.ocsp.cert -rkey ca.privkey -CA ca.cert -ndays 7 -issuer ca.cert -cert server.cert -respout /tmp/amppkg-ocsp
