package entity

type User struct {
	ID           int64  `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	Mobilenumber string `db:"mobileNumber" json:"mobileNumber"`
}
