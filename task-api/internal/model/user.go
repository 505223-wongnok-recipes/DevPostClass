package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	// hashed password
	Password string
}

func (u User) Exists() bool {
	return u.ID != 0 && u.Username != ""
}
