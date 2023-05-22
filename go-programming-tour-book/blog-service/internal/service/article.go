package service

import "github.com/Temptation1/go_blog/go-programming-tour-book/blog-service/internal/model"

type ArticleGetRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type ArticleCreateRequest struct {
	Title    string `form:"title" binding:"required,min=1,max=100"`
	Desc     string `form:"desc" binding:"required,min=2,max=255"`
	Content  string `form:"content" binding:"required,min=1"`
	CreateBy string `form:"create_by" binding:"required,min=2,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleUpdateRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Title      string `form:"title" binding:"max=100"`
	Desc       string `form:"desc" binding:"max=255"`
	Content    string `form:"content"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleDeleteRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) GetArticle(param *ArticleGetRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID, param.State)
}

func (svc *Service) CreateArticle(param *ArticleCreateRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CreateBy, param.State)
}

func (svc *Service) UpdateArticle(param *ArticleUpdateRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.ModifiedBy, param.State)
}

func (svc *Service) DeleteArticle(param *ArticleDeleteRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
