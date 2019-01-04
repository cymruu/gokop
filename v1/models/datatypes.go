package models

import (
	"strings"
	"time"
)

type WykopTime struct {
	time.Time
}

func (wt *WykopTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	wt.Time, err = time.Parse("2006-01-02 15:04:05", s) //IMO its's really retearded idea to use actual date values as placeholders
	if err != nil {
		wt.Time, _ = time.Parse(time.RFC3339Nano, s)
	}
	return
}
