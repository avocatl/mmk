package bm

import "net/url"

func escapePath(v string) string {
	return url.PathEscape(v)
}
