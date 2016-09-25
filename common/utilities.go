package common

import "strings"

// example: www.go ogle.com -> www.go%20ogle.com
func ConvertSpacesToNbsp(incomingURL string) string {
	updatedURL := strings.Replace(incomingURL, " ", "%20", -1)
	return updatedURL
}
