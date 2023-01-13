package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Publisher struct{
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name		  string    `gorm:"size:255;not null;unique" json:"name"`
	Address   string    `gorm:"size:255;not null;" json:"address"`
	Phone		  string    `gorm:"size:255;not null;" json:"phone"`
	Url   		string    `gorm:"size:255;not null;" json:"url"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *Publisher) Prepare(){
	a.ID = 0
	a.Name =  html.EscapeString(strings.TrimSpace(a.Name))
	a.Address =  html.EscapeString(strings.TrimSpace(a.Address))
	a.Phone = html.EscapeString(strings.TrimSpace(a.Phone))
	a.Url = html.EscapeString(strings.TrimSpace(a.Url))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Publisher) Validate() error{
	if a.Name == ""{
		return errors.New("Required Name")
	}
	if a.Address == ""{
		return errors.New("Required Address")
	}
	if a.Phone == ""{
		return errors.New("Required Phone Number")
	}
	if a.Url == ""{
		return errors.New("Required URL")
	}
	return nil
}

func (a *Publisher) StorePublisher(db *gorm.DB) (*Publisher, error){
	var err error
	err = db.Debug().Model(&Publisher{}).Create(&a).Error
	if err != nil {
		return &Publisher{}, err
	}
	return a, nil
}

func (a *Publisher) FindAllPublisher(db *gorm.DB) (*[]Publisher,error){
	var err error
	publishers := []Publisher{}
	err = db.Debug().Model(&Publisher{}).Limit(100).Find(&publishers).Error
	if err != nil {
		return &[]Publisher{}, err
	}
	return &publishers, nil
}

func (a *Publisher) FindPublisherByID(db *gorm.DB, aid uint64) (*Publisher, error){
	var err error
	err = db.Debug().Model(&Publisher{}).Where("id = ?", aid).Take(&a).Error
	if err != nil {
		return &Publisher{}, err
	}
	return a, nil
}

func (a *Publisher) UpdatePublisher(db *gorm.DB, aid uint64) (*Publisher, error){
	var err error
	err = db.Debug().Model(&Publisher{}).Where("id = ?", aid).Updates(Publisher{Name: a.Name, Address: a.Address,Phone: a.Phone, Url: a.Url,UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Publisher{}, err
	}
	return a, nil
}

func (a *Publisher) DeletePublisher (db *gorm.DB, aid uint64) (int64, error) {
	db = db.Debug().Model(&Publisher{}).Where("id = ?",aid).Take(&Publisher{}).Delete(&Publisher{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Publisher not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}


