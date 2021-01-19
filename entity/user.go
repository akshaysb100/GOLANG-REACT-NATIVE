package entity

type User struct {
	ID           string `db:"id" json:"id"`
	USERNAME     string `db:"username" json:"username"`
	MOBILENUMBER string `db:"mobileNumber" json:"mobileNumber"`
}
