package db

import (
	"fmt"
	"log"

	"github.com/niksis02/book-store/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBPostgres struct {
	db *gorm.DB
}

var _ DBService = &DBPostgres{}

// Postgres instance of DBService
func NewPostgres() (DBService, error) {
	db, err := connectPostgres()
	if err != nil {
		return nil, err
	}

	return &DBPostgres{
		db: db,
	}, nil
}

// Creates a new user instance in the database
func (d *DBPostgres) CreateUser(email, password string) error {
	user := User{
		Email:    email,
		Password: password,
	}
	tx := d.db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Retrieves user from database by email address
func (d *DBPostgres) GetUserByEmail(email string) (*User, error) {
	var user User
	tx := d.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

// Creates orders with the given userId and book ids
func (d *DBPostgres) CreateOrders(userId string, bookIds []string) error {
	// Here database transactions would be a good option, but I didn't dive so deep.
	for _, id := range bookIds {
		order := Order{
			UserID: userId,
			BookID: id,
		}
		tx := d.db.Create(&order)
		if tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}

func (d *DBPostgres) ListUserOrders(userId string) ([]Order, error) {
	var orders []Order

	tx := d.db.Where("user_id = ?", userId).Preload("Book").Find(&orders)
	return orders, tx.Error
}

func connectPostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getConnectionStringFromEnv()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Postgres database has been successfully connected")

	err = db.AutoMigrate(&User{}, &Book{}, &Order{})
	if err != nil {
		return nil, err
	}
	log.Println("Database models have been successfully created")

	return db, nil
}

func getConnectionStringFromEnv() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		env.Env.DB_HOST,
		env.Env.DB_USER,
		env.Env.DB_PASSWORD,
		env.Env.DB_NAME,
		env.Env.DB_PORT,
	)
}
