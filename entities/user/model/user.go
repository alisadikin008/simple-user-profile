package model

/*
	created by Ali Sadikin
	this file is intended to return data collection of user from database

	## Function List
	OfEmail(param)
	GetData(param)
	GetOne(params)
	PostData(params)
	CountRecords(param)
	PutOne(params)
	DeleteOne(id)

*/

import (
	sha "crypto/sha256"
	"encoding/hex"
	"errors"

	config "simple-user-profile/config"

	"github.com/jinzhu/gorm"
)

// User Struct
type User struct {
	gorm.Model
	//UserID   int    `gorm:"type:integer(10);primary_key;auto_increment" form:"userid" json:"userId"`
	Username string `gorm:"type:varchar(100)" form:"username" json:"username"`
	Email    string `gorm:"type:varchar(100);unique_index" form:"email" json:"email"`
	Password string `gorm:"type:varchar(255)" form:"password" json:"password"`
	Address  string `gorm:"type:varchar(255)" form:"address" json:"address"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&User{})
}

// OfUsers -()
func (u User) OfUsers(queryParams map[string][]string, page int, limit int) (interface{}, error) {
	var userObject []User
	db, _ := config.ConnectDB()
	query := db
	if result := query.Find(&userObject); result.Error != nil {
		return nil, errors.New("Data not found")
	}
	
	return userObject, nil
}

// OfEmail -()
func (u User) OfEmail(email string) (User, error) {
	var userObject User
	db, _ := config.ConnectDB()
	query := db
	query.Where("email = ?", email).First(&userObject)
	if userObject.ID == 0 {
		return userObject, errors.New("Data not found")
	}

	return userObject, nil
}

// OfUsername -()
func (u User) OfUsername(username string) (User, error) {
	var userObject User
	db, _ := config.ConnectDB()
	query := db
	query.Where("username = ?", username).First(&userObject)
	if userObject.ID == 0 {
		return userObject, errors.New("Data not found")
	}

	return userObject, nil
}

// OfID -()
func (u User) OfID(id int) (interface{}, error) {
	var userObject User
	db, _ := config.ConnectDB()
	query := db
	aha := query.First(&userObject, id)
	if userObject.ID == 0 {
		return userObject, errors.New("Data not found")
	}

	return aha, nil
}

// PostData -()
func (u User) PostData(userObject User) interface{} {
	db, _ := config.ConnectDB()
	hash := sha.Sum256([]byte(userObject.Password))
	userObject.Password = hex.EncodeToString(hash[:])
	tx := db.Begin()
	if err := tx.Create(&userObject).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return userObject
}

// PutOne -()
func (u User) PutOne(id int, userObject User) interface{} {
	var User User
	db, _ := config.ConnectDB()
	if userObject.Password != "" {
		hash := sha.Sum256([]byte(userObject.Password))
		userObject.Password = hex.EncodeToString(hash[:])
	}

	tx := db.Begin()
	if err := tx.First(&User, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Model(&User).Update(userObject)
	tx.Commit()
	return userObject
}
