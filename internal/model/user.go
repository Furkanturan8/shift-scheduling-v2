package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Role int
type Status string

const (
	UserRoleNormal Role = 1
	UserRoleDoctor Role = 2
	UserRoleAdmin  Role = 10
)

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusBanned   Status = "banned"
)

type User struct {
	BaseModel
	Email     string    `json:"email" bun:",unique,notnull"`
	Phone     string    `json:"phone" bun:",unique,notnull"`
	Password  string    `json:"-" bun:"password_hash,notnull"`
	Name      string    `json:"name" bun:"name"`
	Surname   string    `json:"surname" bun:"surname"`
	Role      Role      `json:"role" bun:"type:user_role,notnull,default:'user'"`
	Status    Status    `json:"status" bun:"type:user_status,notnull,default:'active'"`
	LastLogin time.Time `json:"last_login" bun:",nullzero"`

	tableName struct{} `bun:"users"`
}

func (u User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u User) GetStatus() Status {
	return u.Status
}

func (u User) String() string {
	return u.Name + " " + u.Surname
}

func (r Role) String() string {
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
