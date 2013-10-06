package main

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

var readyMsg = hash{"type": "ready"}
var errorMsg = hash{"type": "error"}
var goPongYourselfMsg = hash{"type": "go.pong.yourself"}
