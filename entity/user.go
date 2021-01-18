package entity

type User struct {
	Username     string `db:"Username" json:"username"`
	Mobilenumber string `db:"mobileNumber" json:"mobileNumber"`
}
