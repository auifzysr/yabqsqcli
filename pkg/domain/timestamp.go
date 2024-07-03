package domain

import (
	"fmt"
	"time"
)

const (
	timeLayout = time.RFC3339
)

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
