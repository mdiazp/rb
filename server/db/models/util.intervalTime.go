package models

import "time"

// IntervalTime ...
type IntervalTime struct {
	InitialTime time.Time `gorm:"column:initial_time; not null"`
	FinishTime  time.Time `gorm:"column:finish_time; not null"`
}

// Valid ...
func (o *IntervalTime) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateTrue(&es, "IntervalTime",
		o.InitialTime.Equal(o.FinishTime) || o.InitialTime.Before(o.FinishTime),
	)
	return &es
}

// GetMinTimeValue ...
func GetMinTimeValue() time.Time {
	mn, _ := time.Parse(
		"2006-01-02T15:04:05Z07:00",
		"1900-01-01T00:00:00Z",
	)

	return mn
}

// GetMaxTimeValue ...
func GetMaxTimeValue() time.Time {
	mx, _ := time.Parse(
		"2006-01-02T15:04:05Z07:00",
		"2100-12-31T23:59:59Z",
	)

	return mx
}
