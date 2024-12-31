package data

type Account struct {

	Email		string		`json:"email" validate:"required,email"`

}