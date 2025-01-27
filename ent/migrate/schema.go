// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MemberColumns holds the columns for the "member" table.
	MemberColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "firstName", Type: field.TypeString, Size: 2147483647},
		{Name: "lastName", Type: field.TypeString, Size: 2147483647},
		{Name: "createdAt", Type: field.TypeTime},
		{Name: "updatedAt", Type: field.TypeTime},
	}
	// MemberTable holds the schema information for the "member" table.
	MemberTable = &schema.Table{
		Name:       "member",
		Columns:    MemberColumns,
		PrimaryKey: []*schema.Column{MemberColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MemberTable,
	}
)

func init() {
	MemberTable.Annotation = &entsql.Annotation{
		Table: "member",
	}
}
