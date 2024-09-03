package models

type Role struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users"`
}

type RoleWithoutUsers struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
