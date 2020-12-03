package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	Age     int
	Address []Address
}

type Address struct {
	gorm.Model
	Alamat string
	UserID uint
}
