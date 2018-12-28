package handlers

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Paginator ...
type Paginator struct {
	Offset int
	Limit  int
}

// OrderBy ...
type OrderBy struct {
	By   string
	DESC bool
}

func orderByAndPaginator(db *gorm.DB, orderBy *OrderBy, pag *Paginator, tableName string) *gorm.DB {
	if orderBy != nil {
		kk := "asc"
		if orderBy.DESC {
			kk = "desc"
		}
		db = db.Order(tableName + "." + orderBy.By + " " + kk)
	}
	if pag != nil {
		db = db.Limit(pag.Limit).Offset(pag.Offset)
	}
	return db
}

func makeSubstrFilter(db *gorm.DB, value *string, column string) *gorm.DB {
	if value != nil {
		s := fmt.Sprintf("%s ILIKE ?", column)
		db = db.Where(s, "%"+*value+"%")
	}
	return db
}
func makePrefixFilter(db *gorm.DB, value *string, column string) *gorm.DB {
	if value != nil {
		s := fmt.Sprintf("%s ILIKE ?", column)
		db = db.Where(s, *value+"%")
	}
	return db
}
func makeEqIntFilter(db *gorm.DB, value *int, column string) *gorm.DB {
	if value != nil {
		s := fmt.Sprintf("%s = ?", column)
		db = db.Where(s, *value)
	}
	return db
}
func makeEqBoolFilter(db *gorm.DB, value *bool, column string) *gorm.DB {
	if value != nil {
		s := fmt.Sprintf("%s = ?", column)
		db = db.Where(s, *value)
	}
	return db
}
func makeEqStringFilter(db *gorm.DB, value *bool, column string) *gorm.DB {
	if value != nil {
		s := fmt.Sprintf("%s = ?", column)
		db = db.Where(s, *value)
	}
	return db
}
