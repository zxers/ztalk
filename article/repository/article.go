package repository

import (
	"context"
	"ztalk/article/domain"
)

type ArticleRepositroy interface {
	List(ctx context.Context, AuthorId int, offset int, limit int) []domain.Article
	ListPub(ctx context.Context, offset int, limit int) []domain.Article
	Update()
	Create()
}

var _ ArticleRepositroy = &articleRepositroy{}

type articleRepositroy struct {
}

// Create implements ArticleRepositroy.
func (a *articleRepositroy) Create() {
	panic("unimplemented")
}

// Update implements ArticleRepositroy.
func (a *articleRepositroy) Update() {
	panic("unimplemented")
}

// List implements ArticleRepositroy.
func (a *articleRepositroy) List(ctx context.Context, AuthorId int, offset int, limit int) []domain.Article {
	panic("unimplemented")
}

// ListPub implements ArticleRepositroy.
func (a *articleRepositroy) ListPub(ctx context.Context, offset int, limit int) []domain.Article {
	panic("unimplemented")
}
