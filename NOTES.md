openssl ecparam -out packager.key -name prime256v1 -genkey

openssl req -new -key packager.key -nodes -out packager.csr

openssl req -in packager.csr -noout -text

 openssl ecparam -out privkey.pem -name prime256v1 -genkey

 openssl req -new -key privkey.pem -nodes -out beebo-red.csr -subj "/CN=beebo.red"

 # after validation DigiCert will return a zip file containing beebo_red.crt and DigiCertCA.crt, just concatenate these together to create cert.pem

 cat beebo_red.crt DigiCertCA.crt > cert.pem

 The two files you need are privkey.pem and cert.pem

 grep getURL $GOPATH/src/github.com/ampproject/amppackager/packager/certcache/certcache.go

 getraw -H 'accept: application/signed-exchange;v=b2' -H 'amp-cache-transform: google' http://localhost:8080/priv/doc/https://beebo.red/

 getraw -H 'X-Appengine-Inbound-Appid: foo' -H 'accept: application/signed-exchange;v=b2' -H 'amp-cache-transform: google' http://localhost:8080/priv/doc/https://beebo.red/

/Applications/Google\ Chrome\ Canary.app/Contents/MacOS/Google\ Chrome\ Canary \
--user-data-dir=/tmp/udd \
--ignore-certificate-errors-spki-list=KTiy+EOYZly/v6Za1YH2VVJoL3ZeG3vLMNQvnTgF7UI= \
--enable-features=SignedHTTPExchange \
'data:text/html,<a href="https://beebo.blue/priv/doc/https://beebo.red/">click me</a>'

 gethead -H 'X-Appengine-Inbound-Appid: foo' -H 'accept: application/signed-exchange;v=b2' -H 'amp-cache-transform: google' https://stillers-blue.appspot.com/priv/doc/https://beebo.red/

data:text/html,<a href="https://localhost:8080/priv/doc/https://beebo.red/">click me</a>

getraw -H 'amp-cache-transform: google' -H 'accept: application/signed-exchange;v=b2;q=0.9,*/*;q=0.8' https://beebo.blue/priv/doc/https://beebo.red/

Deploy beebo.red:

cd amp-by-example
goapp deploy -application stillers-red -version 1

Deploy beebo.blue:

cd amp-by-example/packager
gcloud app deploy --project stillers-blue

# returns text/html
getraw https://beebo.red/

# returns application/signed-exchange
getraw -H 'amp-cache-transform: google' -H 'accept: application/signed-exchange;v=b2;q=0.9,*/*;q=0.8' https://beebo.red/