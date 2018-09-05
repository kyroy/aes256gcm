# aws256gcm

A simple tool to encrypt and decrypt aes-256-gcm with key and nonce.

Install
```bash
go get -u github.com/kyroy/aes256gcm
```

Encrypt
```bash
aes256gcm enc -in plain.file -out encrypted.file -key 2e3a40d4fef6b0fcfd664a1424aebac2cdd80df0af6d74fc5d44df84d3255718 -nonce bb8ef84243d2ee95a41c6c57
```

Decrypt
```bash
aes256gcm dec -in encrypted.file -out plain.file -key 2e3a40d4fef6b0fcfd664a1424aebac2cdd80df0af6d74fc5d44df84d3255718 -nonce bb8ef84243d2ee95a41c6c57
```
