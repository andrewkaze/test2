package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("/Users/nurhakimarif/workspaces/g2go/gorm/test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}

	var user User

	db.Table("users").Where("id=?",1).Find(&user)
	db.Delete(&user)

	db.Delete(&User{}, 4)
	db.Delete(&User{}, []int {2,3,4})
}
