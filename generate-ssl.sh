#!/bin/bash

# Create a new directory to store the SSL certificates and keys
mkdir -p /ssl

# Generate a private key
openssl genrsa -out /ssl/server.key 2048

# Generate a certificate signing request (CSR)
openssl req -new -key /ssl/server.key -out /ssl/server.csr -subj "/CN=localhost"

# Generate a self-signed SSL certificate
openssl x509 -req -days 365 -in /ssl/server.csr -signkey /ssl/server.key -out /ssl/server.crt
