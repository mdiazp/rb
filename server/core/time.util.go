package core

import (
	"time"

	"github.com/mdiazp/rb/server/db/models"
)

// MoveDateToNextWeekDay ...
func MoveDateToNextWeekDay(date time.Time, wd models.WeekDay) time.Time {
	for true {
		if date.Weekday().String() == (string)(wd) {
			break
		}
		date.AddDate(0, 0, 1)
	}
	return date
}

// MoveDateToPreviousWeekDay ...
func MoveDateToPreviousWeekDay(date time.Time, wd models.WeekDay) time.Time {
	for true {
		if date.Weekday().String() == (string)(wd) {
			break
		}
		date.AddDate(0, 0, -1)
	}
	return date
}
