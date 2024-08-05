package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MenuCategory holds the schema definition for the MenuCategory entity.
type MenuCategory struct {
	ent.Schema
}

// Fields of the MenuCategory.
func (MenuCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Positive(),
		field.String("name").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the MenuCategory.
func (MenuCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("menu_categories").
			Field("user_id").
			Unique().
			Required(),
	}
}
