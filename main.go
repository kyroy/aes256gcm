package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var (
		in, out    string
		key, nonce string
	)
	flag.StringVar(&in, "in", "", "input file")
	flag.StringVar(&out, "out", "", "output file")
	flag.StringVar(&key, "key", "", "key")
	flag.StringVar(&nonce, "nonce", "", "nonce")
	flag.Parse()
	if in == "" || out == "" || key == "" || nonce == "" {
		printUsage(fmt.Sprintf("in=%s,out=%s,key=%s,nonce=%s", in, out, key, nonce))
		os.Exit(1)
	}
	if flag.NArg() != 1 {
		printUsage("arg expected")
		os.Exit(1)
	}
	var f func(key, nonce string, ciphertext []byte) ([]byte, error)
	switch flag.Arg(0) {
	case "enc":
		log("encrypting %s to %s", in, out)
		f = encrypt
	case "dec":
		log("decrypting %s to %s", in, out)
		f = decrypt
	default:
		printUsage(fmt.Sprintf("unknown command: %s", flag.Arg(0)))
		os.Exit(1)
	}
	input, err := ioutil.ReadFile(in)
	if err != nil {
		log(err.Error())
		os.Exit(1)
	}
	output, err := f(key, nonce, input)
	if err != nil {
		log(err.Error())
		os.Exit(1)
	}
	if err := ioutil.WriteFile(out, output, os.ModePerm); err != nil {
		log(err.Error())
		os.Exit(1)
	}
}

func log(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format = format + "\n"
	}
	fmt.Fprintf(os.Stderr, format, args...)
}

func printUsage(msgs ...string) {
	if len(msgs) > 0 {
		log(strings.Join(msgs, "\n"))
	}
	log("Usage:")
	flag.PrintDefaults()
}

func decrypt(key, nonce string, ciphertext []byte) ([]byte, error) {
	aesgcm, err := parseSymmetricKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse symmetric key: %v", err)
	}
	n, err := hex.DecodeString(nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nonce: %v", err)
	}
	plaintext, err := aesgcm.Open(nil, n, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decode: %v", err)
	}
	return plaintext, nil
}

func encrypt(key, nonce string, payload []byte) ([]byte, error) {
	aesgcm, err := parseSymmetricKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse symmetric key: %v", err)
	}
	n, err := hex.DecodeString(nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nonce: %v", err)
	}
	return aesgcm.Seal(nil, n, payload, nil), nil
}

func parseSymmetricKey(key string) (cipher.AEAD, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return aesgcm, nil
}
