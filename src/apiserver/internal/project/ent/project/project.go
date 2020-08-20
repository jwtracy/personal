// Code generated by entc, DO NOT EDIT.

package project

import (
	"time"
)

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHead holds the string denoting the head field in the database.
	FieldHead = "head"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldStarted holds the string denoting the started field in the database.
	FieldStarted = "started"
	// FieldCompleted holds the string denoting the completed field in the database.
	FieldCompleted = "completed"

	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"

	// Table holds the table name of the project in the database.
	Table = "projects"
	// TagsTable is the table the holds the tags relation/edge. The primary key declared below.
	TagsTable = "topic_projects"
	// TagsInverseTable is the table name for the Topic entity.
	// It exists in this package in order to avoid circular dependency with the "topic" package.
	TagsInverseTable = "topics"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldHead,
	FieldBody,
	FieldCreateTime,
	FieldUpdateTime,
	FieldStarted,
	FieldCompleted,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"topic_id", "project_id"}
)

var (
	// HeadValidator is a validator for the "head" field. It is called by the builders before save.
	HeadValidator func(string) error
	// BodyValidator is a validator for the "body" field. It is called by the builders before save.
	BodyValidator func(string) error
	// DefaultCreateTime holds the default value on creation for the create_time field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the update_time field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	UpdateDefaultUpdateTime func() time.Time
)