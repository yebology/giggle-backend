package post

// Type represents the type of a post.
type Type string

// Constants for possible post types.
const (

	Hire    Type = "Hire"
	Service Type = "Service"
	
)

// AllowedType lists all valid post types.
var AllowedType = []Type{
	Hire, Service,
}
