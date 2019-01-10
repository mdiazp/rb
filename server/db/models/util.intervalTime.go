package models

import "time"

// IntervalTime ...
type IntervalTime struct {
	InitialTime time.Time `gorm:"column:initial_time; type: date; not null"`
	FinishTime  time.Time `gorm:"column:finish_time; type: date; not null"`
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
		"2006-01-02",
		"1900-01-01",
	)

	return mn
}

// GetMaxTimeValue ...
func GetMaxTimeValue() time.Time {
	mx, _ := time.Parse(
		"2006-01-02",
		"2100-12-31",
	)

	return mx
}
