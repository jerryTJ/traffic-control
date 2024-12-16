
# create ca

openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 3650 -out ca.crt -subj "/CN=jerry"

# server

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=traffic.devops.com"
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650

# client

openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 3650

# validate

openssl s_client -connect localhost:10080 -CAfile ca.crt

# use openssl  config

##

openssl genrsa -out server.key 2048

##

openssl req -new -key server.key -out server.csr -config openssl.cnf

##

openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650 -extensions v3_req -extfile openssl.cnf
