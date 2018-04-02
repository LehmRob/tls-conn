# tls-conn

Example project for showing how to run a TLS server and connect to it using
a TLS client. Both are written in go.

## Certificat generation

You can generate the certificate with the following command

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout privateKey.key -out certificate.crt
```

It's very important to use localhost or an other FQDN for the certificate creation.
