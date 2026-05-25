package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Nama     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}

func (User) TableName() string {
	return "users"
}
