package models

import "github.com/jinzhu/gorm"

// PDiskCopy ...
type PDiskCopy struct {
	PServiceBase
}

// TableName ...
func (PDiskCopy) TableName() string {
	return "p_disk_copy"
}

// Valid implementation of Model interface
func (o *PDiskCopy) Valid() *[]ValidationError {
	es := *(o.PServiceBase.Valid())
	return &es
}

// AddSQLConstrainsts implementation of Model interface
func (o *PDiskCopy) AddSQLConstrainsts(db *gorm.DB) (e error) {
	o.PServiceBase.AddSQLConstrainsts(db)
	return
}
