package main

import (
	"time"
)

func init() {
	handlers["queue"] = queueHandler
}

func queueHandler(c *Client, frame hash) {
	trackId := queueMessage(frame)
	// validation and timing and shit
	track, err := trackInfo(trackId)
	if err != nil {
		c.sendQueue <- queueError("api error: " + err.Error())
		return
	}
	if track.duration > 15*time.Minute {
		c.sendQueue <- queueError("song too long")
		return
	}
	enqueueSong(track)
}
