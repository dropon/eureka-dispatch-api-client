package openapi

import (
	"strings"
	time "time"
)

type Time struct {
	time.Time
}

const timeLayout = "2006-01-02T15:04:05"

// brusselsLocation is the timezone assumed for datetimes the API returns
// without timezone information. Per the API docs, such datetimes are stored
// without an offset and depend on context; Europe/Brussels matches the
// carrier's local time.
var brusselsLocation = func() *time.Location {
	loc, err := time.LoadLocation("Europe/Brussels")
	if err != nil {
		return time.UTC
	}
	return loc
}()

// MarshalJSON encodes the time without timezone information. The API ignores
// any offset on input, so the timezone-less layout is the safest choice.
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.Time.Format(timeLayout) + "\""), nil
}

// UnmarshalJSON decodes datetimes returned by the API, which may include a
// timezone ("2026-05-11T10:00:00Z", "...+05:00") or omit it entirely
// ("2026-05-11T10:00:00"). Offset-aware values are parsed with RFC3339;
// offset-less values are interpreted in Europe/Brussels.
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "" || s == "null" {
		return nil
	}

	if parsed, err := time.Parse(time.RFC3339, s); err == nil {
		t.Time = parsed
		return nil
	}

	parsed, err := time.ParseInLocation(timeLayout, s, brusselsLocation)
	if err != nil {
		return err
	}

	t.Time = parsed
	return nil
}
