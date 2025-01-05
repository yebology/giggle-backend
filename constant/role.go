package constant

// Role represents different user roles in the application.
type Role string

// Constants for the available user roles.
const (
	User  Role = "user"
	Guest Role = "guest"
)

// AllowedRole lists all valid user roles.
var AllowedRole = []Role{
	User, Guest,
}