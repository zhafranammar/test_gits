package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Author struct{
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Fullname  string    `gorm:"size:255;not null;unique" json:"fullname"`
	Address   string    `gorm:"size:255;not null;" json:"address"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *Author) Prepare(){
	a.ID = 0
	a.Fullname =  html.EscapeString(strings.TrimSpace(a.Fullname))
	a.Address =  html.EscapeString(strings.TrimSpace(a.Address))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Author) Validate() error{
	if a.Fullname == ""{
		return errors.New("Required Fullname")
	}
	if a.Address == ""{
		return errors.New("Required Address")
	}
	return nil
}

func (a *Author) StoreAuthor(db *gorm.DB) (*Author, error){
	var err error
	err = db.Debug().Model(&Author{}).Create(&a).Error
	if err != nil {
		return &Author{}, err
	}
	return a, nil
}

func (a *Author) FindAllAuthor(db *gorm.DB) (*[]Author,error){
	var err error
	authors := []Author{}
	err = db.Debug().Model(&Author{}).Limit(100).Find(&authors).Error
	if err != nil {
		return &[]Author{}, err
	}
	return &authors, nil
}

func (a *Author) FindAuthorByID(db *gorm.DB, aid uint64) (*Author, error){
	var err error
	err = db.Debug().Model(&Author{}).Where("id = ?", aid).Take(&a).Error
	if err != nil {
		return &Author{}, err
	}
	return a, nil
}

func (a *Author) UpdateAuthor(db *gorm.DB, aid uint64) (*Author, error){
	var err error
	err = db.Debug().Model(&Author{}).Where("id = ?", aid).Updates(Author{Fullname: a.Fullname, Address: a.Address, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Author{}, err
	}
	return a, nil
}

func (a *Author) DeleteAuthor (db *gorm.DB, aid uint64) (int64, error) {
	db = db.Debug().Model(&Author{}).Where("id = ?",aid).Take(&Author{}).Delete(&Author{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Authora not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}


