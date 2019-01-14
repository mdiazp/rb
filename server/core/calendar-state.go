package core

import (
	"time"

	"github.com/mdiazp/rb/server/db/models"
)

// PDRTurnCalendarState ...
type PDRTurnCalendarState struct {
	Date         time.Time
	PDRs         []*models.PDiskReservation
	DCRR         []DiscCategoryRequestReport
	DCRNullTotal int
	DiscsTotal   int
}

// DiscCategoryRequestReport ...
type DiscCategoryRequestReport struct {
	Category string
	DCTotal  int
	DCRTotal int
}
