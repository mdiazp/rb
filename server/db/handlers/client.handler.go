package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// ClientHandler ...
type ClientHandler interface {
	CreateClient(o *models.Client) error
	RetrieveClientByID(id int) (*models.Client, error)
	RetrieveClientByIdentification(identification string) (*models.Client, error)
	UpdateClient(o *models.Client) error
	DeleteClient(id int) error

	RetrieveClientList(filter *ClientFilter,
		orderBy *OrderBy, pag *Paginator) (*[]models.Client, error)
	CountClients(filter *ClientFilter) (count int, e error)
}

// ClientFilter ...
type ClientFilter struct {
	IdentificationPrefix *string
	NameSubstr           *string
	AddressSubstr        *string
	PhonesSubstr         *string
	DescriptionSubstr    *string
	Actived              *bool
}

/////////////////////////////////////////////////////////////////////////////////////

func makeClientFilter(db *gorm.DB, filter *ClientFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	db = makePrefixFilter(db, filter.IdentificationPrefix, "identification")
	db = makeSubstrFilter(db, filter.NameSubstr, "name")
	db = makeSubstrFilter(db, filter.AddressSubstr, "address")
	db = makeSubstrFilter(db, filter.PhonesSubstr, "phones")
	db = makeSubstrFilter(db, filter.DescriptionSubstr, "description")
	db = makeEqBoolFilter(db, filter.Actived, "actived")

	return db
}

func (h *handler) CreateClient(o *models.Client) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveClientByID(id int) (*models.Client, error) {
	o := &models.Client{}
	e := h.Where("id = ?", id).First(o).Error
	return o, e
}

func (h *handler) RetrieveClientByIdentification(identification string) (*models.Client, error) {
	o := &models.Client{}
	e := h.Where("identification = ?", identification).First(o).Error
	return o, e
}

func (h *handler) UpdateClient(o *models.Client) error {
	return h.Save(o).Error
}

func (h *handler) DeleteClient(id int) error {
	return h.Delete(models.Client{ID: id}).Error
}

func (h *handler) RetrieveClientList(filter *ClientFilter,
	orderBy *OrderBy, pag *Paginator) (*[]models.Client, error) {
	db := makeClientFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.Client{}).TableName())

	l := make([]models.Client, 0)
	e := db.Find(&l).Error

	return &l, e
}

func (h *handler) CountClients(filter *ClientFilter) (count int, e error) {
	e = makeClientFilter(h.DB.Model(&models.Client{}), filter).Count(&count).Error
	return
}
