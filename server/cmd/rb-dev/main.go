package main

import (
	"fmt"
)

func main() {
	// testCreateDisk()
	testUpdateDisk()
	// testDeleteDisk()
	/*
		for i := 0; i < 100; i++ {
			act := true
			if rand.Int()%2 == 0 {
				act = false
			}
			g := models.Disk{
				Name:        randString(10),
				Description: "Some description",
				Actived:     act,
			}

			e := db.CreateDisk(&g)
			if e != nil {
				panic(e)
			}
		}
	*/

	fmt.Println("Finish")

}
