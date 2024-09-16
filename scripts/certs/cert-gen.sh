#!/bin/bash

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout ${CERT_PATH}/server-key.pem -out ${CERT_PATH}/server-req.pem -subj "/C=IN/ST=KA/L=Bangalore/O=kinara/OU=kinops/CN=*.kinops.kinara.com/emailAddress=support-kinops@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in ${CERT_PATH}/server-req.pem -days 60 -CA ${CA_PATH}/ca-cert.pem -CAkey ${CA_PATH}/ca-key.pem -CAcreateserial -out ${CERT_PATH}/server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in ${CERT_PATH}/server-cert.pem -noout -text