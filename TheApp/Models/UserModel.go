package Models

import (
	"gorm.io/gorm"
	"theapp/Service"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Password string `json:"password"`
	Role string `json:"role"`
}
func (b *User) TableName() string {
	return "user"
}

func (u *User)BeforeCreate(tx *gorm.DB)  (err error){
	u.Name = Service.Encrypt(u.Name)
	return
}

func (u *User)BeforeUpdate(tx *gorm.DB)  (err error){
	u.Name = Service.Encrypt(u.Name)
	return
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.Name = Service.Decrypt(u.Name)
	return 
}
