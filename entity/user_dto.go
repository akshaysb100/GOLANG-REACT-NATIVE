package entity

type UserDTO struct {
	ID           string `json:"ID,string,string"`
	USERNAME     string `json:"USERNAME,string,string"`
	MOBILENUMBER string `json:"MOBILENUMBER,string"`
}
