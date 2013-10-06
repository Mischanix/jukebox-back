package main

import (
	cryprand "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"io"
	"math/rand"
	"time"
)

var sessionIndex int64
var sessionModifier int64

func init() {
	rand.Seed(time.Now().UnixNano())
	sessionModifier = rand.Int63()
}

func getSessionId() string {
	var result string
	sessionIndex++
	buf := make([]byte, 10)
	binary.PutVarint(buf, sessionIndex^sessionModifier)
	result = base64.StdEncoding.EncodeToString(buf)
	return result
}

func makeSessionSecret() string {
	buf := make([]byte, 512)
	_, err := io.ReadFull(cryprand.Reader, buf)
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
