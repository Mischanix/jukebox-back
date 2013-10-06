package main

func userMessage(nick string, quarters, skipCost float64) hash {
	return hash{
		"type":     "user",
		"nick":     nick,
		"quarters": quarters,
		"skipCost": skipCost,
	}
}
