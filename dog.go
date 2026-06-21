package dog

import "strings"

func When_grown(s string) string {
	return "When the puppy grows up it barks" + strings.ToUpper(s)
}