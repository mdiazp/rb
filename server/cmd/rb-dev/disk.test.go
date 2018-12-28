package main

import (
	"fmt"

	"github.com/mdiazp/rb/server/db/models"
)

func mnInvalidDisk() models.Disk {
	return models.Disk{
		Name:         "",
		SerialNumber: "",
		Capacity:     0,
		Category:     "",
		Deleted:      false,
	}
}

func mxInvalidDisk() models.Disk {
	return models.Disk{
		Name:         randString(100 + 1),
		SerialNumber: randString(255 + 1),
		Capacity:     (1 << 30) + 1,
		Category:     "",
		Deleted:      true,
	}
}

func mnDisk() models.Disk {
	return models.Disk{
		Name:         randString(1),
		SerialNumber: randString(1),
		Capacity:     1,
		Category:     models.DiskCategorySmall,
		Deleted:      false,
	}
}

func mxDisk() models.Disk {
	return models.Disk{
		Name:         randString(100),
		SerialNumber: randString(255),
		Capacity:     (1 << 30),
		Category:     models.DiskCategoryBig,
		Deleted:      false,
	}
}

func printDiskValidationErrors(es *[]models.ValidationError) {
	fmt.Printf("Errors Size: %d\n", len(*es))
	for _, e := range *es {
		fmt.Printf("Error at %s: %s\n", e.PropertyName, e.Error)
	}
	fmt.Println()
}

func testDiskValidation() {
	disk := mnInvalidDisk()
	es := disk.Valid()
	printDiskValidationErrors(es)

	disk = mxInvalidDisk()
	es = disk.Valid()
	printDiskValidationErrors(es)

	disk = mnDisk()
	es = disk.Valid()
	printDiskValidationErrors(es)

	disk = mxDisk()
	es = disk.Valid()
	printDiskValidationErrors(es)
}

func testCreateDisk() {
	n := 100
	for i := 0; i < n; i++ {
		disk := mxDisk()
		e := db.CreateDisk(&disk)
		if e != nil {
			panic(e)
		}
	}
}

func testDeleteDisk() {
	l, e := db.RetrieveDiskList(nil, nil, nil)
	pe(e)

	for _, disk := range *l {
		db.DeleteDisk(disk.ID)
	}
}

func testUpdateDisk() {
	l, e := db.RetrieveDiskList(nil, nil, nil)
	pe(e)

	for _, disk := range *l {
		disk.Name = randString(5)
		disk.SerialNumber = randString(10)
		disk.Category = models.DiskCategorySmall
		disk.Capacity = 1
		disk.Deleted = true

		e = db.UpdateDisk(&disk)
		pe(e)
	}
}
