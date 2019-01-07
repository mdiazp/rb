package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/mdiazp/rb/server/db/models"
)

// PDiskCopyHandler ...
type PDiskCopyHandler interface {
	CreatePDiskCopy(o *models.PDiskCopy) error
	RetrievePDiskCopyByID(id int) (*models.PDiskCopy, error)
	UpdatePDiskCopy(o *models.PDiskCopy) error
	DeletePDiskCopy(id int) error

	RetrievePDiskCopyList(filter *PDiskCopyFilter,
		orderBy *OrderBy, pag *Paginator) (*[]models.PDiskCopy, error)
	CountPDiskCopies(filter *PDiskCopyFilter) (count int, e error)
}

// PDiskCopyFilter ...
type PDiskCopyFilter struct {
	PServiceBaseFilter *PServiceBaseFilter
}

/////////////////////////////////////////////////////////////////////////////////////

func makePDiskCopyFilter(db *gorm.DB, filter *PDiskCopyFilter) *gorm.DB {
	if filter == nil {
		return db
	}
	db = makePServiceBaseFilter(db, filter.PServiceBaseFilter, "p_disk_copy")

	return db
}

func (h *handler) CreatePDiskCopy(o *models.PDiskCopy) error {
	return h.Create(o).Error
}

func (h *handler) RetrievePDiskCopyByID(id int) (*models.PDiskCopy, error) {
	o := &models.PDiskCopy{}
	e := h.Where("p_disk_copy.id = ?", id).First(o).Error
	return o, e
}

func (h *handler) UpdatePDiskCopy(o *models.PDiskCopy) error {
	return h.Save(o).Error
}

func (h *handler) DeletePDiskCopy(id int) error {
	return h.Delete(models.PDiskCopy{
		PServiceBase: models.PServiceBase{
			ID: id,
		},
	}).Error
}

func (h *handler) RetrievePDiskCopyList(filter *PDiskCopyFilter,
	orderBy *OrderBy, pag *Paginator) (*[]models.PDiskCopy, error) {
	db := h.DB.Model(&models.PDiskCopy{})
	db = addJoinToClient(db, "p_disk_copy")
	db = makePDiskCopyFilter(db, filter)
	db = orderByAndPaginator(db, orderBy, pag, (models.PDiskCopy{}).TableName())

	l := make([]models.PDiskCopy, 0)

	db = db.
		Select(
			"p_disk_copy.id, " +
				"p_disk_copy.client_id, " +
				"client.name, " +
				"p_disk_copy.initial_time, " +
				"p_disk_copy.finish_time, " +
				"p_disk_copy.cost",
		)

	fmt.Println("	SQL Query = ", db.QueryExpr())

	rows, e := db.Rows()
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	for rows.Next() {
		o := models.PDiskCopy{}
		e = rows.Scan(
			&o.ID,
			&o.ClientID,
			&o.ClientName,
			&o.InitialTime,
			&o.FinishTime,
			&o.Cost,
		)
		if e != nil {
			return nil, e
		}
		l = append(l, o)
	}

	return &l, e
}

func (h *handler) CountPDiskCopies(filter *PDiskCopyFilter) (count int, e error) {
	db := addJoinToClient(
		h.DB.Model(&models.PDiskCopy{}),
		"p_disk_copy",
	)
	e = makePDiskCopyFilter(db, filter).Count(&count).Error
	return
}
