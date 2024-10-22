#!/bin/zsh

#generate ca-keys
openssl req -x509 -sha256 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=TR/ST=EURASIA/L=ISTANBUL/O=Software/OU=yifeistudio/CN=*.yifeistudio.com/emailAddress=develop@yifeistudio.com" -nodes

#verify ca keys
openssl x509 -in ca-cert.pem -noout -text

#generate server key
openssl req -newkey rsa:2096 -keyout server-key.pem -out server-req.pem -subj "/C=ZH/ST=ZHEJIANG/L=HANGZHOU/O=patrol/OU=OrderService/CN=*.yifeistudio.com/emailAddress=develop@yifeistudio.com" -nodes -sha256

#sign
openssl x509 -req -in server-req.pem -days 60 -CA  ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf -sha256
