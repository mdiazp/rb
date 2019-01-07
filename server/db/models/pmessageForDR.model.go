package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// PMessageForDR ...
type PMessageForDR struct {
	PMessageBase
	PDiskReservationID int `gorm:"not null"`
}

// TableName ...
func (PMessageForDR) TableName() string {
	return "p_message_for_dr"
}

// Valid implementation of Model interface
func (o *PMessageForDR) Valid() *[]ValidationError {
	es := *(o.PMessageBase.Valid())
	return &es
}

const (
	// SQLConstrainstClientPDiskReservationIDTypeUI ...
	SQLConstrainstClientPDiskReservationIDTypeUI = "uix_pdiskreservationid_type"
)

// AddSQLConstrainsts implementation of Model interface
func (o *PMessageForDR) AddSQLConstrainsts(db *gorm.DB) (e error) {

	o.PMessageBase.AddSQLConstrainsts(db)

	e = db.Model(o).AddForeignKey("p_disk_reservation_id", "p_disk_reservation(id)",
		"CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(p_disk_reservation): %s", "PMessageForDR",
			e.Error())
		return
	}

	e = db.Model(o).AddUniqueIndex(SQLConstrainstClientPDiskReservationIDTypeUI,
		"p_disk_reservation_id", "type").Error
	if e != nil {
		e = fmt.Errorf("%s - AddUniqueIndex(p_disk_reservation_id-type): %s", o.TableName(), e.Error())
		return
	}
	return
}
