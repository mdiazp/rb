package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DiskReservation ...
type DiskReservation struct {
	ServiceBase
	Turn
	DiskID int `gorm:"not null"`
}

// TableName ...
func (DiskReservation) TableName() string {
	return "disk_reservation"
}

// Valid implementation of Model interface
func (o *DiskReservation) Valid() *[]ValidationError {
	es := *(o.ServiceBase.Valid())
	es = append(es, *(o.Turn.Valid())...)
	return &es
}

// AddSQLConstrainsts implementation of Model interface
func (o *DiskReservation) AddSQLConstrainsts(db *gorm.DB) (e error) {
	o.ServiceBase.AddSQLConstrainsts(db)
	e = db.Table(o.tableName).AddForeignKey("disk_id", "disk(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(disk_id): %s", "DiskReservation", e.Error())
		return
	}
	return
}
