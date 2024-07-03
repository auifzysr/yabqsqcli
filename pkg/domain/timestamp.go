package domain

import (
	"fmt"
	"time"
)

func TimestampSeconds(timeRepresentation string) (int64, error) {
	if timeRepresentation == "" {
		return 0, fmt.Errorf("time representation string cannot be empty")
	}
	if timeRepresentation == "now" {
		return time.Now().Unix(), nil
	}

	parsedTime, err := time.Parse(time.RFC3339, timeRepresentation)
	if err != nil {
		return 0, fmt.Errorf("Error parsing time: %w", err)
	}

	return parsedTime.Unix(), nil
}
