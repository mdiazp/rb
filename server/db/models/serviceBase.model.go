package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ServiceStatus ...
type ServiceStatus string

const (
	// ServiceStatusPending ...
	ServiceStatusPending = "Pending"
	// ServiceStatusExecuted ...
	ServiceStatusExecuted = "Executed"
	// ServiceStatusCanceled ...
	ServiceStatusCanceled = "Canceled"
)

// GetServicesStatus ...
func GetServicesStatus() []ServiceStatus {
	return []ServiceStatus{
		ServiceStatusPending,
		ServiceStatusExecuted,
		ServiceStatusCanceled,
	}
}

// ValidateServiceStatus ...
func ValidateServiceStatus(o ServiceStatus) bool {
	l := GetServicesStatus()
	for _, x := range l {
		if x == o {
			return true
		}
	}
	return false
}

// ServiceBase ...
type ServiceBase struct {
	ID       int
	ClientID int           `gorm:"column:client_id;not null"`
	Cost     int           `gorm:"column:cost; not null"`
	Status   ServiceStatus `gorm:"type:varchar(20); not null"`
	Date     time.Time     `gorm:"not null"`
	Report   string        `gorm:"type:varchar(500); not null"`

	tableName string `gorm:"-"`
}

// Valid implementation of Model interface
func (o *ServiceBase) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateInt(&es, o.Cost, "Cost", 0, (1 << 30))
	validateTrue(&es, "Status", ValidateServiceStatus(o.Status))
	validateString(&es, o.Report, "Report", 1, 500)
	return &es
}

// AddSQLConstrainsts implementation of Model interface
func (o *ServiceBase) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Table(o.tableName).AddForeignKey("client_id", "client(id)", "CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(client_id): %s", "PServiceBase", e.Error())
		return
	}
	return
}
