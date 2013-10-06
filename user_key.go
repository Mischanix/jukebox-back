package main

import (
	"code.google.com/p/go.crypto/scrypt"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// NewKey creates a new salt and key for the DbUser given a string password. If
// successful, the new key and salt will be present in their base64-encoded
// fields on the DbUser.
func (dbUser *DbUser) NewKey(pass string) {
	salt := make([]byte, 512)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}
	key, err := scrypt.Key([]byte(pass), salt, 16384, 8, 1, 512)
	if err != nil {
		panic(err)
	}
	dbUser.Base64Salt = base64.StdEncoding.EncodeToString(salt)
	dbUser.Base64Key = base64.StdEncoding.EncodeToString(key)
}

func (dbUser *DbUser) TestKey(pass string) bool {
	salt, err := base64.StdEncoding.DecodeString(dbUser.Base64Salt)
	if err != nil {
		panic(err)
	}
	key, err := scrypt.Key([]byte(pass), salt, 16384, 8, 1, 512)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(key) == dbUser.Base64Key
}
