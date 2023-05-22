package model

import (
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

//建立struct和表名的映射，因为crud的接口都是直接操作结构体的，得让它自己知道对哪张表操作
func (t Tag) TableName() string {
	return "blog_tag"
}

func (t Tag) Create(db *gorm.DB) error {
	var isExist int
	db = db.Where("name=? AND is_del=?", t.Name, 0)
	err := db.Model(t).Select(1).Count(&isExist).Error
	if err != nil {
		return err
	}
	if isExist == 0 {
		return db.Create(&t).Error
	}
	return errcode.ErrorCreateTagFail
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//db = db.Model(&Tag{}).Where("id = ? AND is_del = ?", t.ID, 0)
	err := db.Model(t).Where("id=? AND is_del=?", t.ID, 0).Updates(values).Error
	if err != nil {
		return err
	}
	return nil //根据id值索引得到，然后再改内部的其它内容
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.Model.ID, 0).Delete(&t).Error
}

//count时候其实就没有输入param, name查询之类的应该没起作用,返回的就是所有数据的总行数
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("State = ?", t.State)
	err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
