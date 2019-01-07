package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	ID          int
	Provider    string `gorm:"type:varchar(100);index;not null"`
	Username    string `gorm:"type:varchar(50);not null"`
	Name        string `gorm:"type:varchar(100);not null"`
	Rol         string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(1024);not null"`
	Actived     bool   `gorm:"not null"`
}

// TableName ...
func (User) TableName() string {
	return "system_user"
}

// Valid implementation of Model interface
func (o *User) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	validateString(&es, o.Provider, "Provider", 1, 100)
	validateString(&es, o.Username, "Username", 1, 50)
	validateString(&es, o.Name, "Name", 1, 100)
	validateString(&es, o.Rol, "Rol", 1, 100)
	validateString(&es, o.Description, "Description", 0, 1024)
	return &es
}

// SQL Constrainsts
const (
	SQLConstrainstUserUsernameUI = "uix_user_username"
)

// AddSQLConstrainsts implementation of Model interface
func (o *User) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Model(o).AddUniqueIndex(SQLConstrainstUserUsernameUI,
		"username").Error
	if e != nil {
		e = fmt.Errorf("%s - AddUniqueIndex(username): %s", o.TableName(), e.Error())
		return
	}
	return
}
