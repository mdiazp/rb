package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// InitDB ...
func InitDB(db *gorm.DB) (e error) {
	ms := []Model{
		&Disk{},
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
