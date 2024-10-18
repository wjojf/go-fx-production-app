package utils

import (
	"fmt"
	"regexp"
)

func ExtractAuthToken(authHeader string) (string, error) {
	// Regular expression to match only "Bearer {token}" (case-sensitive)
	re := regexp.MustCompile(`^Bearer\s+(.+)$`)

	// Find the submatch (the token)
	matches := re.FindStringSubmatch(authHeader)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid authorization format")
	}

	// Return the token (the second part of the match)
	return matches[1], nil
}
