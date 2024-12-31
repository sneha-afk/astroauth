# astroauth

## Generating keys
The following shows basic/default settings to make these keys, preferably add passphrases and more bytes for security.

### RSA
```bash
openssl genrsa -out priv_key.pem 2048
openssl rsa -in priv_key.pem -outform PEM -pubout -out pub_key.pem
```

### ECDSA
To work with the `x509` library, use the `.der` format for ECDSA keys:
```bash
openssl ecparam -genkey -name secp384r1 -outform der -noout -out priv_ecdsa.der
openssl pkey -in priv_ecdsa.der -pubout > pub_ecdsa.der
```
