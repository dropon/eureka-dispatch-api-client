package openapi

import (
	"encoding/json"
	"strings"
	time "time"
)

type Time struct {
	time.Time
}

const timeLayout = "2006-01-02T15:04:05"

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return
	}

	loc, err := time.LoadLocation("Europe/Brussels")
	if err != nil {
		return err
	}

	parsedTime, err := time.ParseInLocation(timeLayout, s, loc)
	if err != nil {
		return err
	}

	t.Time = parsedTime
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}
