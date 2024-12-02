生成 CA 证书和客户端证书
生成 CA 私钥和证书：

openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 3650 -out ca.crt -subj "/CN=MyCA"
使用 CA 签发服务端证书：

openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -config openssl.cnf
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650 -extensions v3_req -extfile openssl.cnf
生成客户端私钥和证书：

openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=client"
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365

验证证书生效
openssl s_client -connect grpc.devops.com:10080 -CAfile ca.crt

openssl verify -CAfile ca.crt server.crt
openssl verify -CAfile ca.crt client.crt
