package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// DiskHandler ...
type DiskHandler interface {
	CreateDisk(o *models.Disk) error
	RetrieveDiskByID(id int) (*models.Disk, error)
	RetrieveDiskByName(name string) (*models.Disk, error)
	RetrieveDiskBySerialNumber(serialNumber string) (*models.Disk, error)
	UpdateDisk(o *models.Disk) error
	DeleteDisk(id int) error

	RetrieveDiskList(filter *DiskFilter,
		orderBy *OrderBy, pag *Paginator) (*[]models.Disk, error)
	CountDisks(filter *DiskFilter) (count int, e error)
}

// DiskFilter ...
type DiskFilter struct {
	Capacity *models.GB
	Category *models.DiskCategory
	Actived  *bool
}

/////////////////////////////////////////////////////////////////////////////////////

func makeDiskFilter(db *gorm.DB, filter *DiskFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	if filter.Capacity != nil {
		db = db.Where("Capacity = ?", *(filter.Capacity))
	}
	if filter.Category != nil {
		db = db.Where("Category = ?", *(filter.Category))
	}
	db = makeEqBoolFilter(db, filter.Actived, "actived")

	return db
}

func (h *handler) CreateDisk(o *models.Disk) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveDiskByID(id int) (*models.Disk, error) {
	o := &models.Disk{}
	e := h.Where("id = ?", id).First(o).Error
	return o, e
}

func (h *handler) RetrieveDiskByName(name string) (*models.Disk, error) {
	o := &models.Disk{}
	e := h.Where("name = ?", name).First(o).Error
	return o, e
}

func (h *handler) RetrieveDiskBySerialNumber(serialNumber string) (*models.Disk, error) {
	o := &models.Disk{}
	e := h.Where("serial_number = ?", serialNumber).First(o).Error
	return o, e
}

func (h *handler) UpdateDisk(o *models.Disk) error {
	return h.Save(o).Error
}

func (h *handler) DeleteDisk(id int) error {
	return h.Delete(models.Disk{ID: id}).Error
}

func (h *handler) RetrieveDiskList(filter *DiskFilter,
	orderBy *OrderBy, pag *Paginator) (*[]models.Disk, error) {
	db := makeDiskFilter(h.DB, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.Disk{}).TableName())

	l := make([]models.Disk, 0)
	e := db.Find(&l).Error

	return &l, e
}

func (h *handler) CountDisks(filter *DiskFilter) (count int, e error) {
	e = makeDiskFilter(h.DB.Model(&models.Disk{}), filter).Count(&count).Error
	return
}
