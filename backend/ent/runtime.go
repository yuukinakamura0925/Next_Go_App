// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"Next_Go_App/ent/book"
	"Next_Go_App/ent/schema"
	"Next_Go_App/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	bookFields := schema.Book{}.Fields()
	_ = bookFields
	// bookDescUserID is the schema descriptor for user_id field.
	bookDescUserID := bookFields[0].Descriptor()
	// book.UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	book.UserIDValidator = bookDescUserID.Validators[0].(func(int) error)
	// bookDescTitle is the schema descriptor for title field.
	bookDescTitle := bookFields[1].Descriptor()
	// book.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	book.TitleValidator = bookDescTitle.Validators[0].(func(string) error)
	// bookDescBody is the schema descriptor for body field.
	bookDescBody := bookFields[2].Descriptor()
	// book.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	book.BodyValidator = bookDescBody.Validators[0].(func(string) error)
	// bookDescCreatedAt is the schema descriptor for created_at field.
	bookDescCreatedAt := bookFields[3].Descriptor()
	// book.DefaultCreatedAt holds the default value on creation for the created_at field.
	book.DefaultCreatedAt = bookDescCreatedAt.Default.(func() time.Time)
	// bookDescUpdatedAt is the schema descriptor for updated_at field.
	bookDescUpdatedAt := bookFields[4].Descriptor()
	// book.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	book.DefaultUpdatedAt = bookDescUpdatedAt.Default.(func() time.Time)
	// book.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	book.UpdateDefaultUpdatedAt = bookDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
