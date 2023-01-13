package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct{
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	AuthorID 	int		`gorm:"type:int;not null;" json:"author_id"`
	PubID 		int		`gorm:"type:int;not null;" json:"publisher_id"`
	Title  		string    `gorm:"size:255;not null;unique" json:"title"`
	Price   	int    `gorm:"type:int;not null;" json:"price"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *Book) Prepare(){
	a.ID = 0
	a.Title =  html.EscapeString(strings.TrimSpace(a.Title))
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
}

func (a *Book) Validate() error{
	if a.Title == ""{
		return errors.New("Required Title")
	}
	if a.Price == 0{
		return errors.New("Required Price")
	}
	if a.AuthorID == 0{
		return errors.New("Required Author ID")
	}
	if a.PubID == 0{
		return errors.New("Required Publisher ID")
	}
	return nil
}

func (a *Book) StoreBook(db *gorm.DB) (*Book, error){
	var err error
	err = db.Debug().Model(&Book{}).Create(&a).Error
	if err != nil {
		return &Book{}, err
	}
	return a, nil
}

func (a *Book) FindAllBook(db *gorm.DB) (*[]Book,error){
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Limit(100).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, nil
}

func (a *Book) FindBookByAuthor(db *gorm.DB, aid uint64) (*[]Book, error){
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Where("author_id = ?", aid).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, nil
}

func (a *Book) FindBookByPublisher(db *gorm.DB, aid uint64) (*[]Book, error){
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Where("pub_id = ?", aid).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, nil
}

func (a *Book) FindBookByID(db *gorm.DB, aid uint64) (*Book, error){
	var err error
	err = db.Debug().Model(&Book{}).Where("id = ?", aid).Take(&a).Error
	if err != nil {
		return &Book{}, err
	}
	return a, nil
}

func (a *Book) UpdateBook(db *gorm.DB, aid uint64) (*Book, error){
	var err error
	err = db.Debug().Model(&Book{}).Where("id = ?", aid).Updates(Book{Title: a.Title, Price: a.Price, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Book{}, err
	}
	return a, nil
}

func (a *Book) DeleteBook (db *gorm.DB, aid uint64) (int64, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?",aid).Take(&Book{}).Delete(&Book{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Book not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}


