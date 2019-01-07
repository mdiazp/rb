package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Messenger ...
type Messenger struct {
	ID   int
	Name string `gorm:"type:varchar(100); not null"`
}

// TableName ...
func (Messenger) TableName() string {
	return "messenger"
}

// Valid ...
func (o *Messenger) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateString(&es, o.Name, "Name", 1, 100)
	return &es
}

const (
	// SQLConstrainstMessengerNameUI ...
	SQLConstrainstMessengerNameUI = "uix_messenger_name"
)

// AddSQLConstrainsts implementation of Model interface
func (o *Messenger) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddUniqueIndex(SQLConstrainstMessengerNameUI, "name").Error
	if e != nil {
		e = fmt.Errorf("%s - AddUniqueIndex(name): %s", o.TableName(), e.Error())
		return
	}
	return
}
