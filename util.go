package binance

import "strings"

// stripQueryParams takes in a URL path, and removes all query parameters that
// have been added on. This is used mainly to sanitize the path for metrics in
// order to prevent posting dynamic paths to prometheus.
func stripQueryParams(path string) string {
	index := strings.LastIndex(path, "?")
	if index == -1 {
		return path
	}

	return path[0:index]
}
