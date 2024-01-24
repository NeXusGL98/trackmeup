package interval_test

import (
	"errors"
	"slices"
	"testing"
	"time"

	"github.com/NexusGL98/trackmeup/internal/domain/interval"
)

func TestTimeInterval(t *testing.T) {

	t.Run("Should create a new TimeInterval on Valid start and end dates", func(t *testing.T) {

		startTimeToStr := "2020-01-01"
		endTimeToStr := "2020-01-31"

		ti, err := interval.NewTimeInterval(startTimeToStr, endTimeToStr, []time.Weekday{})

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		fromFormatted := ti.From.Format(time.DateOnly)
		toFormatted := ti.To.Format(time.DateOnly)

		if fromFormatted != startTimeToStr {
			t.Errorf("Expected From to be %s, got %s", startTimeToStr, fromFormatted)
		}

		if toFormatted != endTimeToStr {
			t.Errorf("Expected To to be %s, got %s", endTimeToStr, toFormatted)
		}

		if len(ti.Skip) != 0 {
			t.Errorf("Expected Skip to be empty, got %v", ti.Skip)
		}

	})

	t.Run("Should return an error on invalid start date", func(t *testing.T) {

		startTimeToStr := "2020-01-32"
		endTimeToStr := "2020-01-31"

		_, err := interval.NewTimeInterval(startTimeToStr, endTimeToStr, []time.Weekday{})

		if err == nil {
			t.Error("Expected an error, got nil")
		}

		var parseErr *time.ParseError

		if !errors.As(err, &parseErr) {
			t.Errorf("Expected a time.ParseError, got %v", err)
		}

	})

	t.Run("Should return an error on invalid end date", func(t *testing.T) {

		startTimeToStr := "2020-01-31"
		endTimeToStr := "2020-01-40"

		_, err := interval.NewTimeInterval(startTimeToStr, endTimeToStr, []time.Weekday{})

		if err == nil {
			t.Error("Expected an error, got nil")
		}

		var parseErr *time.ParseError

		if !errors.As(err, &parseErr) {
			t.Errorf("Expected a time.ParseError, got %v", err)
		}

	})

	t.Run("Should return an error when invalid date ranges", func(t *testing.T) {

		startTimeToStr := "2020-01-25"
		endTimeToStr := "2020-01-20"

		_, err := interval.NewTimeInterval(startTimeToStr, endTimeToStr, []time.Weekday{})

		if err == nil {
			t.Error("Expected an error, got nil")
		}

		if err != interval.ErrInvalidTimeRange {
			t.Errorf("Expected ErrInvalidTimeRange, got %v", err)
		}

	})

}

func TestTimeIntervalGenerateInterval(t *testing.T) {

	type IntervalCheck struct {
		interval.TimeInterval
		expected []time.Time
	}

	tableTests := map[string]IntervalCheck{
		"Should generate a valid interval with no skipping days from 2020-01-01 to 2020-01-05": {
			TimeInterval: interval.TimeInterval{
				From: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
			},
			expected: []time.Time{
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
			},
		},

		"Should generate a valid interval with skipping days from 2024-01-24 to 2024-01-28": {
			TimeInterval: interval.TimeInterval{
				From: time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2024, 1, 28, 0, 0, 0, 0, time.UTC),
				Skip: []time.Weekday{time.Saturday, time.Sunday},
			},

			expected: []time.Time{
				time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 25, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 1, 26, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for name, icheck := range tableTests {

		t.Run(name, func(t *testing.T) {

			interval, err := icheck.GenerateInterval()

			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}

			if len(interval) != len(icheck.expected) {
				t.Errorf("Expected interval to have 5 days, got %d", len(interval))
			}

			if !slices.Equal(interval, icheck.expected) {
				t.Errorf("Expected interval to be %v, got %v", icheck.expected, interval)
			}

		})

	}

}
