package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// PServiceBaseFilter ...
type PServiceBaseFilter struct {
	ClientID            *int
	ActivedClient       *bool
	ActivedIntervalTime *models.IntervalTime
}

/////////////////////////////////////////////////////////////////////////////////////

func addJoinToClient(db *gorm.DB, tableName string) *gorm.DB {
	db = db.Joins("left join client on " + tableName + ".client_id = client.id")
	return db
}

func makePServiceBaseFilter(db *gorm.DB, filter *PServiceBaseFilter, tableName string) *gorm.DB {
	if filter == nil {
		return db
	}

	db = makeEqIntFilter(db, filter.ClientID, tableName+".client_id")

	if filter.ActivedClient != nil {
		db = makeEqBoolFilter(db, filter.ActivedClient, "client.actived")
	}

	if filter.ActivedIntervalTime != nil {
		db = db.Where("NOT("+
			tableName+".finish_time < ? OR ? < "+tableName+".initial_time"+
			")",
			filter.ActivedIntervalTime.InitialTime.Format("2006-01-02 15:04:05"),
			filter.ActivedIntervalTime.FinishTime.Format("2006-01-02 15:04:05"),
		)
	}

	return db
}
