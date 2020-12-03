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
	db.Table("users").Where("id=?",3).First(&user)

	user.Name = "Caca Handika"
	user.Age = 20
	db.Save(&user)

	db.Model(&User{}).Where("id=?", 2).Update("name", "Mansyur S.")

	db.Model(&User{}).Where("id=?", 1).Updates(User{ Name: "Super Senior", Age:  100, })

}

