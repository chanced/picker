package search

import (
	"github.com/chanced/dynamic"
)

type WithTimeZone interface {
	TimeZone() string
	SetTimeZone(v string)
}

// timeZoneParam is a mixin that adds the time_zone param to queries
//
// Coordinated Universal Time (UTC) offset or IANA time zone used to convert
// date values in the query to UTC.
//
// Valid values are ISO 8601 UTC offsets, such
// as +01:00 or -08:00, and IANA time zone IDs, such as America/Los_Angeles.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html#range-query-time-zone
type timeZoneParam struct {
	timeZone string
}

func (tz timeZoneParam) TimeZone() string {
	return tz.timeZone
}
func (tz *timeZoneParam) SetTimeZone(v string) {
	tz.timeZone = v
}
func unmarshalTimeZoneParam(data dynamic.RawJSON, target interface{}) error {
	if a, ok := target.(WithTimeZone); ok {
		a.SetTimeZone(data.UnquotedString())
	}
	return nil
}
func marshalTimeZoneParam(data dynamic.Map, source interface{}) (dynamic.Map, error) {
	if b, ok := source.(WithTimeZone); ok {
		if b.TimeZone() != "" {
			data[paramTimeZone] = b.TimeZone()
		}
	}
	return data, nil
}