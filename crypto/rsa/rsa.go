package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GenKey(bits int) (prvKey, pubKey []byte, err error) {
	// Generates private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvKey = pem.EncodeToMemory(block)

	// Generates public key from private key.
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return
	}
	block = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	pubKey = pem.EncodeToMemory(block)
	return
}

func Encrypt(pubKey, data []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("decode public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), data)
}

func Decrypt(prvKey, cipher []byte) ([]byte, error) {
	block, _ := pem.Decode(prvKey)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	prv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, prv, cipher)
}

func Sign(prvKey []byte, hash crypto.Hash, data []byte) ([]byte, error) {
	block, _ := pem.Decode(prvKey)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// MD5 and SHA1 are not supported as they are not secure.
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	return rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
}

func VerifySign(pubKey []byte, hash crypto.Hash, data, sig []byte) error {
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return errors.New("decode public key error")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	// SHA1 and MD5 are not supported as they are not secure.
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), hash, hashed, sig)
}
