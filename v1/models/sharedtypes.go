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

type OK_ID struct {
	ID int64 `json:"id"`
}

type Sex string

const (
	Female Sex = "female"
	Male   Sex = "male"
)

type Usergroup uint16

var groupToColor = map[Usergroup]string{
	0:    "zielony",
	1:    "pomaranczowy",
	2:    "bordowy",
	5:    "administrator",
	1001: "zbanowany",
	1002: "usunięty",
	2001: "niebieski",
}

func (u *Usergroup) ToString() string {
	color, ok := groupToColor[*u]
	if ok {
		return color
	}
	return "unknown"
}
