package model

import (
	"github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content" gorm:"type:longtext"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
	Tags          []Tag  `json:"tags" gorm:"many2many:blog_tag_article"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Get(db *gorm.DB) (*Article, error) {
	article := &Article{}
	db = db.Where("is_del=?", 0)

	if err := db.Preload("Tags").First(article, a.ID).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (a Article) Create(db *gorm.DB) error {
	//根据文章标题判断是否重复
	var isExist int
	db = db.Where("title=? AND is_del=?", a.Title, 0)
	err := db.Model(a).Select(1).Count(&isExist).Error
	if err != nil {
		return err
	}
	if isExist == 0 {
		//return db.Create(&a).Error
		return db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&a).Error; err != nil {
				return err
			}
			return nil
		})
	}
	return errcode.ErrorCreateArticle
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	db = db.Where("id=? and is_del=?", a.ID, 0)
	err := db.Model(a).Updates(values).Error
	if err != nil {
		return err
	}
	return nil
}

//删除有问题
func (a Article) Delete(db *gorm.DB) error {
	//global.Logger.Info(a.ID)
	return db.Select("Tags").Where("id=? AND is_del=?", a.ID, 0).Delete(&a).Error
	//return db.Model(a).Delete(&a).Error

}
