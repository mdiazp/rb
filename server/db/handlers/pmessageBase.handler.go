package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// PMessageBaseFilter ...
type PMessageBaseFilter struct {
	MessengerID *int
	TurnWeekDay *models.WeekDay
	TurnNum     *models.TurnNum
	Type        *models.MessageType
	NotesSubstr *string
}

/////////////////////////////////////////////////////////////////////////////////////

func makePMessageBaseFilter(db *gorm.DB, filter *PMessageBaseFilter, tableName string) *gorm.DB {
	if filter == nil {
		return db
	}

	db = makeEqIntFilter(db, filter.MessengerID, tableName+".messenger_id")

	if filter.TurnWeekDay != nil {
		db = db.Where(tableName+".turn_week_day = ?", *(filter.TurnWeekDay))
	}
	if filter.TurnNum != nil {
		db = db.Where(tableName+".turn_num = ?", *(filter.TurnNum))
	}
	if filter.Type != nil {
		db = db.Where(tableName+".type = ?", *(filter.Type))
	}

	db = makeSubstrFilter(db, filter.NotesSubstr, "notes")

	return db
}
