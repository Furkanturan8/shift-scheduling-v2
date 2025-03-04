package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserRole int

const (
	UserRoleNormal UserRole = 1
	UserRoleDoctor UserRole = 2
	UserRoleAdmin  UserRole = 10
)

type User struct {
	ID        uuid.UUID
	Name      string   `gorm:"column:name"`
	Surname   string   `gorm:"column:surname"`
	UserName  string   `gorm:"column:username"`
	Email     string   `gorm:"column:email;unique"`
	Phone     string   `gorm:"column:phone;unique"`
	Password  string   `gorm:"column:password"`
	Role      UserRole `gorm:"column:role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (User) EntityName() string {
	return "user"
}

func (u User) String() string {
	return u.Name + " " + u.Surname
}

func (r UserRole) String() string {
	switch r {
	case UserRoleNormal:
		return "normal"
	case UserRoleDoctor:
		return "doctor"
	case UserRoleAdmin:
		return "admin"
	default:
		return "unknown"
	}
}
