openssl ecparam -out packager.key -name prime256v1 -genkey

openssl req -new -key packager.key -nodes -out packager.csr

openssl req -in packager.csr -noout -text