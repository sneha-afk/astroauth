#!/bin/bash

# https://auth0.com/blog/brute-forcing-hs256-is-possible-the-importance-of-using-strong-keys-to-sign-jwts/

PRIV_LOC="private_key.pem"
PUB_LOC="public_key.pem"

# Generate a private key
openssl genpkey -algorithm RSA -out $PRIV_LOC -pkeyopt rsa_keygen_bits:2048

# Derive the public key from the private key
openssl rsa -pubout -in $PRIV_LOC -out $PUB_LOC
