package main

import (
	"strconv"
)

type hash map[string]interface{}

func (h hash) stringOrEmpty(key string) string {
	switch val := h[key].(type) {
	case string:
		return val
	default:
		return ""
	}
}

func (h hash) boolOrFalse(key string) bool {
	switch val := h[key].(type) {
	case bool:
		return val
	default:
		return false
	}
}

func (h hash) int64OrZero(key string) int64 {
	switch val := h[key].(type) {
	case string:
		n, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0
		} else {
			return n
		}
	default:
		return 0
	}
}

var readyMsg = hash{"type": "ready"}
var errorMsg = hash{"type": "error"}
var goPongYourselfMsg = hash{"type": "go.pong.yourself"}
