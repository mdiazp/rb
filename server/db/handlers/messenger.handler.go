package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// MessengerHandler ...
type MessengerHandler interface {
	CreateMessenger(o *models.Messenger) error
	RetrieveMessengerByID(id int) (*models.Messenger, error)
	RetrieveMessengerByName(name string) (*models.Messenger, error)
	UpdateMessenger(o *models.Messenger) error
	DeleteMessenger(id int) error

	RetrieveMessengerList(filter *MessengerFilter, orderBy *OrderBy,
		pag *Paginator) (*[]models.Messenger, error)
	CountMessengers(filter *MessengerFilter) (count int, e error)
}

// MessengerFilter ...
type MessengerFilter struct {
	NameSubstr *string
}

/////////////////////////////////////////////////////////////////////////////////////

func makeMessengerFilter(db *gorm.DB, filter *MessengerFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	db = makeSubstrFilter(db, filter.NameSubstr, "messenger.name")
	return db
}

func (h *handler) CreateMessenger(o *models.Messenger) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveMessengerByID(id int) (*models.Messenger, error) {
	o := &models.Messenger{}
	e := h.Where("messenger.id = ?", id).First(o).Error
	return o, e
}

func (h *handler) RetrieveMessengerByName(name string) (*models.Messenger, error) {
	o := &models.Messenger{}
	e := h.Where("messenger.name = ?", name).First(o).Error
	return o, e
}

func (h *handler) UpdateMessenger(o *models.Messenger) error {
	return h.Save(o).Error
}

func (h *handler) DeleteMessenger(id int) error {
	return h.Delete(models.Messenger{ID: id}).Error
}

func (h *handler) RetrieveMessengerList(filter *MessengerFilter, orderBy *OrderBy,
	pag *Paginator) (*[]models.Messenger, error) {
	db := makeMessengerFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.Messenger{}).TableName())

	l := make([]models.Messenger, 0)
	e := db.Find(&l).Error

	return &l, e
}

func (h *handler) CountMessengers(filter *MessengerFilter) (count int, e error) {
	e = makeMessengerFilter(h.DB.Model(&models.Messenger{}), filter).Count(&count).Error
	return
}
