package dao

import "github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/internal/model"

func (d *Dao) GetArticle(id uint32, state uint8) (*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) CreateArticle(title string, desc string, content string, create_by string, state uint8) error {
	article := model.Article{
		Title:   title,
		Desc:    desc,
		Content: content,
		State:   state,
		Model:   &model.Model{CreateBy: create_by},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(id uint32, title string, desc string, content string, modified_by string, state uint8) error {
	article := model.Article{
		Model: &model.Model{
			ID: id,
		},
	}

	// update的时候要注意判空，判0，用map更新而不是用struct，防止gorm无法分辨是空值还是传了0
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modified_by,
	}
	if title != "" {
		values["title"] = title
	}
	if desc != "" {
		values["desc"] = desc
	}
	if content != "" {
		values["content"] = content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{
		Model: &model.Model{
			ID: id,
		},
	}
	return article.Delete(d.engine)
}
