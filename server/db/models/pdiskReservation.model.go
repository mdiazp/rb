package models

import "github.com/jinzhu/gorm"

const (
	// DiskCategoryRequestNull ...
	DiskCategoryRequestNull DiskCategory = "NULL"
)

// PDiskReservation ...
type PDiskReservation struct {
	PServiceBase
	Turn
	DiskCategoryRequest DiskCategory `gorm:"type:varchar(50); not null"`
}

// TableName ...
func (PDiskReservation) TableName() string {
	return "p_disk_reservation"
}

// Valid implementation of Model interface
func (o *PDiskReservation) Valid() *[]ValidationError {
	es := *(o.PServiceBase.Valid())

	if o.DiskCategoryRequest != DiskCategoryRequestNull {
		validateTrue(&es, "DiskCategoryRequest",
			ValidateDiskCategory(o.DiskCategoryRequest))
	}

	/*
		validateTrue(&es, "InitialTime",
			o.InitialTime.Weekday().String() == (string)(o.TurnWeekDay))
		validateTrue(&es, "FinishTime",
			o.FinishTime.Weekday().String() == (string)(o.TurnWeekDay))
	*/

	es = append(es, *(o.Turn.Valid())...)

	return &es
}

// AddSQLConstrainsts implementation of Model interface
func (o *PDiskReservation) AddSQLConstrainsts(db *gorm.DB) (e error) {
	o.PServiceBase.AddSQLConstrainsts(db)
	return
}
