package repositories

import (
	"context"

	"Next_Go_App/ent"
	"Next_Go_App/ent/menucategory"
)

type IMenuCategoryRepository interface {
	CreateMenuCategory(ctx context.Context, userID int, name string) (*ent.MenuCategory, error)
	GetMenuCategoryList(ctx context.Context, userID int) ([]*ent.MenuCategory, error)
	GetMenuCategoryDetail(ctx context.Context, id int) (*ent.MenuCategory, error)
	UpdateMenuCategory(ctx context.Context, id int, name string) (*ent.MenuCategory, error)
	DeleteMenuCategory(ctx context.Context, id int) error
}

type menuCategoryRepository struct {
	client *ent.Client
}

func NewMenuCategoryRepository(client *ent.Client) IMenuCategoryRepository {
	return &menuCategoryRepository{client: client}
}

func (r *menuCategoryRepository) CreateMenuCategory(ctx context.Context, userID int, name string) (*ent.MenuCategory, error) {
	return r.client.MenuCategory.
		Create().
		SetUserID(userID).
		SetName(name).
		Save(ctx)
}

func (r *menuCategoryRepository) GetMenuCategoryList(ctx context.Context, userID int) ([]*ent.MenuCategory, error) {
	return r.client.MenuCategory.Query().Where(menucategory.UserIDEQ(userID)).All(ctx)
}

func (r *menuCategoryRepository) GetMenuCategoryDetail(ctx context.Context, id int) (*ent.MenuCategory, error) {
	return r.client.MenuCategory.Query().Where(menucategory.IDEQ(id)).Only(ctx)
}

func (r *menuCategoryRepository) UpdateMenuCategory(ctx context.Context, id int, name string) (*ent.MenuCategory, error) {
	return r.client.MenuCategory.UpdateOneID(id).
		SetName(name).
		Save(ctx)
}

func (r *menuCategoryRepository) DeleteMenuCategory(ctx context.Context, id int) error {
	return r.client.MenuCategory.DeleteOneID(id).Exec(ctx)
}
