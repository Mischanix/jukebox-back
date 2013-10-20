package main

func queueMessage(frame hash) (track string) {
	return frame.stringOrEmpty("track")
}

func queueError(reason string) hash {
	return hash{
		"type":   "queue.error",
		"reason": reason,
	}
}
