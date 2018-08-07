# aws256gcm

A simple tool to encrypt and decrypt aes-256-gcm with key and nonce.

Install
```bash
go get -u github.com/kyroy/aes256gcm
```

Encrypt
```bash
aes256gcm -in plain.file -out encrypted.file -key AES256Key-32Characters1234567890 -nonce bb8ef84243d2ee95a41c6c57 enc
```

Decrypt
```bash
aes256gcm -in encrypted.file -out plain.file -key AES256Key-32Characters1234567890 -nonce bb8ef84243d2ee95a41c6c57 dec
```
