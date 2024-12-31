package constant

type Role string

const (

	User Role = "user"
	Guest Role = "guest"
	
)

var AllowedRole = []Role{
	User, Guest,
}