# astroauth

## Setting `.env`

In a `.env` at the root of the directory, set the following variables:
1. `PRIV_KEY_LOC`: file path to the private key
2. `PUBLIC_KEY_LOC`: file path to the public key
3. `AUTH_KEY_TYPE`: type of the key

The following authentication keys are supported:
- RSA
- ES256
- ES512

## Generating keys
The following shows basic/default settings to make these keys, preferably add passphrases and more bytes for security.

### RSA
```bash
openssl genrsa -out priv_key.pem 2048

openssl rsa -in priv_key.pem -outform PEM -pubout -out pub_key.pem
```

### ECDSA: ES256
```bash
openssl ecparam -name prime256v1 -genkey -noout -out priv_ecdsa.pem

openssl pkey -in priv_ecdsa.pem -pubout > pub_ecdsa.pem
```

### ECDSA: ES512
```bash
openssl ecparam -name secp521r1 -genkey -noout -out priv_ecdsa.pem

openssl ec -in priv_ecdsa.pem -pubout -out pub_ecdsa.pem
```
