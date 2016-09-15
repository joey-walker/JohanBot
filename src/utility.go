package main

import "strings"


// example: www.go ogle.com -> www.go%20ogle.com
func convertSpacesToNbsp(incomingURL string) string {
	updatedURL := strings.Replace(incomingURL, " ", "%20", -1)
	return updatedURL
}
