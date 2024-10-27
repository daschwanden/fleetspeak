#!/bin/bash
cp -r spanner testing

cd testing

openssl ecparam -list_curves

# generate a private key for a curve
openssl ecparam -name prime256v1 -genkey -noout -out key.pem

# optional: generate corresponding public key
openssl ec -in key.pem -pubout -out public-key.pem

# create a self-signed certificate
openssl req -new -x509 -key key.pem -out cert.pem -days 365 -subj "/C=AU/CN=fleetspeak-frontend" -addext "subjectAltName = DNS:fleetspeak-frontend"

FRONTEND_CERTIFICATE=$(sed ':a;N;$!ba;s/\n/\\\\n/g' cert.pem)
FRONTEND_KEY=$(sed ':a;N;$!ba;s/\n/\\\\n/g' key.pem)

sed -i 's@FRONTEND_CERTIFICATE@'"$FRONTEND_CERTIFICATE"'@' ./config/fleetspeak-client/config.textproto

sed -i 's@FRONTEND_CERTIFICATE@'"$FRONTEND_CERTIFICATE"'@' ./config/fleetspeak-server/components.textproto

sed -i 's@FRONTEND_KEY@'"$FRONTEND_KEY"'@' ./config/fleetspeak-server/components.textproto

cd ..
