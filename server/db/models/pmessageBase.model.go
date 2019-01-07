package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// MessageType ...
type MessageType string

const (
	// MessageTypeEntrega ...
	MessageTypeEntrega = "Entrega"
	// MessageTypeRecogida ...
	MessageTypeRecogida = "Recogida"
)

// GetMessageTypes ...
func GetMessageTypes() []MessageType {
	return []MessageType{
		MessageTypeEntrega,
		MessageTypeRecogida,
	}
}

// ValidateMessageType ...
func ValidateMessageType(o MessageType) bool {
	l := GetMessageTypes()
	for _, x := range l {
		if x == o {
			return true
		}
	}
	return false
}

// PMessageBase ...
type PMessageBase struct {
	ID          int
	MessengerID int `gorm:"not null"`
	Turn
	SortPosition int         `gorm:"not null"`
	Type         MessageType `gorm:"type:varchar(50); not null"`
	Notes        string      `gorm:"type:varchar(500); not null"`
	tableName    string      `gorm:"-"`
}

// Valid implementation of Model interface
func (o *PMessageBase) Valid() *[]ValidationError {
	es := make([]ValidationError, 0)
	es = append(es, *(o.Turn.Valid())...)
	validateInt(&es, o.SortPosition, "SortPosition", 0, (1 << 30))
	validateTrue(&es, "Type", ValidateMessageType(o.Type))
	validateString(&es, o.Notes, "Notes", 0, 500)
	return &es
}

// AddSQLConstrainsts implementation of Model interface
func (o *PMessageBase) AddSQLConstrainsts(db *gorm.DB) (e error) {
	e = db.Table(o.tableName).AddForeignKey("messenger_id", "messenger(id)",
		"CASCADE", "CASCADE").Error
	if e != nil {
		e = fmt.Errorf("%s - AddForeignKey(messenger_id): %s", "PMessageBase",
			e.Error())
		return
	}
	return
}
