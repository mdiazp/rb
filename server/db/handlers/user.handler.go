package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// UserHandler ...
type UserHandler interface {
	CreateUser(o *models.User) error
	RetrieveUserByID(id int) (*models.User, error)
	RetrieveUserByUsername(username string) (*models.User, error)
	UpdateUser(o *models.User) error
	DeleteUser(id int) error

	RetrieveUserList(filter *UserFilter,
		orderBy *OrderBy, pag *Paginator) (*[]models.User, error)
	CountUsers(filter *UserFilter) (count int, e error)
}

// UserFilter ...
type UserFilter struct {
	UsernameSubstr *string
	NameSubstr     *string
	Provider       *string
	Rol            *string
	Actived        *bool
}

/////////////////////////////////////////////////////////////////////////////////////

func makeUserFilter(db *gorm.DB, filter *UserFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	db = makeSubstrFilter(db, filter.UsernameSubstr, "username")
	db = makeSubstrFilter(db, filter.NameSubstr, "name")
	db = makeEqStringFilter(db, filter.Provider, "provider")
	db = makeEqStringFilter(db, filter.Rol, "rol")
	db = makeEqBoolFilter(db, filter.Actived, "actived")

	return db
}

func (h *handler) CreateUser(o *models.User) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveUserByID(id int) (*models.User, error) {
	o := &models.User{}
	e := h.Where("id = ?", id).First(o).Error
	return o, e
}

func (h *handler) RetrieveUserByUsername(username string) (*models.User, error) {
	o := &models.User{}
	e := h.Where("username = ?", username).First(o).Error
	return o, e
}

func (h *handler) RetrieveUserBySerialNumber(serialNumber string) (*models.User, error) {
	o := &models.User{}
	e := h.Where("serial_number = ?", serialNumber).First(o).Error
	return o, e
}

func (h *handler) UpdateUser(o *models.User) error {
	return h.Save(o).Error
}

func (h *handler) DeleteUser(id int) error {
	return h.Delete(models.User{ID: id}).Error
}

func (h *handler) RetrieveUserList(filter *UserFilter,
	orderBy *OrderBy, pag *Paginator) (*[]models.User, error) {
	db := makeUserFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.User{}).TableName())

	l := make([]models.User, 0)
	e := db.Find(&l).Error

	return &l, e
}

func (h *handler) CountUsers(filter *UserFilter) (count int, e error) {
	e = makeUserFilter(h.DB.Model(&models.User{}), filter).Count(&count).Error
	return
}
