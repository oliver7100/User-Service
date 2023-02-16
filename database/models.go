package database

type User struct {
	Email    string `gorm:"type:varchar(255);uniqueIndex"`
	Password string ``
	Name     string
}
