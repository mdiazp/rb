package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// InitDB ...
func InitDB(db *gorm.DB) (e error) {
	ms := []Model{
		&DiskReservation{
			ServiceBase: ServiceBase{
				tableName: "disk_reservation",
			},
		},
		&PMessageForDR{
			PMessageBase: PMessageBase{
				tableName: "p_message_for_dr",
			},
		},
		&PDiskReservation{
			PServiceBase: PServiceBase{
				tableName: "p_disk_reservation",
			},
		},
		&PDiskCopy{
			PServiceBase: PServiceBase{
				tableName: "p_disk_copy",
			},
		},
		&Messenger{},
		&Disk{},
		&Client{},
		&User{},
	}

	model := make([]interface{}, 0)
	for _, x := range ms {
		model = append(model, x)
	}

	// db.DropTableIfExists(model...)
	// return
	db.SingularTable(true)
	e = db.AutoMigrate(model...).Error
	if e != nil {
		return fmt.Errorf("db.Automigrate: %s", e.Error())
	}

	for _, x := range ms {
		if e = x.AddSQLConstrainsts(db); e != nil {
			return e
		}
	}

	return
}
