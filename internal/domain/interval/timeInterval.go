package interval

import "time"

// TimeInterval represents a time interval from a start time to an end time.
// It receives skip days of the week to exclude from the interval.
type TimeInterval struct {
	From time.Time      `json:"from"`
	To   time.Time      `json:"to"`
	Skip []time.Weekday `json:"skip"`
}

// NewTimeInterval creates a new TimeInterval.
// It receives a start time, an end time as String and a list of days of the week to skip.

func NewTimeInterval(from, to string, skip []time.Weekday) (*TimeInterval, error) {
	var ti TimeInterval

	fromTime, err := time.Parse(time.DateOnly, from)

	if err != nil {
		return nil, err
	}

	toTime, err := time.Parse(time.DateOnly, to)

	if err != nil {
		return nil, err
	}

	// check if the start date is before the end date, otherwise return an error

	if fromTime.After(toTime) {
		return nil, ErrInvalidTimeRange
	}

	ti = TimeInterval{
		From: fromTime,
		To:   toTime,
		Skip: skip,
	}

	return &ti, nil

}

func (ti *TimeInterval) isSkipDay(day time.Weekday) bool {

	if len(ti.Skip) == 0 {
		return false
	}

	for _, d := range ti.Skip {
		if d == day {
			return true
		}
	}
	return false
}

func (ti *TimeInterval) GenerateInterval() []time.Time {

	var interval []time.Time

	// Loop through the days between the start and end date, including the end date.
	// If the day is not in the skip list, add it to the interval.

	for start := ti.From; start.Before(ti.To) || start.Equal(ti.To); start = start.AddDate(0, 0, 1) {
		if !ti.isSkipDay(start.Weekday()) {
			interval = append(interval, start)
		}
	}

	return interval
}
