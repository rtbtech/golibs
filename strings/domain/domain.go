package domain

import (
	"strings"
)

func Name(uri string) string {
	if len(uri) > 0 {
		if pos := strings.Index(uri, "//"); pos != -1 {
			uri = uri[pos+2:]
		}

		if strings.HasPrefix(uri, "www.") {
			uri = uri[4:]
		}

		if pos := strings.Index(uri, "/"); pos != -1 {
			uri = uri[:pos]
		}
	}

	return strings.ToLower(uri)
}
