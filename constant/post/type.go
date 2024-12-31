package post

type Type string

const (

	Hire Type = "Hire"
	Service Type = "Service"

)

var AllowedType = []Type{
	Hire, Service,
}