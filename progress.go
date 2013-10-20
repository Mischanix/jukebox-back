package main

func init() {
	handlers["progress"] = progressHandler
}

func progressHandler(c *Client, frame hash) {
	c.notifyProgress()
}
