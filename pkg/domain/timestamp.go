package domain

import (
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02T15:04:05Z"
)

// ScheduleOptions' comment:
// Represents seconds of UTC time since Unix epoch
// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
// 9999-12-31T23:59:59Z inclusive.
func TimestampSeconds(timeRepresentation string) (int64, error) {
	if timeRepresentation == "" {
		return 0, fmt.Errorf("time representation string cannot be empty")
	}
	if timeRepresentation == "now" {
		return time.Now().Unix(), nil
	}

	parsedTime, err := time.Parse(timeLayout, timeRepresentation)
	if err != nil {
		return 0, fmt.Errorf("error parsing time: %w", err)
	}

	return parsedTime.Unix(), nil
}
