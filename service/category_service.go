package service

import (
	"context"
	"go-restful-api/model/web"
)

type Service interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context)
	FindById(ctx context.Context)
	FindAll(ctx context.Context)
}
