package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("/Users/nurhakimarif/workspaces/g2go/gorm/test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}

	var user User

	db.First(&user)
	fmt.Println(user.Name)

	result := map[string]interface{}{}
	db.Model(&User{}).First(result)
	fmt.Println(result["name"])

	result = map[string]interface{}{}
	db.Table("users").Take(&result)
	fmt.Println(result)

	var users User
	all,_ := db.Model(&users).Rows()

	for all.Next(){
		db.ScanRows(all, &users)
		fmt.Println(users.ID)
	}

	var u User
	where,_ := db.Table("users").Where("age=?", 17).Rows()

	for where.Next(){
		db.ScanRows(where, &u)
		fmt.Println(u.ID)
	}

}
