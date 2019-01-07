package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// PServiceBase ...
type PServiceBase struct {
	ID       int
	ClientID int `gorm:"column:client_id;not null"`
	ClientName string `gorm:"-"`
	IntervalTime
	Cost      int    `gorm:"column:cost; not null"`
	tableName string `gorm:"-"`
}

// Valid implementation of Model interface
func (o *PServiceBase) Valid() *[]ValidationError {
	es := *(o.IntervalTime.Valid())
	validateInt(&es, o.Cost, "Cost", 0, (1 << 30))
	return &es
}

// AddSQLConstrainsts implementation of Model interface
func (o *PServiceBase) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Table(o.tableName).AddForeignKey("client_id", "client(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(client_id): %s", "PServiceBase", e.Error())
		return
	}
	return
}
