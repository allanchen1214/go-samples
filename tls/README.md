# Generate private key (.key)

```shell
# Key considerations for algorithm "RSA" ≥ 2048-bit
openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
# https://safecurves.cr.yp.to/
# List ECDSA the supported curves (openssl ecparam -list_curves)
openssl ecparam -genkey -name secp384r1 -out server.key
```

# Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
```shell
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```


### Knowledge
.crt — Alternate synonymous most common among *nix systems .pem (pubkey).
.csr — Certficate Signing Requests (synonymous most common among *nix systems).
.cer — Microsoft alternate form of .crt, you can use MS to convert .crt to .cer (DER encoded .cer, or base64[PEM] encoded .cer).
.pem = The PEM extension is used for different types of X.509v3 files which contain ASCII (Base64) armored data prefixed with a «—– BEGIN …» line. These files may also bear the cer or the crt extension.
.der — The DER extension is used for binary DER encoded certificates.