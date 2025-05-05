package entities

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey;type:uuid;"`
	Username string `json:"username" gorm:"type:varchar(150);"`
	PwdHash	 string `json:"pwd_hash" gorm:"type:varchar(255);"`
	ProfilePicture string `json:"profile_picture" gorm:"type:varchar(255);"`
	Active	 bool   `json:"active" gorm:"type:boolean"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func NewUser(username, pwdHash string) *User {
	return &User{
		ID:			 uuid.NewString(),
		Username: 	 username,
		PwdHash: 	 pwdHash,
		Active:		 true,
		CreatedAt: 	 time.Now().UTC(),
		UpdatedAt: 	 time.Now().UTC(),
		DeletedAt: 	 gorm.DeletedAt{},
	}
}

func (u *User) SetActive() {
	u.Active = true
}

func (u *User) SetInactive() {	
	u.Active = false
}	
