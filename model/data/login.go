package data

type Login struct {

	UserIdentifier 		string 		`json:"userIdentifier" validate:"required,min=8"`
	Password 			string 		`json:"password" validate:"required,min=8"`

}