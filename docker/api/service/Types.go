package service

type User struct {
	Id          int    `json:"id"`
	DisplayName string `json:"display_name"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	CardNo      string `json:"card_no"`
}
