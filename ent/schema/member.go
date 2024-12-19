package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("id"),
		field.Text("firstName").
			NotEmpty().StorageKey("firstName"),
		field.Text("lastName").
			NotEmpty().StorageKey("lastName"),
		field.Time("createdAt").
			Default(time.Now).
			Immutable().StorageKey("createdAt"),
		field.Time("updatedAt").
			Default(time.Now).
			Immutable().StorageKey("updatedAt"),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return nil
}

// Annotations of the User.
func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member"},
	}
}
