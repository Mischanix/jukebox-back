package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func getSessionId() string {
	buf := make([]byte, 10)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(err)
	}
	result := base64.StdEncoding.EncodeToString(buf)
	if UserWithSid(result) != nil {
		return getSessionId()
	}
	return result
}

func makeSessionSecret() string {
	buf := make([]byte, 512)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(err)
	}
	result := base64.StdEncoding.EncodeToString(buf)
	return result
}

type Session struct {
	Id     string
	OldId  string
	Secret string
}
