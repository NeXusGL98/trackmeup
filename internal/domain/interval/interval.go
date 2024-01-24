package interval

import (
	"errors"
	"time"
)

// Intervaler represents an interface for generating time intervals.
type Intervaler interface {
	GenerateInterval() ([]time.Time, error)
}

var (
	ErrInvalidTimeRange = errors.New("invalid time range")
)
