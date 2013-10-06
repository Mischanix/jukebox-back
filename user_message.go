package main

func userMessage(nick string, quarters, skipCost float64, fake bool) hash {
	return hash{
		"type":     "user",
		"nick":     nick,
		"fake":     fake,
		"quarters": quarters,
		"skipCost": skipCost,
	}
}
