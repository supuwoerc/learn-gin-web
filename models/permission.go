package models

type Permission struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Resource string  `json:"resource"`
	Roles    []*Role `json:"roles,omitempty"`
}
