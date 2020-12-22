package room_user

type User struct {
	CardNo      string `json:"card_no" form:"card_no" query:"card_no" gorm:"primaryKey"`
	DisplayName string `json:"display_name" form:"display_name" query:"display_name"`
	FullName    string `json:"full_name" form:"full_name" query:"full_name"`
	Email       string `json:"email" form:"email" query:"email"`
}

type CurrentUser struct {
	CardNo string `json:"card_no"`
}

func (CurrentUser) TableName() string {
	return "current_users"
}

type Tabler interface {
	TableName() string
}
