package service

import (
	"context"
	"ztalk/article/domain"
	"ztalk/article/repository"
)

type ArticleService interface {
	List(ctx context.Context, AuthorId int, offset int, limit int) []domain.Article
	ListPub(ctx context.Context, offset int, limit int) []domain.Article
	Save(ctx context.Context, article domain.Article) error
}

var _ ArticleService = &articleService{}

type articleService struct {
	ArticleRepositroy repository.ArticleRepositroy
}



func NewArticleService(articleRepositroy repository.ArticleRepositroy) *articleService {

	return &articleService{
		ArticleRepositroy: articleRepositroy,
	}
}

// List implements ArticleService.
func (a *articleService) List(ctx context.Context, AuthorId int, offset int, limit int) []domain.Article {
	panic("unimplemented")
}

// ListPub implements ArticleService.
func (a *articleService) ListPub(ctx context.Context, offset int, limit int) []domain.Article {
	panic("unimplemented")
}

// Save implements ArticleService.
func (a *articleService) Save(ctx context.Context, article domain.Article) error {
	panic("unimplemented")
}