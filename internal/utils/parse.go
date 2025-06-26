package utils

import (
	"github.com/mssola/user_agent"
)

// ParseUserAgent extracts browser name, device type, and bot detection
func ParseUserAgent(ua string) (browser string, device string, isBot bool) {
	parser := user_agent.New(ua)

	browserName, _ := parser.Browser()
	browser = browserName

	if parser.Mobile() {
		device = "mobile"
	} else {
		device = "desktop"
	}

	isBot = parser.Bot()

	return
}
