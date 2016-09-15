package main

import "strings"

func convertSpacesToNbsp(incomingURL string) string {
	updatedURL := strings.Replace(incomingURL, " ", "%20", -1)
	return updatedURL
}
