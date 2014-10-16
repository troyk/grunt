package sanitize

import (
	"regexp"
	"strings"
)

var (
	reWebsite = regexp.MustCompile(`^http://|^https://|/$|\?.*|[^0-9a-z-_\.]`)
	rePhone   = regexp.MustCompile(`[^0-9]`)
)

// Removes all protocol and query from a url and returns just the hostname
func Website(s string) string {
	return reWebsite.ReplaceAllString(strings.ToLower(strings.TrimSpace(s)), "")
}

// Nothing fancy, trimmed-lowered string
func Email(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

// Removes anything not a digit and attempts to normalize 10 digits
// e.g. 1-555-212-6969 becomes 5552226969
func Phone(s string) string {
	s = rePhone.ReplaceAllString(s, "")
	if len(s) == 11 && s[0:1] == "1" { // strip us country codes
		return s[1:]
	}
	return s
}
