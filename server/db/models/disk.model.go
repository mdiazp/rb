package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// DiskCategory ...
type DiskCategory string

const (
	// DiskCategoryBig ...
	DiskCategoryBig = "BIG"
	// DiskCategorySmall ...
	DiskCategorySmall = "SMALL"
)

// GetDiskCategories ...
func GetDiskCategories() []DiskCategory {
	return []DiskCategory{
		DiskCategoryBig,
		DiskCategorySmall,
	}
}

// ValidateCategoryDisk ...
func ValidateCategoryDisk(c DiskCategory) bool {
	cs := GetDiskCategories()
	for _, cv := range cs {
		if cv == c {
			return true
		}
	}
	return false
}

// GB GigaByte unit
type GB int

// Disk ...
type Disk struct {
	ID           int
	Name         string       `gorm:"type:varchar(100);not null"`
	SerialNumber string       `gorm:"type:varchar(255);not null"`
	Capacity     GB           `gorm:"not null"`
	Category     DiskCategory `gorm:"type:varchar(50);not null"`
	Actived      bool         `gorm:"not null"`
}

// TableName ...
func (Disk) TableName() string {
	return "disk"
}

// Valid implementation of Model interface
func (o *Disk) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateString(&es, o.Name, "Name", 1, 100)
	validateString(&es, o.SerialNumber, "SerialNumber", 1, 255)
	validateInt(&es, (int)(o.Capacity), "Capacity", 1, (1 << 30))
	validateValue(&es, "Category", ValidateCategoryDisk(o.Category))
	return &es
}

// SQL Constrainsts
const (
	SQLConstrainstDiskNameUI         = "uix_disk_name"
	SQLConstrainstDiskSerialNumberUI = "uix_disk_serial_number"
)

// AddSQLConstrainsts implementation of Model interface
func (o *Disk) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddUniqueIndex(SQLConstrainstDiskNameUI, "name").Error
	if e != nil {
		e = fmt.Errorf("%s - AddUniqueIndex(name): %s", o.TableName(), e.Error())
		return
	}

	e = db.Model(o).AddUniqueIndex(SQLConstrainstDiskSerialNumberUI, "serial_number").Error
	if e != nil {
		e = fmt.Errorf("%s - AddUniqueIndex(serial_number): %s", o.TableName(), e.Error())
		return
	}
	return
}
