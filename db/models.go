package db

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Base is common for all the models
type Base struct {
	ID        string `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// Create primary key uuid for all the models
func (base *Base) BeforeCreate(_ *gorm.DB) error {
	id := uuid.New()
	base.ID = id.String()
	return nil
}

type User struct {
	Base
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

// Hash the user password before storing it
func (user *User) BeforeCreate(_ *gorm.DB) error {
	// As BeforeCreate method is overridden from Base here we need to set uuid part again
	id := uuid.New()
	user.ID = id.String()

	bts, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bts)
	return nil
}

type Book struct {
	Base
	Title  string `gorm:"unique;not null" json:"title"`
	Price  int    `json:"price"`
	Author string `json:"author"`
}

// Order is the combination of Book and User ids, which indicates that a user has ordered a book
type Order struct {
	Base
	UserID string
	User   User `gorm:"foreignkey:UserID" json:"-"`

	BookID string
	Book   Book `gorm:"foreignkey:BookID"`
}
