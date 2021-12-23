package bm

import (
	"net/url"
	"strings"
	"time"
)

func escapePath(v string) string {
	return url.PathEscape(v)
}

type offerTime struct {
	time.Time
}

func (ot *offerTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	ss := strings.Trim(s, "\"")
	t, err := time.Parse("2006-01-02 15:04:05", ss)
	if err != nil {
		return err
	}
	ot.Time = t
	return nil
}
