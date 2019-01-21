package core

import (
	"time"

	"github.com/mdiazp/rb/server/db/models"
)

// GetPDRFreeTurns return free turns from ini time untill end time
func GetPDRFreeTurns(pdrs []models.PDiskReservation,
	ini time.Time, end time.Time, dcr string) []models.Turn {

	return nil
}

// TurnAndDiscCategory ...
type TurnAndDiscCategory struct {
	Weekday             string
	TurnNum             int
	DiscCategoryRequest string
}
