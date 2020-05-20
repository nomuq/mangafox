package mangareader

import "strings"

func TrimAndSplitURL(u string) []string {
	u = strings.TrimSuffix(u, "/")
	return strings.Split(u, "/")
}

// IsURLValid will exclude those url containing `.gif` and `logo`.
func IsURLValid(value string) bool {
	check := value != "" && !strings.Contains(value, ".gif") && !strings.Contains(value, "logo") && !strings.Contains(value, "mobilebanner")

	if check {
		return strings.HasPrefix(value, "http") || strings.HasPrefix(value, "https")
	}

	return check
}
