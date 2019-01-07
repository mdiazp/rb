package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Client ...
type Client struct {
	ID             int
	Identification string `gorm:"type:varchar(100);not null"`
	Name           string `gorm:"type:varchar(100);not null"`
	Address        string `gorm:"type:varchar(1024);not null"`
	Phones         string `gorm:"type:varchar(100);not null"`
	Description    string `gorm:"type:varchar(1024);not null"`
	Actived        bool   `gorm:"not null"`
}

// TableName ...
func (Client) TableName() string {
	return "client"
}

// Valid implementation of Model interface
func (o *Client) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateString(&es, o.Identification, "Identification", 0, 100)
	validateString(&es, o.Name, "Name", 1, 100)
	validateString(&es, o.Address, "Address", 0, 1024)
	validateString(&es, o.Phones, "Phones", 0, 100)
	validateString(&es, o.Description, "Description", 0, 1024)
	return &es
}

// SQL Constrainsts
const (
	SQLConstrainstClientIdentificationUI = "uix_client_identification"
)

// AddSQLConstrainsts implementation of Model interface
func (o *Client) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddUniqueIndex(SQLConstrainstClientIdentificationUI,
		"identification").Error
	if e != nil {
		e = fmt.Errorf("%s - AddUniqueIndex(identification): %s", o.TableName(), e.Error())
		return
	}
	return
}
