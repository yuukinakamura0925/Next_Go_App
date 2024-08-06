package usecases

import (
	"Next_Go_App/ent"
	"context"

	"Next_Go_App/internal/repositories"
)

type IMenuCategoryUsecase interface {
	CreateMenuCategory(ctx context.Context, userID int, name string) (*ent.MenuCategory, error)
	GetMenuCategoryList(ctx context.Context, userID int) ([]*ent.MenuCategory, error)
	GetMenuCategoryDetail(ctx context.Context, id int) (*ent.MenuCategory, error)
	UpdateMenuCategory(ctx context.Context, id int, name string) (*ent.MenuCategory, error)
	DeleteMenuCategory(ctx context.Context, id int) error
}

type menuCategoryUsecase struct {
	menuCategoryRepo repositories.IMenuCategoryRepository
}

func NewMenuCategoryUsecase(menuCategoryRepo repositories.IMenuCategoryRepository) IMenuCategoryUsecase {
	return &menuCategoryUsecase{menuCategoryRepo: menuCategoryRepo}
}

func (u *menuCategoryUsecase) CreateMenuCategory(ctx context.Context, userID int, name string) (*ent.MenuCategory, error) {
	return u.menuCategoryRepo.CreateMenuCategory(ctx, userID, name)
}

func (u *menuCategoryUsecase) GetMenuCategoryList(ctx context.Context, userID int) ([]*ent.MenuCategory, error) {

	return u.menuCategoryRepo.GetMenuCategoryList(ctx, userID)
}

func (u *menuCategoryUsecase) GetMenuCategoryDetail(ctx context.Context, id int) (*ent.MenuCategory, error) {
	return u.menuCategoryRepo.GetMenuCategoryDetail(ctx, id)
}

func (u *menuCategoryUsecase) UpdateMenuCategory(ctx context.Context, id int, name string) (*ent.MenuCategory, error) {
	return u.menuCategoryRepo.UpdateMenuCategory(ctx, id, name)
}

func (u *menuCategoryUsecase) DeleteMenuCategory(ctx context.Context, id int) error {
	return u.menuCategoryRepo.DeleteMenuCategory(ctx, id)
}
