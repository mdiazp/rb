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

	GetTotalsByDiscCategory() (map[string]int, int, error)
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
		db = db.Where("disk.capacity = ?", *(filter.Capacity))
	}
	if filter.Category != nil {
		db = db.Where("disk.category = ?", *(filter.Category))
	}
	db = makeEqBoolFilter(db, filter.Actived, "disk.actived")

	return db
}

func (h *handler) CreateDisk(o *models.Disk) error {
	return h.Create(o).Error
}

func (h *handler) RetrieveDiskByID(id int) (*models.Disk, error) {
	o := &models.Disk{}
	e := h.Where("disk.id = ?", id).First(o).Error
	return o, e
}

func (h *handler) RetrieveDiskByName(name string) (*models.Disk, error) {
	o := &models.Disk{}
	e := h.Where("disk.name = ?", name).First(o).Error
	return o, e
}

func (h *handler) RetrieveDiskBySerialNumber(serialNumber string) (*models.Disk, error) {
	o := &models.Disk{}
	e := h.Where("disk.serial_number = ?", serialNumber).First(o).Error
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

func (h *handler) GetTotalsByDiscCategory() (map[string]int, int, error) {
	vtrue := true
	discs, e := h.RetrieveDiskList(
		&DiskFilter{
			Actived: &vtrue,
		},
		nil, nil,
	)
	if e != nil {
		return nil, 0, e
	}
	T := 0
	totals := make(map[string]int)
	for _, d := range *discs {
		totals[string(d.Category)]++
		T++
	}
	return totals, T, nil
}
