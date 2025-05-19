package dog

import (
	"strings"
)

// WhenGrownUp returns an uppercase message based on the input string
func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
